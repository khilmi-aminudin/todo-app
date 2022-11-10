package web

type ActivityCreateRequest struct {
	Email string `json:"email" binding:"required,email"`
	Title string `json:"title" binding:"required,min=5"`
}

type ActivityUpdateRequest struct {
	ID    int64  `json:"id"`
	Email string `json:"email" binding:"required,email"`
	Title string `json:"title" binding:"required,min=5"`
}
