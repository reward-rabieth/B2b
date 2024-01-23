package Procurer

import (
	"context"
	"github.com/reward-rabieth/b2b/core/components/Procurer/models"
)

type Component interface {
	Create(ctx context.Context, kyc *models.Requisition) (*models.Requisition, error)
}

type component struct {
	repo Repo
}

func NewComponent(repo Repo) *component {
	return &component{
		repo: repo,
	}
}

func (c *component) Create(ctx context.Context, requisition *models.Requisition) (*models.Requisition, error) {
	createdRequisition, err := c.repo.Create(ctx, requisition)
	if err != nil {
		return nil, err
	}
	return createdRequisition, nil
}
