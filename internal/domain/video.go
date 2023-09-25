package domain

import (
	"github.com/go-playground/validator/v10"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"time"
)

type Video struct {
	Base
	Title       string  `gorm:"type:varchar(255);not null"`
	Description string  `gorm:"type:varchar(255);not null"`
	URL         *vo.URL `gorm:"type:varchar(255);not null"`
}

type VideoDto struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type VideoInput struct {
	Title       *string `json:"title" validate:"required,min=8,max=255"`
	Description *string `json:"description" validate:"required,min=8,max=255"`
	URL         *string `json:"url" validate:"required,url"`
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

	return &video, nil
}

// MapTo maps Video to VideoDto struct
func (v *Video) MapTo() *VideoDto {
	return &VideoDto{
		ID:          v.ID.ToString(),
		Title:       v.Title,
		Description: v.Description,
		URL:         v.URL.ToString(),
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
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

	return &Video{base, d.Title, d.Description, newUrl}, nil
}
