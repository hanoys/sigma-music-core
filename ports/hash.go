package ports

import "github.com/hanoys/sigma-music-core/domain"

type IHashPasswordService interface {
	EncodePassword(password string) domain.SaltedPassword
	ComparePasswordWithHash(password string, saltedPassword domain.SaltedPassword) bool
}
