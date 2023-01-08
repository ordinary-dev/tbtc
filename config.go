// Environment variable handler
package main

import (
	"errors"
	"os"
)

type EnvConfig struct {
	TargetDomain        string
	AcmeFilePath        string
	CertificateFilePath string
	KeyFilePath         string
}

// Returns the value of an environment variable,
// or the default value if the variable does not exist
func getEnvVarOrDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

// Get settings from environment variables
func GetConfigFromEnv() (EnvConfig, error) {
	// Get target domain
	targetDomain, ok := os.LookupEnv("TBTC_TARGET_DOMAIN")
	if !ok {
		return EnvConfig{}, errors.New("TBTC_TARGET_DOMAIN is undefined")
	}

	// Get acme.json file path
	var acmeFilePath = getEnvVarOrDefault("TBTC_ACME_FILE_PATH", "acme.json")

	// Get certificate path
	var certificateFilePath = getEnvVarOrDefault("TBTC_CERTIFICATE_FILE_PATH", "fullchain.pem")

	// Get key path
	var keyFilePath = getEnvVarOrDefault("TBTC_KEY_FILE_PATH", "privkey.pem")

	var config = EnvConfig{
		targetDomain,
		acmeFilePath,
		certificateFilePath,
		keyFilePath,
	}
	return config, nil
}
