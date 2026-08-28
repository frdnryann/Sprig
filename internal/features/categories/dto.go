package categories

type CreateCategoryRequest struct {
	UserID uint64  `json:"user_id"`
	Name   *string `json:"name"`
}

type CategoryResponse struct {
	ID     uint64 `json:"id"`
	UserID uint64 `json:"user_id"`
	Name   string `json:"name"`
}
