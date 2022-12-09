package model

import (
	uuid "github.com/satori/go.uuid"
)

type Article struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key;"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	Title     string    `json:"title" gorm:"type:varchar(50);not null"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	HeadImage string    `json:"head_image"`
	CreatedAt Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time      `json:"updated_at" gorm:"type:timestamp"`
}

type ArticleInfo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	HeadImage string `json:"head_image"`
	CreatedAt Time   `json:"created_at"`
}

type CreateArticleRequest struct {
	Title     string `json:"title" binging:"required"`
	Content   string `json:"content" binging:"required"`
	HeadImage string `json:"head_image"`
}
