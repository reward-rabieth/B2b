package Procurer

import (
	"context"
	"github.com/reward-rabieth/b2b/core/components/Procurer/models"
	"gorm.io/gorm"
)

type Repo interface {
	Create(ctx context.Context, requisition *models.Requisition) (*models.Requisition, error)
	GetByID(id int) (*models.Requisition, error)
	Update(ctx context.Context, requisition *models.Requisition) error
	Delete(ctx context.Context, id int) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (r *repo, err error) {
	r = &repo{
		db: db,
	}
	return
}

func (r *repo) Create(ctx context.Context, requisition *models.Requisition) (*models.Requisition, error) {
	return nil, nil
}

func (r *repo) GetByID(id int) (*models.Requisition, error) {
	return nil, nil
}

func (r *repo) Update(ctx context.Context, requisition *models.Requisition) error {
	return nil
}

func (r *repo) Delete(ctx context.Context, id int) error {
	return nil
}
