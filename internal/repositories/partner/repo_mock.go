package partner

import (
	"context"
	"modalrakyat/skeleton-golang/internal/model"

	"github.com/stretchr/testify/mock"
)

type PartnerRepositoryMock struct {
	Mock mock.Mock
}

func (r *PartnerRepositoryMock) GetByID(ctx context.Context, ID uint) (*model.Partner, error) {
	arguments := r.Mock.Called(ctx, ID)
	if arguments.Get(0) != nil {
		obj := arguments.Get(0).(model.Partner)
		return &obj, nil
	} else {
		return nil, arguments.Error(1)
	}
}

func (r *PartnerRepositoryMock) GetDropdown(ctx context.Context, fields ...string) ([]*model.Partner, error) {
	arguments := r.Mock.Called(ctx, fields)
	if arguments.Get(0) != nil {
		objs := arguments.Get(0).([]*model.Partner)
		return objs, nil
	} else {
		return nil, arguments.Error(1)
	}
}
