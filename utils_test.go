package omdb

import "testing"

func TestParseRating(t *testing.T) {
	t.Parallel()

	if got := ParseRating("8.6"); got != 8.6 {
		t.Fatalf("ParseRating(\"8.6\") = %v, want 8.6", got)
	}

	if got := ParseRating("N/A"); got != 0 {
		t.Fatalf("ParseRating(\"N/A\") = %v, want 0", got)
	}
}

func TestSplitAndTrim(t *testing.T) {
	t.Parallel()

	got := SplitAndTrim("Action, Drama, N/A")
	want := []string{"Action", "Drama"}
	if len(got) != len(want) {
		t.Fatalf("SplitAndTrim length = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("SplitAndTrim[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestFirstYear(t *testing.T) {
	t.Parallel()

	if got := FirstYear("2014-2016"); got != "2014" {
		t.Fatalf("FirstYear(2014-2016) = %q, want 2014", got)
	}
}

func TestFirstYearFromEpisodes(t *testing.T) {
	t.Parallel()

	episodes := []SeasonEpisode{
		{Released: "N/A"},
		{Released: "2011-04-17"},
	}
	if got := FirstYearFromEpisodes(episodes); got != "2011" {
		t.Fatalf("FirstYearFromEpisodes() = %q, want 2011", got)
	}
}
