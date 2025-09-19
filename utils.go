package omdb

import (
	"strconv"
	"strings"
)

// ParseRating converts an OMDb rating string (e.g. "8.6") into a float. Returns 0 on failure.
func ParseRating(raw string) float32 {
	raw = strings.TrimSpace(raw)
	if raw == "" || strings.EqualFold(raw, "N/A") {
		return 0
	}
	value, err := strconv.ParseFloat(raw, 32)
	if err != nil {
		return 0
	}
	return float32(value)
}

// SplitAndTrim splits a comma separated string and removes empty or "N/A" items.
func SplitAndTrim(raw string) []string {
	if raw == "" || strings.EqualFold(raw, "N/A") {
		return nil
	}
	parts := strings.Split(raw, ",")
	cleaned := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" || strings.EqualFold(part, "N/A") {
			continue
		}
		cleaned = append(cleaned, part)
	}
	if len(cleaned) == 0 {
		return nil
	}
	return cleaned
}

// FirstYear extracts the first four digit year from a string (e.g. "2014-2016").
func FirstYear(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}
	var digits strings.Builder
	for i := 0; i < len(raw); i++ {
		ch := raw[i]
		if ch >= '0' && ch <= '9' {
			digits.WriteByte(ch)
			if digits.Len() == 4 {
				break
			}
		} else if digits.Len() > 0 {
			break
		}
	}
	return digits.String()
}

// FirstYearFromEpisodes finds the first year among a slice of SeasonEpisode release dates.
func FirstYearFromEpisodes(episodes []SeasonEpisode) string {
	for _, ep := range episodes {
		if year := FirstYear(ep.Released); year != "" {
			return year
		}
	}
	return ""
}
