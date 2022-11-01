package partner

import (
	"context"
	"modalrakyat/skeleton-golang/internal/model"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"modalrakyat/skeleton-golang/pkg/utils/null"
	stringutil "modalrakyat/skeleton-golang/pkg/utils/strings"
	"modalrakyat/skeleton-golang/pkg/utils/syncs"
	"sync/atomic"
)

func (s *partnerService) FindByID(ctx context.Context, ID *uint) (*PartnerResponse, error) {
	where := model.Partner{
		ID: *ID,
	}

	partnerModel, err := s.partnerRepo.Find(ctx, &where)
	if err != nil {
		return nil, err
	}
	if null.IsNil(partnerModel) {
		return nil, nil
	}

	response := NewPartnerResponse(
		&partnerModel.ID,
		&partnerModel.UniqueID,
		&partnerModel.Code,
		&partnerModel.Name,
		&partnerModel.CreatedAt,
		&partnerModel.UpdatedAt,
	)

	return &response, nil
}

func (s *partnerService) FindAll(ctx context.Context) (*[]PartnerResponse, *int64, error) {
	var (
		asyncGroup       syncs.AsyncGroup
		arrayAtomic      atomic.Value
		int64Atomic      atomic.Int64
		errAtomic        atomic.Value
		partnerResponses []PartnerResponse
		count            int64
	)

	asyncGroup.Async(func() {
		if partnerModels, err := s.partnerRepo.FindAll(ctx); err != nil {
			errAtomic.Store(err)
		} else {
			arrayAtomic.Store(*partnerModels)
		}
	})
	asyncGroup.Async(func() {
		if count, err := s.partnerRepo.CountAll(ctx); err != nil {
			errAtomic.Store(err)
		} else {
			int64Atomic.Store(*count)
		}
	})
	asyncGroup.AsyncWait()

	if err, ok := errAtomic.Load().(error); ok {
		return nil, nil, err
	}

	if partnerModels, ok := arrayAtomic.Load().([]model.Partner); !ok {
		return nil, nil, errors.NewGenericError(errors.ERROR_CODE_INTERNAL_SERVER_ERROR)
	} else {
		for _, v := range partnerModels {
			partnerResponses = append(partnerResponses, NewPartnerResponse(
				&v.ID,
				&v.UniqueID,
				&v.Code,
				&v.Name,
				&v.CreatedAt,
				&v.UpdatedAt,
			))
		}
	}

	count = int64Atomic.Load()

	return &partnerResponses, &count, nil
}

func (s *partnerService) Create(ctx context.Context, payload *CreatePartnerPayload) error {
	return s.txRepo.Run(ctx, func(ctx context.Context) error {
		uniqueID := stringutil.GenerateUniqueID()
		partnerModel := NewPartnerModelCreatePayload(&uniqueID, &payload.Code, &payload.Name)
		return s.partnerRepo.Create(ctx, &partnerModel)
	})
}
