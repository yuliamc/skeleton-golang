package partner

import (
	"context"
	"modalrakyat/skeleton-golang/internal/model"
)

func (r *partnerRepo) Find(ctx context.Context, where *model.Partner) (*model.Partner, error) {
	var partnerModel *model.Partner
	query := r.dbClient.Get(ctx).Where(where)

	if err := query.Find(&partnerModel).Error; err != nil {
		return nil, err
	}

	if query.RowsAffected == 0 {
		return nil, nil
	}

	return partnerModel, nil
}
