package partner

import (
	"context"
	"fmt"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"modalrakyat/skeleton-golang/pkg/utils/logs"
	"time"

	"github.com/go-redis/redis"
)

func (s *partnerService) GetByID(ctx context.Context, ID uint) (*PartnerResponse, error) {
	cacheKey := "partner:" + fmt.Sprint(ID) + ":detail"

	var response *PartnerResponse
	err := s.redisDel.Get(cacheKey, &response)
	if err != nil && err != redis.Nil {
		logs.PushErrorLog(err)
		return nil, errors.NewGenericError(errors.INTERNAL_SERVER_ERROR)
	}

	if response == nil {
		partner, err := s.partnerRepo.GetByID(ctx, ID)
		if err != nil {
			return nil, err
		}

		response = &PartnerResponse{}
		response.PartnerStructResponse(*partner)

		err = s.redisDel.Set(cacheKey, response, time.Duration(5*time.Minute))
		if err != nil && err != redis.Nil {
			logs.PushErrorLog(err)
			return nil, errors.NewGenericError(errors.INTERNAL_SERVER_ERROR)
		}
	}

	return response, nil
}
