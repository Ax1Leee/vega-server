package api

type Movie struct {
	Cover        string  `json:"cover"`
	Title        string  `json:"title"`
	CriticRating float32 `json:"criticRating"`
	UserRating   float32 `json:"userRating"`
}

type AdvancedMovie struct {
	Cover        string   `json:"cover"`
	Title        string   `json:"title"`
	Genres       []string `json:"genres"`
	ReleaseDate  string   `json:"releaseDate"`
	Location     string   `json:"location"`
	Director     string   `json:"director"`
	Stars        []string `json:"stars"`
	Language     string   `json:"language"`
	Runtime      string   `json:"runtime"`
	Storyline    string   `json:"storyline"`
	CriticRating float32  `json:"criticRating"`
	UserRating   float32  `json:"userRating"`
}

type GetMovieRequest struct {
	ID uint `form:"movieID" binding:"required"`
}

type GetAdvancedMovieRequest struct {
	ID uint `form:"movieID" binding:"required"`
}

type GetMoviesRequest struct {
	Genre    string `form:"genre" binding:"required"`
	Category string `form:"category" binding:"required"`
	Field    string `form:"field"`
}

type GetMovieResponseData struct {
	Movie Movie `json:"movie"`
}

type GetAdvancedMovieResponseData struct {
	Movie   AdvancedMovie `json:"movie"`
	Reviews []uint        `json:"reviews"`
}

type GetMoviesResponseData struct {
	Movies []uint `json:"movies"`
}

type GetAdvancedMoviesResponseData struct {
	Titles []string `json:"titles"`
}

type GetHotMoviesResponseData struct {
	HotMovies []uint `json:"hotMovies"`
}
