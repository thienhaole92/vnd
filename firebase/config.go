package firebase

import "github.com/caarlos0/env/v10"

type Config struct {
	CredentialsFile string `env:"FIREBASE_CREDENTIALS_FILE,required"`
	DatabaseUrl     string `env:"FIREBASE_DATABASE_URL,required"`
	StorageBucket   string `env:"FIREBASE_STORAGE_BUCKET,required"`
}

func NewConfig() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
