package partner

import (
	"context"
	"modalrakyat/skeleton-golang/internal/model"

	"gorm.io/gorm"
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

func (r *partnerRepo) FindAll(ctx context.Context) (*[]model.Partner, error) {
	var partnerModels []model.Partner
	query := r.baseFindAllQuery(ctx)

	if err := query.Find(&partnerModels).Error; err != nil {
		return nil, err
	}

	return &partnerModels, nil
}

func (r *partnerRepo) CountAll(ctx context.Context) (*int64, error) {
	var count int64
	query := r.baseFindAllQuery(ctx)

	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
}

func (r *partnerRepo) Create(ctx context.Context, partner *model.Partner) error {
	if err := r.dbClient.Get(ctx).Create(partner).Error; err != nil {
		return err
	}
	return nil
}

func (r *partnerRepo) baseFindAllQuery(ctx context.Context) *gorm.DB {
	return r.dbClient.Get(ctx).Model(model.Partner{})
}
