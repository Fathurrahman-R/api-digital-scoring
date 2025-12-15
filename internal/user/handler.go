package user

import (
	"api-digital-scoring/internal/entity"
	"api-digital-scoring/internal/user/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	r.POST("/user", h.Create)
	r.PUT("/user", h.Update)
	r.GET("/user/:id", h.GetById)
	r.DELETE("/user/:id", h.Delete)
}

func (h *Handler) GetById(c *gin.Context) {
	u, err := h.service.GetById(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dto.Response{
		ID:       u.ID,
		Fullname: u.Fullname,
		Username: u.Username,
		Email:    u.Email,
	}})
}

func (h *Handler) Create(c *gin.Context) {
	var req dto.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	u, err := h.service.Create(c, &entity.User{
		Fullname: req.Fullname,
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashed),
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dto.Response{
		ID:       u.ID,
		Fullname: u.Fullname,
		Username: u.Username,
		Email:    u.Email,
	}})
}

func (h *Handler) Update(c *gin.Context) {

}

func (h *Handler) Delete(c *gin.Context) {

}
