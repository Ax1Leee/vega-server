package api

type Review struct {
	Rating  float32 `json:"rating"`
	Content string  `json:"content"`
}

type GetReviewFromUserRequest struct {
	ID uint `form:"reviewID" binding:"required"`
}

type GetReviewFromUserResponseData struct {
	Movie  Movie  `json:"movie"`
	Review Review `json:"review"`
}

type GetReviewToMovieRequest struct {
	ID uint `form:"reviewID" binding:"required"`
}

type GetReviewToMovieResponseData struct {
	User   User   `json:"user"`
	Review Review `json:"review"`
}

type GetReviewResponseData struct {
	Rating  float32 `json:"rating"`
	Content string  `json:"content"`
}

type SetReviewRequest struct {
	Rating  float32 `json:"rating" binding:"required"`
	Content string  `json:"content" binding:"required"`
}

type SetReviewResponseData struct {
	Rating  float32 `json:"rating"`
	Content string  `json:"content"`
}
