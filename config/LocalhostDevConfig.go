package config

// BechamelRuntimeConfig configuration implementation for local development
// with this running on localhost and not using any external resources
type LocalhostDevConfig struct {
}

func (l *LocalhostDevConfig) PasswordSalt() []byte {
	return []byte("Don't be salty!")
}

func (l *LocalhostDevConfig) AccessJWTTTL() int {
	return 10 * 60 // 10 minutes
}

func (l *LocalhostDevConfig) AccessJWTSigningKey() []byte {
	return []byte("GetThisFromENV")
}

func (l *LocalhostDevConfig) RefreshJWTTTL() int {
	return 7 * 24 * 60 * 60 // 7 days
}

func (l *LocalhostDevConfig) RefreshJWTSigningKey() []byte {
	return []byte("GetThisFromENV")
}

func NewLocalhostDevConfig() *LocalhostDevConfig {
	return &LocalhostDevConfig{}
}
