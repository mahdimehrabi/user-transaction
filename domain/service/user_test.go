package service

import (
	entity "bbdk/domain/entity"
	mock_logger "bbdk/mock/infrastructure"
	mock_user "bbdk/mock/repository"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestUserService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	var tests = []struct {
		name         string
		loggerMock   func() *mock_logger.MockLogger
		userRepoMock func() *mock_user.MockRepository
		user         *entity.User
		ctx          context.Context
		error        error
	}{
		{
			name: "success",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(nil)
				return userRepoMock
			},
			user:  &entity.User{Name: "fsddfs", Email: "ma@gmail.com", Password: "A12345678"},
			ctx:   context.Background(),
			error: nil,
		},
		{
			name: "UserRepoError",
			loggerMock: func() *mock_logger.MockLogger {
				loggerInfra := mock_logger.NewMockLogger(ctrl)
				loggerInfra.EXPECT().Errorf(gomock.Any(), gomock.Any())
				return loggerInfra
			},
			userRepoMock: func() *mock_user.MockRepository {
				userRepoMock := mock_user.NewMockRepository(ctrl)
				userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(err)
				return userRepoMock
			},
			user:  &entity.User{Name: "fsddfs", Email: "ma@gmail.com", Password: "A12345678"},
			ctx:   context.Background(),
			error: err,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userRepoMock := test.userRepoMock()
			loggerMock := test.loggerMock()
			service := NewUserService(userRepoMock, loggerMock)
			err := service.CreateUser(test.user)

			if !errors.Is(err, test.error) {
				t.Errorf("error:%s is not equal to %s", err, test.error)
			}
			loggerMock.EXPECT()
			userRepoMock.EXPECT()
		})
	}
}

func BenchmarkService_Create(b *testing.B) {
	ctrl := gomock.NewController(b)
	userRepoMock := mock_user.NewMockRepository(ctrl)
	userRepoMock.EXPECT().CreateUser(gomock.Any()).Return(nil)
	loggerMock := mock_logger.NewMockLogger(ctrl)
	b.ResetTimer()
	service := NewUserService(userRepoMock, loggerMock)
	user := &entity.User{Name: "fsddfs", Email: "ma@gmail.com", Password: "A12345678"}
	service.CreateUser(user)
	if b.Elapsed() > 100*time.Microsecond {
		b.Error("address_user service-createBatchAddresses takes too long to run")
	}
	loggerMock.EXPECT()
	userRepoMock.EXPECT()
}
