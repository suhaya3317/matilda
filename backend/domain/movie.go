package domain

type Movies struct {
	Page         int       `json:"page"`
	TotalResults int       `json:"total_results"`
	TotalPages   int       `json:"total_pages"`
	Results      []Results `json:"results"`
}

type Results struct {
	VoteCount        int     `json:"vote_count"`
	ID               int     `json:"id"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	Title            string  `json:"title"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	GenreIDs         []int   `json:"genre_ids"`
	BackdropPath     string  `json:"backdrop_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
}

type MatildaMovie struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	PosterPath string `json:"poster_path"`
}
