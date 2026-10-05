package Repository

import (
	"database/sql"

	"github.com/task-engine/src"
)

type TaskRepository struct {
	Instance *sql.DB
}

func GetInstanceTask() *TaskRepository {
	db := src.Postgres()
	return &TaskRepository{
		Instance: db,
	}
}

func (t *TaskRepository) CreateTask() {

}
