package domain

import (
	"github.com/go-playground/validator/v10"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"github.com/lncitador/alura-flix-backend/pkg/errors"
)

type Video struct {
	Base
	Title       string             `gorm:"type:varchar(255);not null"`
	Description string             `gorm:"type:varchar(255);not null"`
	URL         *vo.URL            `gorm:"type:varchar(255);not null;index"`
	CategoryID  *vo.UniqueEntityID `gorm:"type:uuid, not null;index"`
	Category    *Categoria         `gorm:"foreignKey:CategoryID"`
	UsuarioID   *vo.UniqueEntityID `gorm:"type:uuid, not null;index"`
	Usuario     *Usuario           `gorm:"foreignKey:UsuarioID"`
}

// NewVideo creates a new Video instance
func NewVideo(input VideoInput) (*Video, errors.Error) {
	video := Video{}
	video.prepare()

	if err := input.validate(); err != nil {
		return nil, errors.NewErrorByValidation(err)
	}

	newUrl, err := vo.NewURL(*input.URL)
	if err != nil {
		return nil, errors.NewErrorByValidation(err)
	}

	video.Title = *input.Title
	video.Description = *input.Description
	video.URL = newUrl
	video.CategoryID, err = vo.NewUniqueEntityID(input.CategoryID)
	if err != nil {
		return nil, errors.NewErrorByValidation(err)
	}
	video.UsuarioID, err = vo.NewUniqueEntityID(input.UsuarioID)
	if err != nil {
		return nil, errors.NewErrorByValidation(err)
	}

	return &video, nil
}

// MapTo maps Video to VideoDto struct
func (v *Video) MapTo() *VideoDto {
	video := &VideoDto{
		ID:          v.ID.ToString(),
		Title:       v.Title,
		Description: v.Description,
		URL:         v.URL.ToString(),
		CategoryID:  v.CategoryID.ToString(),
		UsuarioID:   v.UsuarioID.ToString(),
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}

	if v.Category != nil {
		video.Category = v.Category
	}

	if v.Usuario != nil {
		video.Usuario = v.Usuario
	}

	return video
}

// Fill fills Video with VideoInput data
func (v *Video) Fill(input VideoInput) errors.Error {
	err := validate.Struct(input)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				continue
			}
			return errors.NewErrorByValidation(err)
		}
	}

	if input.Title != nil {
		v.Title = *input.Title
	}

	if input.Description != nil {
		v.Description = *input.Description
	}

	if input.URL != nil {
		newUrl, err := vo.NewURL(*input.URL)
		if err != nil {
			return errors.NewErrorByValidation(err)
		}
		v.URL = newUrl
	}

	if input.CategoryID != nil {
		categoriaId, err := vo.NewUniqueEntityID(input.CategoryID)
		if err != nil {
			return errors.NewErrorByValidation(err)
		}
		v.CategoryID = categoriaId
	}

	return nil
}
