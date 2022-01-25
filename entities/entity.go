package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (entity *Entity) BeforeCreate(scope *gorm.DB) error {
	id, err := uuid.NewUUID()
	scope.Statement.SetColumn("Id", id)
	return err
}
