package handler

import (
	"ginlayout/internal/request"
	"net/http"

	"ginlayout/internal/service"
	"ginlayout/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// Register godoc
// @Summary User Registration
// @Description Register a new user
// @Tags user
// @Accept json
// @Produce json
// @Param request body request.RegisterReq true "Register Request"
// @Success 200 {object} response.Response
// @Router /api/v1/user/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req request.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.svc.Register(c, req.Username, req.Password, req.Email); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}
