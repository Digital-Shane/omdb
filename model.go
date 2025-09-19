package omdb

// QueryData is the type to create a search query.
type QueryData struct {
	Title      string
	Year       string
	ImdbID     string
	SearchType string
	Plot       string
	Page       string
	Season     string
	Episode    string
}

// resultEnvelope will be used to unmarshall API response for checking Type.
// Based on Type, the response can be unmarshalled to MovieResult/SeriesResult/
// EpisodeResult structs.
type resultEnvelope struct {
	Type     string
	Response string
	Error    string
}

// MovieResult will hold information of a single movie.
type MovieResult struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []Rating
	Metascore  string
	ImdbRating string
	ImdbVotes  string
	ImdbID     string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
}

// SeriesResult will hold information of a single series.
type SeriesResult struct {
	Title        string
	Year         string
	Rated        string
	Released     string
	Runtime      string
	Genre        string
	Director     string
	Writer       string
	Actors       string
	Plot         string
	Language     string
	Country      string
	Awards       string
	Poster       string
	Ratings      []Rating
	Metascore    string
	ImdbRating   string
	ImdbVotes    string
	ImdbID       string
	TotalSeasons string
}

// EpisodeResult will hold information of a single episode.
type EpisodeResult struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []Rating
	Metascore  string
	ImdbRating string
	ImdbVotes  string
	ImdbID     string
	SeriesID   string
	Season     string
	Episode    string
}

// Rating will hold rating information from a single source.
type Rating struct {
	Source string
	Value  string
}

// SearchResponse is a container holding one or more SearchResults.
type SearchResponse struct {
	Search       []SearchResult
	TotalResults string
	Response     string
	Error        string
}

// SearchResult represents a single result from API search by text.
type SearchResult struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
	Poster string
}

// SeasonResult represents a season listing containing multiple episodes.
type SeasonResult struct {
	Title        string          `json:"Title"`
	Season       string          `json:"Season"`
	TotalSeasons string          `json:"totalSeasons"`
	Episodes     []SeasonEpisode `json:"Episodes"`
	Response     string          `json:"Response"`
	Error        string          `json:"Error"`
}

// SeasonEpisode contains summary information for an episode returned in a season listing.
type SeasonEpisode struct {
	Title      string `json:"Title"`
	Released   string `json:"Released"`
	Episode    string `json:"Episode"`
	ImdbID     string `json:"imdbID"`
	ImdbRating string `json:"imdbRating"`
}
