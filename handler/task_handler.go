package handler

import (
	"errors"
	"fmt"
	"net/http"
	"seleksi-javan/model/task"
	uctask "seleksi-javan/usecase/uc_task"
	"seleksi-javan/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TaskHandler struct {
	taskUsecase uctask.TaskUsecase
}

func NewTaskHandler(taskUsecase uctask.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		taskUsecase: taskUsecase,
	}
}

func (th *TaskHandler) AddTask(c *gin.Context) {
	var taskReq task.TaskRequest

	err := c.ShouldBindJSON(&taskReq)
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
				case "oneof":
					out[i] = fmt.Sprintf("%s must be one of: %s", field, fe.Param())
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

	err = th.taskUsecase.AddTask(taskReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Failed to add task",
		})
		return
	}

	c.JSON(http.StatusOK, util.ApiResponse{
		Success: true,
	})
}

func (th *TaskHandler) GetAllTask(c *gin.Context) {
	tasks, err := th.taskUsecase.GetAllTask()
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Failed to get tasks",
		})
		return
	}

	c.JSON(http.StatusOK, util.ApiResponse{
		Success: true,
		Data:    tasks,
	})
}

func (th *TaskHandler) GetTaskByID(c *gin.Context) {
	taskIdParam := c.Param("id")
	taskId, err := strconv.Atoi(taskIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Invalid task ID",
		})
		return
	}

	taskResp, err := th.taskUsecase.GetTaskByID(uint(taskId))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Failed to get task",
		})
		return
	}

	c.JSON(http.StatusOK, util.ApiResponse{
		Success: true,
		Data:    taskResp,
	})
}

func (th *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	targetId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Bad Request param",
		})
		return
	}

	var updateTask task.TaskUpdateRequest
	err = c.ShouldBindJSON(&updateTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Bad Request",
		})
		return
	}

	err = th.taskUsecase.UpdateTask(uint(targetId), updateTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Failed to update task",
		})
		return
	}

	c.JSON(http.StatusOK, util.ApiResponse{
		Success: true,
	})
}

func (th *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	targetId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Bad Request param",
		})
		return
	}

	err = th.taskUsecase.DeleteTask(uint(targetId))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ApiResponse{
			Success: false,
			Message: "Failed to delete task",
		})
		return
	}

	c.JSON(http.StatusOK, util.ApiResponse{
		Success: true,
	})
}

func (th *TaskHandler) Route(r *gin.Engine, authMiddleware gin.HandlerFunc) *gin.Engine {
	protected := r.Group("/api/task", authMiddleware)
	protected.POST("/", th.AddTask)
	protected.GET("/", th.GetAllTask)
	protected.GET("/:id", th.GetTaskByID)
	protected.PATCH("/:id", th.UpdateTask)
	protected.DELETE("/:id", th.DeleteTask)

	return r
}
