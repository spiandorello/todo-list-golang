package dtos

type TaskCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
