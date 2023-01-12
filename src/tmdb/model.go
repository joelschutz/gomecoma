package tmdb

type ResponseSearch struct {
	Page         int64        `json:"page"`
	Results      []ResultBase `json:"results"`
	TotalPages   int64        `json:"total_pages:`
	TotalResults int64        `json:"total_results:`
}

type Genre struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ResultBase struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Popularity       float64 `json:"popularity"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int64   `json:"vote_count"`
	Id               int64   `json:"id"`
	GenreIds         []int64 `json:"genre_ids"`
	Video            bool    `json:"video"`
}

type ResultFull struct {
	ResultBase
	Runtime int64   `json:"runtime"`
	Genres  []Genre `json:"genres"`
}
