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

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	var tests = []struct {
		name         string
		addresses    []*entity.User
		loggerMock   func() *mock_logger.MockLogger
		userRepoMock func() *mock_user.MockRepository
		user         *entity.User
		ctx          context.Context
	}{
		{
			name: "success",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				return loggerInfra
			},
			addressRepoMock: func() *repository.MockAddress {
				addrRepoMock := repository.NewMockAddress(ctrl)
				addrRepoMock.EXPECT().BatchCreate(gomock.Any(), gomock.Any()).Return(nil)
				return addrRepoMock
			},
			userRepoMock: func() *repository.MockUser {
				userRepoMock := repository.NewMockUser(ctrl)
				userRepoMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				return userRepoMock
			},
			addresses: []*entity.Address{entity.NewAddress("c", "s", "co", "str", "3tgdsgds")},
			user:      &entity.User{Name: "Dgsgds", Lastname: "sfafsasf"},
			ctx:       context.Background(),
		},
		{
			name: "AddrRepoError",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).MinTimes(2).Return()
				return loggerInfra
			},
			addressRepoMock: func() *repository.MockAddress {
				addrRepoMock := repository.NewMockAddress(ctrl)
				addrRepoMock.EXPECT().BatchCreate(gomock.Any(), gomock.Any()).MinTimes(2).Return(err)
				return addrRepoMock
			},
			userRepoMock: func() *repository.MockUser {
				userRepoMock := repository.NewMockUser(ctrl)
				userRepoMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				return userRepoMock
			},
			addresses: []*entity.Address{entity.NewAddress("c", "s", "co", "str", "3tgdsgds")},
			ctx:       context.Background(),
		},
		{
			name: "UserRepoError",
			loggerMock: func() *mock_log.MockLog {
				loggerInfra := mock_log.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).MinTimes(2).Return()
				return loggerInfra
			},
			addressRepoMock: func() *repository.MockAddress {
				addrRepoMock := repository.NewMockAddress(ctrl)
				return addrRepoMock
			},
			userRepoMock: func() *repository.MockUser {
				userRepoMock := repository.NewMockUser(ctrl)
				userRepoMock.EXPECT().Create(gomock.Any(), gomock.Any()).MinTimes(2).Return(int64(0), err)
				return userRepoMock
			},
			addresses: []*entity.Address{entity.NewAddress("c", "s", "co", "str", "3tgdsgds")},
			ctx:       context.Background(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			addressRepoMock := test.addressRepoMock()
			userRepoMock := test.userRepoMock()
			loggerMock := test.loggerMock()
			service := NewService(loggerMock, addressRepoMock, userRepoMock)
			service.Create(test.addresses, test.user)

			time.Sleep(600 * time.Millisecond)
			loggerMock.EXPECT()
			addressRepoMock.EXPECT()
			userRepoMock.EXPECT()
		})
	}
}
