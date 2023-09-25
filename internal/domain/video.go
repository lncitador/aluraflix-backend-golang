package domain

import (
	"github.com/go-playground/validator/v10"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"time"
)

type Video struct {
	Base
	Title       string             `gorm:"type:varchar(255);not null"`
	Description string             `gorm:"type:varchar(255);not null"`
	URL         *vo.URL            `gorm:"type:varchar(255);not null;index"`
	CategoriaID *vo.UniqueEntityID `gorm:"type:uuid, not null;index"`
	Categoria   *Categoria         `gorm:"foreignKey:CategoriaID"`
}

type VideoDto struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	URL         string     `json:"url"`
	CategoriaID string     `json:"categoriaId"`
	Categoria   *Categoria `json:"categoria,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type VideoInput struct {
	Title       *string `json:"title" validate:"required,min=8,max=255"`
	Description *string `json:"description" validate:"required,min=8,max=255"`
	URL         *string `json:"url" validate:"required,url"`
	CategoriaID *string `json:"categoriaId" validate:"required,uuid4"`
}

func (i VideoInput) validate() error {
	if err := validate.Struct(i); err != nil {
		return err
	}

	return nil
}

// NewVideo creates a new Video instance
func NewVideo(input VideoInput) (*Video, error) {
	video := Video{}
	video.prepare()

	if err := input.validate(); err != nil {
		return nil, err
	}

	newUrl, err := vo.NewURL(*input.URL)
	if err != nil {
		return nil, err
	}

	video.Title = *input.Title
	video.Description = *input.Description
	video.URL = newUrl
	video.CategoriaID, err = vo.NewUniqueEntityID(input.CategoriaID)
	if err != nil {
		return nil, err
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
		CategoriaID: v.CategoriaID.ToString(),
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}

	if v.Categoria != nil {
		video.Categoria = v.Categoria
	}

	return video
}

func (v *Video) Fill(input VideoInput) error {
	err := validate.Struct(input)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				continue
			}
			return err
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
			return err
		}
		v.URL = newUrl
	}

	if input.CategoriaID != nil {
		categoriaId, err := vo.NewUniqueEntityID(input.CategoriaID)
		if err != nil {
			return err
		}
		v.CategoriaID = categoriaId
	}

	return nil
}

func VideosToDto(videos []Video) *[]VideoDto {
	var videosDto []VideoDto

	for _, video := range videos {
		videosDto = append(videosDto, *video.MapTo())
	}

	return &videosDto
}

// MapFrom maps VideoDto to Video struct
func (d VideoDto) MapFrom() (*Video, error) {
	id, err := vo.NewUniqueEntityID(&d.ID)
	if err != nil {
		return nil, err
	}

	newUrl, err := vo.NewURL(d.URL)
	if err != nil {
		return nil, err
	}

	base := Base{
		ID:        id,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}

	categoriaId, err := vo.NewUniqueEntityID(&d.CategoriaID)
	if err != nil {
		return nil, err
	}

	return &Video{
		Base:        base,
		Title:       d.Title,
		Description: d.Description,
		URL:         newUrl,
		CategoriaID: categoriaId,
		Categoria:   d.Categoria,
	}, nil
}
