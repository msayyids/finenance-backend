package model

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetCategoriesByUserResponse struct {
	UserID     int                `json:"user_id"`
	Categories []CategoryResponse `json:"categories"`
}

type CategoryRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
