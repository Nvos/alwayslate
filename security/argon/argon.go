package argon

import (
	"alwayslate/security"
	"github.com/matthewhartstonge/argon2"
)

type hasher struct {
	cfg argon2.Config
}

func (a *hasher) Decode(value string) (string, error) {
	decode, err := argon2.Decode([]byte(value))
	if err != nil {
		return "", err
	}

	return string(decode.Encode()), nil
}

func (a *hasher) Encode(value string) (string, error) {
	raw, err := a.cfg.HashRaw([]byte(value))
	if err != nil {
		return "", err
	}

	return string(raw.Encode()), nil
}

func (a *hasher) Verify(value string, against string) (bool, error) {
	return argon2.VerifyEncoded([]byte(value), []byte(against))
}

func NewDefaultHasher() security.Hasher {
	cfg := argon2.DefaultConfig()

	return &hasher{
		cfg: cfg,
	}
}