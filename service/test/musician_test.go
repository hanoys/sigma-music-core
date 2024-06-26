package test

import (
	"context"
	"errors"
	"github.com/hanoys/sigma-music-core/domain"
	"github.com/hanoys/sigma-music-core/ports"
	"github.com/hanoys/sigma-music-core/service"
	"github.com/hanoys/sigma-music-core/service/test/mocks/repository"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"testing"
)

var createMusicianReq = ports.MusicianServiceCreateRequest{
	Name:        "TestUser",
	Email:       "email",
	Password:    "password",
	Country:     "testCountry",
	Description: "testDescription",
}

func TestMusicianServiceRegister(t *testing.T) {
	tests := []struct {
		name           string
		repositoryMock func(repository *mocks.MusicianRepository)
		req            ports.MusicianServiceCreateRequest
		expected       error
	}{
		{
			name: "register success",
			req:  createMusicianReq,
			repositoryMock: func(repository *mocks.MusicianRepository) {
				repository.
					On("Create", context.Background(), mock.AnythingOfType("domain.Musician")).
					Return(domain.Musician{}, nil).
					On("GetByName", context.Background(), createMusicianReq.Name).
					Return(domain.Musician{}, ports.ErrMusicianNameNotFound).
					On("GetByEmail", context.Background(), createMusicianReq.Email).
					Return(domain.Musician{}, ports.ErrMusicianEmailNotFound)
			},
			expected: nil,
		},
		{
			name: "musician name already exists",
			req:  createMusicianReq,
			repositoryMock: func(repository *mocks.MusicianRepository) {
				repository.
					On("GetByName", context.Background(), createMusicianReq.Name).
					Return(domain.Musician{}, nil)
			},
			expected: ports.ErrMusicianWithSuchNameAlreadyExists,
		},
		{
			name: "musician email already exists",
			req:  createMusicianReq,
			repositoryMock: func(repository *mocks.MusicianRepository) {
				repository.
					On("GetByName", context.Background(), createMusicianReq.Name).
					Return(domain.Musician{}, ports.ErrUserNameNotFound).
					On("GetByEmail", context.Background(), createMusicianReq.Email).
					Return(domain.Musician{}, nil)
			},
			expected: ports.ErrMusicianWithSuchEmailAlreadyExists,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			musicianRepository := mocks.NewMusicianRepository(t)
			logger, _ := zap.NewProduction()
			hashProvider := service.NewHashPasswordService()
			musicianService := service.NewMusicianService(musicianRepository, hashProvider, logger)
			test.repositoryMock(musicianRepository)

			_, err := musicianService.Register(context.Background(), test.req)
			if !errors.Is(err, test.expected) {
				t.Errorf("got %v, want %v", err, test.expected)
			}
		})
	}
}
