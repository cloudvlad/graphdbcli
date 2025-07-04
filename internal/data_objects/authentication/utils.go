package authentication

import (
	"graphdbcli/internal/tool_configurations/logging"
	"os"
)

func SetupGraphDBAuthentication() {
	logging.LOGGER.Debug("Setting GraphDB authentication...")

	if AuthToken.AuthToken == "" {
		logging.LOGGER.Debug("Bearer Token Authentication was not set with flag")
		logging.LOGGER.Debug("Searching for environment variable 'GRAPHDB_AUTH_TOKEN' as an alternative")
		value, isSet := os.LookupEnv("GRAPHDB_AUTH_TOKEN")
		if isSet {
			AuthToken.AuthToken = value
			return
		} else {
			logging.LOGGER.Debug("Environment variable 'GRAPHDB_AUTH_TOKEN' has not been found")
		}
	}

	if BasicCredentials.Username == "" {
		logging.LOGGER.Debug("Basic Authentication Username was not set with flag")
		logging.LOGGER.Debug("Searching for environment variable 'GRAPHDB_USERNAME' as an alternative")
		value, isSet := os.LookupEnv("GRAPHDB_USERNAME")
		if isSet {
			BasicCredentials.Username = value
		} else {
			logging.LOGGER.Debug("Environment variable 'GRAPHDB_USERNAME' has not been found")
		}
	}

	if BasicCredentials.Password == "" {
		logging.LOGGER.Debug("Basic Authentication Password was not set with flag")
		logging.LOGGER.Debug("Searching for environment variable 'GRAPHDB_PASSWORD' as an alternative")
		value, isSet := os.LookupEnv("GRAPHDB_PASSWORD")
		if isSet {
			BasicCredentials.Password = value
		} else {
			logging.LOGGER.Debug("Environment variable 'GRAPHDB_PASSWORD' has not been found")
		}
	}
}

func SetupS3Authentication() {
	logging.LOGGER.Debug("Setting S3 authentication...")
	if S3Authentication.AccessKeyID == "" {
		logging.LOGGER.Debug("S3 Access Key ID was not set with flag")
		logging.LOGGER.Debug("Searching for environment variable 'S3_ACCESS_KEY_ID' as an alternative")
		value, isSet := os.LookupEnv("S3_ACCESS_KEY_ID")
		if isSet {
			S3Authentication.AccessKeyID = value
		} else {
			logging.LOGGER.Debug("Environment variable 'S3_ACCESS_KEY_ID' has not been found")
		}
	}

	if S3Authentication.SecretAccessKey == "" {
		logging.LOGGER.Debug("S3 Secret Access Key was not set with flag")
		logging.LOGGER.Debug("Searching for environment variable 'S3_ACCESS_KEY_TOKEN' as an alternative")
		value, isSet := os.LookupEnv("S3_ACCESS_KEY_TOKEN")
		if isSet {
			S3Authentication.SecretAccessKey = value
		} else {
			logging.LOGGER.Debug("Environment variable 'S3_ACCESS_KEY_TOKEN' has not been found")
		}
	}
}
