package data

import (
	"context"

	"github.com/alirezazeynali75/grpc/pkg/models"
	"gorm.io/gorm"
)

type ToDoRepo struct {
	db *gorm.DB
}

func NewToDoRepo(db *gorm.DB) *ToDoRepo {
	return &ToDoRepo{
		db: db,
	}
}

func (repo *ToDoRepo) AddNewToDo(ctx context.Context, title string, description string, status string) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		todo := ToDoModel{
			Title:       title,
			Description: description,
			Status:      status,
		}
		if err := tx.WithContext(ctx).Create(&todo).Error; err != nil {
			return err
		}
		return nil
	})
}

func (repo *ToDoRepo) GetById(ctx context.Context, id int) (*models.ToDo, error) {
	var todo = new(ToDoModel)
	err := repo.db.WithContext(ctx).
		Where("id = ?", id).
		Find(&todo).Error
	if err != nil {
		return nil, err
	}
	return &models.ToDo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}, nil
}


func (repo *ToDoRepo) GetToDoByStatus(ctx context.Context, status string) ([]models.ToDo, error) {
	var todos []ToDoModel
	err := repo.db.WithContext(ctx).Where("status = ?", status).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	todosDomain := make([]models.ToDo, len(todos))
	for i, value := range todos {
		todosDomain[i] = models.ToDo{
			ID: value.ID,
			Title: value.Title,
			Description: value.Description,
			Status: value.Status,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
	}
	return todosDomain, nil
}
