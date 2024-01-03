package product

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type Handler struct {
	Service ServiceI
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		Service: s,
	}
}
func (h *Handler) GetAll(ctx *gin.Context) {
	opts := &options.FindOptions{}
	result, err := h.Service.GetAll(ctx, bson.M{}, opts)
	if err != nil {
		logrus.Errorf("Error when getting Product %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
