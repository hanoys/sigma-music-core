package test

//
//import (
//	"context"
//	"errors"
//	"github.com/google/uuid"
//	"github.com/hanoys/sigma-music-core/domain"
//	"github.com/hanoys/sigma-music-core/ports"
//	"github.com/hanoys/sigma-music-core/service"
//	mocks2 "github.com/hanoys/sigma-music-core/service/test/mocks/auth"
//	mocks3 "github.com/hanoys/sigma-music-core/service/test/mocks/hash"
//	"github.com/hanoys/sigma-music-core/service/test/mocks/repository"
//	"go.uber.org/zap"
//	"testing"
//)
//
//var credentials = ports.LogInCredentials{
//	Name:     "test",
//	Password: "test",
//}
//
//var foundUser = domain.User{
//	ID:       uuid.New(),
//	Name:     "test",
//	Email:    "test",
//	Phone:    "test",
//	Password: "test",
//	Salt:     "test",
//	Country:  "test",
//}
//
//func TestAuthServiceLogIn(t *testing.T) {
//	tests := []struct {
//		name              string
//		userRepoMock      func(repository *mocks.UserRepository)
//		musicianRepoMock  func(repository *mocks.MusicianRepository)
//		tokenProviderMock func(provider *mocks2.TokenProvider)
//		expected          error
//	}{
//		{
//			name: "success login",
//			userRepoMock: func(repository *mocks.UserRepository) {
//				repository.
//					On("GetByName", context.Background(), credentials.Name).
//					Return(foundUser, nil)
//			},
//			musicianRepoMock: func(repository *mocks.MusicianRepository) {
//				repository.
//					On("GetByName", context.Background(), credentials.Name).
//					Return(foundUser, nil)
//			},
//			tokenProviderMock: func(provider *mocks2.TokenProvider) {
//				provider.
//					On("NewSession", context.Background(), domain.Payload{
//						UserID: foundUser.ID,
//						Role:   domain.UserRole,
//					}).Return(domain.TokenPair{}, nil)
//			},
//			expected: nil,
//		},
//		{
//			name: "fail login: incorrect password",
//			userRepoMock: func(repository *mocks.UserRepository) {
//				repository.
//					On("GetByName", context.Background(), credentials.Name).
//					Return(foundUser, nil)
//			},
//			musicianRepoMock: func(repository *mocks.MusicianRepository) {
//
//			},
//			tokenProviderMock: func(provider *mocks2.TokenProvider) {
//
//			},
//			expected: ports.ErrIncorrectPassword,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			userRepository := mocks.NewUserRepository(t)
//			musicianRepository := mocks.NewMusicianRepository(t)
//			tokenProvider := mocks2.NewTokenProvider(t)
//			hashProvider := service.NewHashPasswordService()
//			logger, _ := zap.NewProduction()
//			authService := service.NewAuthorizationService(userRepository, musicianRepository, tokenProvider, hashProvider, logger)
//			test.userRepoMock(userRepository)
//			test.musicianRepoMock(musicianRepository)
//			test.tokenProviderMock(tokenProvider)
//
//			_, err := authService.LogIn(context.Background(), credentials)
//			if !errors.Is(err, test.expected) {
//				t.Errorf("got %v, want %v", err, test.expected)
//			}
//		})
//	}
//}
//
//var tokenString = "testTokenString"
//
//func TestAuthServiceVerifyToken(t *testing.T) {
//	tests := []struct {
//		name              string
//		userRepoMock      func(repository *mocks.UserRepository)
//		musicianRepoMock  func(repository *mocks.MusicianRepository)
//		tokenProviderMock func(provider *mocks2.TokenProvider)
//		hashProviderMock  func(provider *mocks3.HashPasswordProvider)
//		expected          error
//	}{
//		{
//			name: "success verification",
//			userRepoMock: func(repository *mocks.UserRepository) {
//			},
//			musicianRepoMock: func(repository *mocks.MusicianRepository) {
//			},
//			tokenProviderMock: func(provider *mocks2.TokenProvider) {
//				provider.
//					On("VerifyToken", context.Background(), tokenString).
//					Return(domain.Payload{}, nil)
//			},
//			hashProviderMock: func(provider *mocks3.HashPasswordProvider) {
//			},
//			expected: nil,
//		},
//		{
//			name: "fail verification",
//			userRepoMock: func(repository *mocks.UserRepository) {
//			},
//			musicianRepoMock: func(repository *mocks.MusicianRepository) {
//			},
//			tokenProviderMock: func(provider *mocks2.TokenProvider) {
//				provider.
//					On("VerifyToken", context.Background(), tokenString).
//					Return(domain.Payload{}, ports.ErrTokenProviderInvalidToken)
//			},
//			hashProviderMock: func(provider *mocks3.HashPasswordProvider) {
//			},
//			expected: ports.ErrTokenProviderInvalidToken,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			userRepository := mocks.NewUserRepository(t)
//			musicianRepository := mocks.NewMusicianRepository(t)
//			tokenProvider := mocks2.NewTokenProvider(t)
//			logger, _ := zap.NewProduction()
//			hashProvider := service.NewHashPasswordService()
//			authService := service.NewAuthorizationService(userRepository, musicianRepository, tokenProvider, hashProvider, logger)
//			test.userRepoMock(userRepository)
//			test.musicianRepoMock(musicianRepository)
//			test.tokenProviderMock(tokenProvider)
//
//			_, err := authService.VerifyToken(context.Background(), tokenString)
//			if !errors.Is(err, test.expected) {
//				t.Errorf("got %v, want %v", err, test.expected)
//			}
//		})
//	}
//}
