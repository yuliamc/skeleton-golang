package partner

import (
	"context"
	"modalrakyat/skeleton-golang/internal/model"
)

func (r *partnerRepo) MLGetDropdown(ctx context.Context, fields ...string) ([]*model.Partner, error) {
	partnerModels := []*model.Partner{}
	query := r.dbdget.Get(ctx)

	if len(fields) == 0 {
		// Set default selected fields
		fields = []string{"code", "name"}
	}

	err := query.
		Select(fields).
		Order(model.Partner{}.TableName() + ".name ASC").
		Find(&partnerModels).
		Error

	if err != nil {
		return nil, err
	}

	return partnerModels, nil
}
