package handler

import (
	"errors"
	"fmt"
	"net/http"
	"seleksi-javan/model/user"
	ucuser "seleksi-javan/usecase/uc_user"
	"seleksi-javan/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUsecase ucuser.UserUsecase
}

func NewUserHandler(userUsecase ucuser.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (uh *UserHandler) Register(c *gin.Context) {
	var registerReq user.RegisterRequest
	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]string, len(ve))
			for i, fe := range ve {
				field := fe.Field()
				tag := fe.Tag()
				switch tag {
				case "required":
					out[i] = fmt.Sprintf("%s is required", field)
				case "min":
					out[i] = fmt.Sprintf("%s must be at least %s characters long", field, fe.Param())
				case "email":
					out[i] = "Email format is invalid"
				default:
					out[i] = fmt.Sprintf("%s failed on '%s' validation", field, tag)
				}
			}

			c.JSON(http.StatusBadRequest, util.ApiResponse{
				Success: false,
				Message: strings.Join(out, ", "),
			})
			return
		}

		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Invalid request payload",
		})
		return
	}

	if err := uh.userUsecase.Register(registerReq); err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Failed to register",
		})
		return
	}

	c.JSON(http.StatusOK, util.ApiResponse{
		Success: true,
	})
}

func (uh *UserHandler) Login(c *gin.Context) {
	loginReq := user.LoginRequest{}
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]string, len(ve))
			for i, fe := range ve {
				field := fe.Field()
				tag := fe.Tag()
				switch tag {
				case "required":
					out[i] = fmt.Sprintf("%s is required", field)
				case "min":
					out[i] = fmt.Sprintf("%s must be at least %s characters long", field, fe.Param())
				case "max":
					out[i] = fmt.Sprintf("%s must be shorter than %s characters", field, fe.Param())
				default:
					out[i] = fmt.Sprintf("%s failed on '%s' validation", field, tag)
				}
			}

			c.JSON(http.StatusBadRequest, util.ApiResponse{
				Success: false,
				Message: strings.Join(out, ", "),
			})
			return
		}

		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Invalid request payload",
		})
		return
	}

	authResp, err := uh.userUsecase.Login(loginReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiResponse{
			Success: false,
			Message: "Failed to login",
		})
		return
	}

	c.JSON(http.StatusOK, util.ApiResponse{
		Success: true,
		Data: gin.H{
			"token": gin.H{
				"access_token":  authResp.Token.AccessToken,
				"refresh_token": authResp.Token.RefreshToken,
			},
			"user": gin.H{
				"username": authResp.User.Username,
				"email":    authResp.User.Email,
			},
		},
	})
}

func (uh *UserHandler) ChangePassword(c *gin.Context) {
	id := c.Param("id")
	targetId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Bad Request param",
		})
		return
	}

	var changePasswordRequest user.ChangePasswordRequest
	err = c.ShouldBindJSON(&changePasswordRequest)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]string, len(ve))
			for i, fe := range ve {
				field := fe.Field()
				tag := fe.Tag()
				switch tag {
				case "required":
					out[i] = fmt.Sprintf("%s is required", field)
				case "min":
					out[i] = fmt.Sprintf("%s must be at least %s characters long", field, fe.Param())
				case "max":
					out[i] = fmt.Sprintf("%s must be shorter than %s characters", field, fe.Param())
				default:
					out[i] = fmt.Sprintf("%s failed on '%s' validation", field, tag)
				}
			}

			c.JSON(http.StatusBadRequest, util.ApiResponse{
				Success: false,
				Message: strings.Join(out, ", "),
			})
			return
		}

		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Invalid request payload",
		})
		return
	}

	err = uh.userUsecase.ChangePassword(uint(targetId), changePasswordRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiResponse{
			Success: false,
			Message: "Internal server error",
		})

		return
	}

	c.JSON(http.StatusOK, util.ApiResponse{
		Success: true,
	})
}

func (uh *UserHandler) GetAllUser(c *gin.Context) {
	users, err := uh.userUsecase.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiResponse{
			Success: false,
			Message: "Failed to get users",
		})
		return
	}

	c.JSON(http.StatusOK, util.ApiResponse{
		Success: true,
		Data:    users,
	})
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	userResp, err := uh.userUsecase.GetUserByID(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiResponse{
			Success: false,
			Message: "Failed to get user",
		})
		return
	}

	c.JSON(http.StatusOK, util.ApiResponse{
		Success: true,
		Data:    userResp,
	})
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	targetId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Bad Request param",
		})
		return
	}

	var updateUserRequest user.UpdateUserRequest
	err = c.ShouldBindJSON(&updateUserRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Bad Request",
		})
		return
	}

	err = uh.userUsecase.UpdateUser(uint(targetId), updateUserRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiResponse{
			Success: false,
			Message: "Internal server error",
		})

		return
	}

	c.JSON(http.StatusOK, util.ApiResponse{
		Success: true,
	})
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	targetId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Bad Request param",
		})
		return
	}

	err = uh.userUsecase.DeleteUser(uint(targetId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ApiResponse{
			Success: false,
			Message: "Internal server error",
		})

		return
	}

	c.JSON(http.StatusOK, util.ApiResponse{
		Success: true,
	})
}

func (uh *UserHandler) Route(r *gin.Engine, authMiddleware gin.HandlerFunc) *gin.Engine {
	public := r.Group("/api/user")

	public.POST("/register", uh.Register)
	public.POST("/login", uh.Login)

	protected := public.Group("/", authMiddleware)

	protected.PATCH("/:id/password", uh.ChangePassword)
	protected.GET("/", uh.GetAllUser)
	protected.GET("/:id", uh.GetUserByID)
	protected.PATCH("/:id", uh.UpdateUser)
	protected.DELETE("/:id", uh.DeleteUser)

	return r
}
