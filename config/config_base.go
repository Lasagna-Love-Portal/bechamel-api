package config

// This interface represents configuration parameters for the Bechamel API.
// In a deployed setting, an implementation will typically get values
// from either a secure source like Microsoft Azure Key Vault,
// or from the deployment environment for less security sensitive values

type BechamelRuntimeConfig interface {
	PasswordSalt() []byte

	// Number of seconds generated JWT tokens are valid for before expiration
	// when generating with the method not requiring an expiration period.
	AccessJWTTTL() int
	RefreshJWTTTL() int

	AccessJWTSigningKey() []byte
	RefreshJWTSigningKey() []byte
}

var RuntimeConfig BechamelRuntimeConfig
