package models

import (
	"gorm.io/gorm"
)

type (
	FormRequest struct {
		gorm.Model
		Title             string  `json:"title"`
		Fields            []Field `json:"fields"`
		ParticipantsLimit int     `json:"participants_limit"`
		AccountLimit      int     `json:"account_limit"`
		UserId            int     `json:"user_id"`
	}

	Form struct {
		gorm.Model
		Title             string `json:"title"`
		PublicId          string `json:"public_id"` // Visible id in type of hashed string
		ParticipantsCount int    `json:"participants_count"`
		ParticipantsLimit int    `json:"participants_limit"`
		AccountLimit      int    `json:"account_limit"`
		UserId            int    `json:"user_id"`
	}

	FormResponse struct {
		ID           uint            `json:"id"`
		Title        string          `json:"title"`
		IsFull       bool            `json:"is_full"`
		AccountLimit int             `json:"account_limit"`
		Fields       []FieldResponse `json:"fields"`
		Varieties    []Variety       `json:"varieties"`
		Layout       LayoutResponse  `json:"layout"`
	}

	MyFormResponse struct {
		ID           uint            `json:"id"`
		Title        string          `json:"title"`
		IsFull       bool            `json:"is_full"`
		AccountLimit int             `json:"account_limit"`
		Fields       []FieldResponse `json:"fields"`
		Varieties    []Variety       `json:"varieties"`
		Layout       LayoutResponse  `json:"layout"`
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
)
