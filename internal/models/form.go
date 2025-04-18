package models

import "gorm.io/gorm"

type (
	Form struct {
		gorm.Model
		Title  string  `json:"title"`
		Fields []Field `json:"fields"`
		FormId string  `json:"form_id"`
		UserId int     `json:"user_id"`
	}

	Field struct {
		Name   string `bson:"name" json:"name" validate:"required"`
		Type   string `bson:"type" json:"type" validate:"required"`
		FormId uint   `bson:"type" json:"form_id" validate:"required"`
	}

	FieldResponse struct {
		Name string `json:"name" validate:"required"`
		Type string `json:"type" validate:"required"`
	}

	FormResponse struct {
		Title  string          `json:"title"`
		Fields []FieldResponse `json:"fields"`
	}
)
