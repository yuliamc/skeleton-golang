package partner

import (
	"context"
	"fmt"
	"modalrakyat/skeleton-golang/internal/model"
	"modalrakyat/skeleton-golang/pkg/utils/null"
)

func (s *partnerService) FindByID(ctx context.Context, ID *uint) (*PartnerResponse, error) {
	where := model.Partner{
		ID: *ID,
	}

	partnerModel, err := s.partnerRepo.Find(ctx, &where)
	if err != nil {
		fmt.Println("17", partnerModel)
		return nil, err
	}
	if null.IsNil(partnerModel) {
		fmt.Println("20", partnerModel)
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
