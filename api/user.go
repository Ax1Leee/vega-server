package api

type User struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

type AdvancedUser struct {
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthDate"`
	Location  string `json:"location"`
	Bio       string `json:"bio"`
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignUpResponseData struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignInResponseData struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type GetUserResponseData struct {
	User User `json:"user"`
}

type GetAdvancedUserResponseData struct {
	User    AdvancedUser `json:"user"`
	Reviews []uint       `json:"reviews"`
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
