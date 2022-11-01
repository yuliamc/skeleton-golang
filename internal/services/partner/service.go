package partner

import (
	"context"
	"modalrakyat/skeleton-golang/internal/model"
)

func (s *partnerService) FindByID(ctx context.Context, ID *uint) (*PartnerResponse, error) {
	where := model.Partner{
		ID: *ID,
	}

	partnerModel, err := s.partnerRepo.Find(ctx, &where)
	if err != nil {
		return nil, err
	}

	response := NewPartnerResponse(&partnerModel.ID, &partnerModel.UniqueID, &partnerModel.Code, &partnerModel.Name)

	return &response, nil
}
