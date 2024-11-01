package repositories

import (
	"keizer-auth-api/internal/models"

	"gorm.io/gorm"
)

type DomainRepository struct {
	db *gorm.DB
}

func NewDomainRepository(db *gorm.DB) *DomainRepository {
	return &DomainRepository{db: db}
}

func (r *DomainRepository) GetActiveDomain(
	origin string,
) (*models.Domain, error) {
	var session models.Domain
	err := r.db.First(
		&session,
		"origin = ?",
		origin,
	).Error
	if err != nil {
		return nil, err
	}

	return &session, nil
}
