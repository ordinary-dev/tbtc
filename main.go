package main

import (
	"encoding/base64"
	"errors"
	"log"
	"os"
	"time"
)

// Checks the list of certificates and looks for the target domain
func findCertificateFromProvider(provider CertProvider, targetDomain string) (string, string, error) {
	for _, record := range provider.Certificates {
		if record.Domain.Main == targetDomain {
			return record.Cert, record.Key, nil
		}
	}
	return "", "", errors.New("domain was not found")
}

// Decodes the string and writes the result to a file
func writeBase64ToFile(content string, filepath string) error {
	// Decode base64 string
	decodedContent, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return err
	}

	// Create a file
	outputFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Write bytes to file
	outputFile.Write(decodedContent)
	return nil
}

// Main function that is called periodically.
func updateCertificates(config EnvConfig) error {
	// Parse information about certificates
	certProviders, err := ParseAcmeJsonFile(config.AcmeFilePath)
	if err != nil {
		return err
	}

	filesWereUpdated := false
	for _, provider := range certProviders {
		// Try to find the encoded certificate and key
		cert64, key64, err := findCertificateFromProvider(provider, config.TargetDomain)
		// The required domain was not found, try the next provider
		if err != nil {
			continue
		}

		// Save the certificate
		err = writeBase64ToFile(cert64, config.CertificateFilePath)
		if err != nil {
			return err
		}

		// Save the key
		err = writeBase64ToFile(key64, config.KeyFilePath)
		if err != nil {
			return err
		}

		filesWereUpdated = true
		break
	}

	if !filesWereUpdated {
		return errors.New("target domain was not found")
	}

	return nil
}

func main() {
	config, err := GetConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = updateCertificates(config)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("The certificate and key have been updated, sleeping for 3 hours")
		time.Sleep(3 * time.Hour)
	}
}
