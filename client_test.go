package omdb

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func newTestClient(handler roundTripFunc) *http.Client {
	return &http.Client{Transport: handler}
}

func jsonResponse(body string) *http.Response {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     make(http.Header),
	}
	resp.Header.Set("Content-Type", "application/json")
	return resp
}

func TestSearchByTitleSeasonRequest(t *testing.T) {
	t.Parallel()

	var query string
	client := NewClient("apikey", newTestClient(func(req *http.Request) (*http.Response, error) {
		query = req.URL.RawQuery
		return jsonResponse(`{"Title":"Game of Thrones","Season":"1","totalSeasons":"8","Episodes":[{"Title":"Winter Is Coming","Released":"2011-04-17","Episode":"1","imdbID":"tt1480055","imdbRating":"8.9"}],"Response":"True"}`), nil
	}))

	result, err := client.SearchByTitle(QueryData{Title: "Game of Thrones", Season: "1"})
	if err != nil {
		t.Fatalf("SearchByTitle() error = %v", err)
	}

	if !strings.Contains(query, "Season=1") {
		t.Fatalf("query = %q, want to contain Season=1", query)
	}

	season, ok := result.(SeasonResult)
	if !ok {
		t.Fatalf("result type = %T, want SeasonResult", result)
	}

	want := SeasonResult{
		Title:        "Game of Thrones",
		Season:       "1",
		TotalSeasons: "8",
		Episodes: []SeasonEpisode{{
			Title:      "Winter Is Coming",
			Released:   "2011-04-17",
			Episode:    "1",
			ImdbID:     "tt1480055",
			ImdbRating: "8.9",
		}},
		Response: "True",
	}

	if diff := cmp.Diff(want, season); diff != "" {
		t.Errorf("SearchByTitle() (-want +got)\n%s", diff)
	}
}

func TestSearchByTitleEpisodeRequest(t *testing.T) {
	t.Parallel()

	var query string
	client := NewClient("apikey", newTestClient(func(req *http.Request) (*http.Response, error) {
		query = req.URL.RawQuery
		return jsonResponse(`{"Title":"Winter Is Coming","Year":"2011","Season":"1","Episode":"1","Runtime":"62 min","Genre":"Drama","imdbRating":"8.9","imdbID":"tt1480055","seriesID":"tt0944947","Type":"episode","Response":"True"}`), nil
	}))

	result, err := client.SearchByTitle(QueryData{Title: "Game of Thrones", Season: "1", Episode: "1", Plot: "full"})
	if err != nil {
		t.Fatalf("SearchByTitle() error = %v", err)
	}

	if !strings.Contains(query, "Season=1") || !strings.Contains(query, "Episode=1") || !strings.Contains(query, "plot=full") {
		t.Fatalf("query = %q, want to contain Season=1, Episode=1, plot=full", query)
	}

	episode, ok := result.(EpisodeResult)
	if !ok {
		t.Fatalf("result type = %T, want EpisodeResult", result)
	}

	if episode.ImdbID != "tt1480055" {
		t.Fatalf("EpisodeResult.ImdbID = %q, want tt1480055", episode.ImdbID)
	}
}

func TestSearchByTitleEpisodeRequiresSeason(t *testing.T) {
	t.Parallel()

	client := NewClient("apikey", newTestClient(func(req *http.Request) (*http.Response, error) {
		return jsonResponse(`{"Response":"False","Error":"Movie not found!"}`), nil
	}))

	_, err := client.SearchByTitle(QueryData{Title: "Game of Thrones", Episode: "1"})
	if err == nil {
		t.Fatalf("SearchByTitle() error = nil, want error")
	}
}
