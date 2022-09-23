package schemas

type CreatePost struct {
	Title    string `json:"title" binding:"required"`
	Category string `json:"category" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Image    string `json:"image" binding:"required"`
}

type UpdatePost struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Content  string `json:"content"`
	Image    string `json:"image"`
}
