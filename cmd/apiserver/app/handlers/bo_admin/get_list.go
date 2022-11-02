package bo_admin

import (
	"modalrakyat/skeleton-golang/pkg/utils/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BOAdminHandler) GetList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, api.BaseWithMeta{
		Data: []struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			{Name: "Budi Suparjo", Email: "budi.suparjo@modalrakyat.id"},
			{Name: "Jim Gordon", Email: "jim.gordon@modalrakyat.id"},
			{Name: "Thomas Wayne", Email: "thomas.wayne@modalrakyat.id"},
		},
		Meta: map[string]interface{}{
			"total": 2,
		},
	})
}
