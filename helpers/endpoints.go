package helpers

import "context"

type EncryptService interface {
	Encrypt(context.Context, string, string) (string, error)
	Decrypt(context.Context, string, string) (string, error)
}
