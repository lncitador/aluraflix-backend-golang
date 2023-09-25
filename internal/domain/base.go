package domain

import (
	"github.com/go-playground/validator/v10"
	v "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"time"
)

var validate = validator.New()

type Base struct {
	ID        *v.UniqueEntityID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time         `gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt time.Time         `gorm:"type:timestamp;not null;autoUpdateTime"`
}

func (b *Base) prepare() {
	now := time.Now()

	b.ID, _ = v.NewUniqueEntityID(nil)

	b.CreatedAt = now
	b.UpdatedAt = now
}
