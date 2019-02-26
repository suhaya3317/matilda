package domain

type MoviesAPI struct {
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

type Movie struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	PosterPath string `json:"poster_path"`
}

type MovieAPI struct {
	Adult               bool                  `json:"adult"`
	BackdropPath        string                `json:"backdrop_path"`
	BelongsToCollection BelongsToCollection   `json:"belongs_to_collection"`
	Budget              int                   `json:"budget"`
	Genres              []Genres              `json:"genres"`
	Homepage            string                `json:"homepage"`
	ID                  int                   `json:"id"`
	ImdbId              string                `json:"imdb_id"`
	OriginalLanguage    string                `json:"original_language"`
	OriginalTitle       string                `json:"original_title"`
	Overview            string                `json:"overview"`
	Popularity          float64               `json:"popularity"`
	PosterPath          string                `json:"poster_path"`
	ProductionCompanies []ProductionCompanies `json:"production_companies"`
	ProductionCountries []ProductionCountries `json:"production_countries"`
	ReleaseDate         string                `json:"release_date"`
	Revenue             int                   `json:"revenue"`
	Runtime             int                   `json:"runtime"`
	SpokenLanguages     []SpokenLanguages     `json:"spoken_languages"`
	Status              string                `json:"status"`
	Tagline             string                `json:"tagline"`
	Title               string                `json:"title"`
	Video               bool                  `json:"video"`
	VoteAverage         float64               `json:"vote_average"`
	VoteCount           int                   `json:"vote_count"`
}

type BelongsToCollection struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}

type Genres struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductionCompanies struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type ProductionCountries struct {
	Iso_3166_1 string `json:"iso_3166_1"`
	Name       string `json:"name"`
}

type SpokenLanguages struct {
	Iso_639_1 string `json:"iso_639_1"`
	Name      string `json:"name"`
}
