package models

type CreatePost struct {
	Title    string `json:"title" binding:"required"`
	Category string `json:"category" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Image    string `json:"image" binding:"required"`
}
