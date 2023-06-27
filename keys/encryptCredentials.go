package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// DarajaCredentials Generate the daraja credentials from the provided certificate and the passKey provided from the portal -.
func DarajaCredentials(passKey string) (string, error) {
	certPath := "keys/ProductionCertificate.cer"

	// Read the certificate file -.
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		return "", fmt.Errorf("failed to read certificate file: %w", err)
	}

	// Decode the PEM-encoded certificate
	block, _ := pem.Decode(certData)

	if block == nil {
		return "", fmt.Errorf("failed to decode PEM certificate")
	}

	// Parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)

	if err != nil {
		return "", fmt.Errorf("failed to parse certificate: %w", err)
	}

	// Encrypt the plaintext using the certificate's public key -.
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, cert.PublicKey.(*rsa.PublicKey), []byte(passKey))

	if err != nil {
		return "", fmt.Errorf("failed to encrypt plaintext: %w", err)
	}

	// Encode the encrypted data as base64 -.
	encodedSignature := base64.StdEncoding.EncodeToString(encryptedBytes)

	return encodedSignature, nil
}
