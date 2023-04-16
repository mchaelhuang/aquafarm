package entity

import (
	"time"

	"gorm.io/gorm"
)

type Pond struct {
	ID        int            `json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	FarmID int    `json:"farm_id"`
	Label  string `json:"label"`
	Volume int    `json:"volume"` // in cm
}

type PondFilter struct {
	Pond
	WithDeleted bool // include deleted in result
}

type PondFilterRequest struct {
	FarmID int    `query:"farm_id"`
	Label  string `query:"label"`
	Volume int    `query:"location"`
}

type PondFormRequest struct {
	FarmID int    `form:"farm_id"`
	Label  string `form:"label"`
	Volume int    `form:"volume"`
}

type PondFormResponse struct {
	ID int `json:"id"`
}
