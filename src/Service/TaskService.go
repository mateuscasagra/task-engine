package Service

import (
	"github.com/gin-gonic/gin"
	dto "github.com/task-engine/src/DTO"
	"github.com/task-engine/src/Repository"
)

type TaskService struct {
}

func CreateTask(ctx *gin.Context) {
	var taskDto dto.TaskDTO
	ctx.ShouldBindBodyWithJSON(&taskDto)
	taskRepo := Repository.GetInstanceTask()
	defer taskRepo.Instance.Close()

	// idempotency := ctx.Request.Header.Get("X-Idempotency-Key")

}
