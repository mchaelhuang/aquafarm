package entity

import (
	"time"

	"gorm.io/gorm"
)

type Farm struct {
	ID        int            `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Name     string `json:"name"`
	Location string `json:"location"`
	Ponds    []Pond `json:"ponds"`
}

type FarmFilter struct {
	Farm
	WithDeleted bool // include deleted in result
}

type FarmFilterRequest struct {
	Name     string `query:"name"`
	Location string `query:"location"`
}

type FarmFormRequest struct {
	Name     string `form:"name"`
	Location string `form:"location"`
}

type FarmFormResponse struct {
	ID int `json:"id"`
}
