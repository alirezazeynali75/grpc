package todo

import (
	"context"

	"github.com/alirezazeynali75/grpc/pkg/models"
	"github.com/sirupsen/logrus"
)

type toDoRepo interface {
	AddNewToDo(ctx context.Context, title string, description string, status string) error
	GetToDoById(ctx context.Context, id int) (models.ToDo, error)
	GetToDoByStatus(ctx context.Context, status string) ([]models.ToDo, error)
}

type ToDoService struct {
	logger *logrus.Logger
	repo   toDoRepo
}

func NewToDoService(logger *logrus.Logger, repo toDoRepo) *ToDoService {
	return &ToDoService{
		logger: logger.WithField("module", "ToDoService").Logger,
		repo: repo,
	}
}


func (svc *ToDoService) AddNew(ctx context.Context, title string, description string, status string) error {
	return svc.repo.AddNewToDo(ctx, title, description, status)
}

func (svc *ToDoService) GetById(ctx context.Context, id int) ( models.ToDo, error) {
	return svc.repo.GetToDoById(ctx, id)
}

func (svc *ToDoService) GetByStatus(ctx context.Context, status string) ([]models.ToDo, error) {
	return svc.repo.GetToDoByStatus(ctx, status)
}