package keys

import (
	"cli_cobra/crypto_utils"
	"cli_cobra/logger"
	"cli_cobra/utils"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type PrivateKeyGen struct {
	outputPath string
	keyBitSize int
	saltSize   int
}

func init() {
	keysCmd.AddCommand(keysGenerateCmd)
	keysGenerateCmd.Flags().String("pub-out", "pub_key.pem", "Path to public key")
	keysGenerateCmd.Flags().String("priv-out", "priv_key.pem", "Path to private key")
	keysGenerateCmd.Flags().Int("priv-size", 2048, "Private key size in bits")
	keysGenerateCmd.Flags().Int("salt-size", 16, "Salt size in bits")

}

var keysGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate RSA keys",
	Long:  `Generate RSA keys`,
	Run: func(cmd *cobra.Command, args []string) {
		pkOut, _ := cmd.Flags().GetString("priv-out")
		pkSize, _ := cmd.Flags().GetInt("priv-size")
		saltSize, _ := cmd.Flags().GetInt("salt-size")

		pkGenConfig := PrivateKeyGen{
			outputPath: pkOut,
			keyBitSize: pkSize,
			saltSize:   saltSize,
		}

		privateKey, err := generatePrivKey(pkGenConfig)
		logger.HaltOnErr(err)

		pubOut, _ := cmd.Flags().GetString("pub-out")

		err = generatePubKey(pubOut, privateKey)

		logger.HaltOnErr(err)
	},
}

func generatePubKey(path string, privKey *rsa.PrivateKey) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("Failed to get absolute path of key")
	}

	pubASN1, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		return fmt.Errorf("Failed to marshal public key")
	}

	file, err := os.Create(absPath)
	if err != nil {
		return fmt.Errorf("Failed to create key file")
	}

	defer file.Close()

	if err := pem.Encode(file, &pem.Block{Type: "RSA PUBLIC KEY", Bytes: pubASN1}); err != nil {
		return fmt.Errorf("Failed to encode public key: %w", err)
	}

	return nil
}

func generatePrivKey(pkGenConfig PrivateKeyGen) (*rsa.PrivateKey, error) {
	absPath, err := filepath.Abs(pkGenConfig.outputPath)
	if err != nil {
		return nil, fmt.Errorf("unable to determine absolute path of private key: %w", err)
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, pkGenConfig.keyBitSize)
	if err != nil {
		return nil, fmt.Errorf("unable to generate private key: %s", err)
	}

	passphrase, err := utils.GetPassphrase()
	if err != nil {
		return nil, fmt.Errorf("unable to get passphrase: %s", err)
	}

	privateKeyPEM := x509.MarshalPKCS1PrivateKey(privateKey)

	salt, err := makeSalt(pkGenConfig.saltSize)
	if err != nil {
		return nil, fmt.Errorf("unable to generate salt: %s", err)
	}

	key, err := crypto_utils.DeriveKey(crypto_utils.KeyDerivationConfig{
		Passphrase: passphrase,
		Salt:       salt,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to derive private key: %s", err)
	}

	crypter, err := crypto_utils.MakeCrypter(key)
	if err != nil {
		return nil, fmt.Errorf("unable to create crypter: %s", err)
	}

	nonce, err := crypto_utils.MakeNonce(crypter)
	if err != nil {
		return nil, fmt.Errorf("unable to create nonce: %s", err)
	}

	encryptedData := crypter.Seal(nil, nonce, privateKeyPEM, nil)

	encryptedPEMBlock := &pem.Block{
		Type:  "ENCRYPTED PRIVATE KEY",
		Bytes: encryptedData,
		Headers: map[string]string{
			"Nonce":                   base64.StdEncoding.EncodeToString(nonce),
			"Salt":                    base64.StdEncoding.EncodeToString(salt),
			"Key-Derivation-Function": "Argon2",
		},
	}

	err = savePrivKeyToPEM(absPath, encryptedPEMBlock)
	if err != nil {
		return nil, fmt.Errorf("unable to save private key: %s", err)
	}

	return privateKey, nil
}

func savePrivKeyToPEM(absPath string, encryptedPEMBlock *pem.Block) error {
	privateKeyFile, err := os.Create(absPath)
	if err != nil {
		return fmt.Errorf("unable to create private key file: %s", err)
	}

	defer privateKeyFile.Close()

	if err := pem.Encode(privateKeyFile, encryptedPEMBlock); err != nil {
		return fmt.Errorf("unable to encode private key: %s", err)
	}

	return nil
}

func makeSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("unable to generate salt: %s", err)
	}

	return salt, nil
}
