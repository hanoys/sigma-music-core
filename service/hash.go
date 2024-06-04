package service

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/hanoys/sigma-music-core/domain"
	"time"
)

type HashPasswordService struct {
}

func NewHashPasswordService() *HashPasswordService {
	return &HashPasswordService{}
}

func (h *HashPasswordService) genSalt() string {
	salt := sha256.Sum256([]byte(time.Now().String()))
	return hex.EncodeToString(salt[:])
}

func (h *HashPasswordService) EncodePassword(password string) domain.SaltedPassword {
	salt := h.genSalt()
	hash := sha256.Sum256([]byte(password + salt))
	return domain.SaltedPassword{
		HashPassword: hex.EncodeToString(hash[:]),
		Salt:         salt,
	}
}
func (h *HashPasswordService) ComparePasswordWithHash(password string, saltedPassword domain.SaltedPassword) bool {
	passwordHash := sha256.Sum256([]byte(password + saltedPassword.Salt))
	return hex.EncodeToString(passwordHash[:]) == saltedPassword.HashPassword
}
