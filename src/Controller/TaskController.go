package Controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Session *sql.DB
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	//validar o indepotencykey no redis
	// se nao existe -> cria em tabela e adiciona no redis
	// se existe ignora
}

func (c *TaskController) GetTask(ctx *gin.Context) {

}
