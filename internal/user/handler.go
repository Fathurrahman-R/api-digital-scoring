package user

import (
	"api-digital-scoring/internal/helper"
	"api-digital-scoring/internal/user/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("user", h.Create)
	r.PUT("user", h.Update)
	r.GET("user/:id", h.GetById)
	r.DELETE("user/:id", h.Delete)
}

func (h *Handler) GetById(c *gin.Context) {

	u, err := h.service.GetById(c, c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": helper.NewErrorResponse(err.Error())})
		return
	}

	data := BindResponse(&u)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (h *Handler) Create(c *gin.Context) {
	var req dto.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": helper.NewErrorResponse(err.Error())})
		return
	}

	u, err := h.service.Create(c, BindRequest(&req))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": helper.NewErrorResponse(err.Error())})
		return
	}

	data := BindResponse(&u)

	c.JSON(http.StatusCreated, gin.H{"data": data})
}

func (h *Handler) Update(c *gin.Context) {

}

func (h *Handler) Delete(c *gin.Context) {

}
