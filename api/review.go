package api

type Review struct {
	Rating    float32 `json:"rating"`
	Content   string  `json:"content"`
	UpdatedAt string  `json:"updatedAt"`
}

type GetReviewFromUserRequest struct {
	ID uint `form:"reviewID" binding:"required"`
}

type GetReviewToMovieRequest struct {
	ID uint `form:"reviewID" binding:"required"`
}

type GetReviewRequest struct {
	MovieID uint `form:"movieID" binding:"required"`
}

type SetReviewRequest struct {
	Rating  float32 `json:"rating" binding:"required"`
	Content string  `json:"content" binding:"required"`
}

type GetReviewFromUserResponseData struct {
	Movie  Movie  `json:"movie"`
	Review Review `json:"review"`
}

type GetReviewToMovieResponseData struct {
	User   User   `json:"user"`
	Review Review `json:"review"`
}

type GetReviewResponseData struct {
	Rating  float32 `json:"rating"`
	Content string  `json:"content"`
}

type SetReviewResponseData struct {
	Rating  float32 `json:"rating"`
	Content string  `json:"content"`
}
