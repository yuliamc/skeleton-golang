package partner

import (
	"context"
	"modalrakyat/skeleton-golang/internal/model"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
)

func (r *partnerRepo) GetByID(ctx context.Context, ID uint) (*model.Partner, error) {
	partner := model.Partner{}
	partner.ID = ID

	query := r.dbdget.Get(ctx).Where(partner)

	if err := query.Find(&partner).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, errors.NewGenericError(errors.DATA_NOT_FOUND)
	}

	return &partner, nil
}
