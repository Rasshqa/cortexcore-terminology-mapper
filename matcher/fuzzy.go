package matcher

import (
	"regexp"
	"sort"
	"strings"

	"terminology-mapper/dictionary"
)

var reNonAlpha = regexp.MustCompile(`[^a-z0-9\s]`)

// Tokenize cleans and splits a clinical description into comparable search tokens.
func Tokenize(text string) []string {
	s := strings.ToLower(text)
	s = reNonAlpha.ReplaceAllString(s, " ")
	words := strings.Fields(s)

	var tokens []string
	stopWords := map[string]bool{
		"dan": true, "atau": true, "yang": true, "dengan": true,
		"tanpa": true, "di": true, "ke": true, "dari": true,
	}

	for _, w := range words {
		if !stopWords[w] && len(w) > 1 {
			tokens = append(tokens, w)
		}
	}
	return tokens
}

// ScoreTokens calculates Jaccard similarity between query tokens and reference tokens.
func ScoreTokens(queryTokens, refTokens []string) float64 {
	if len(queryTokens) == 0 || len(refTokens) == 0 {
		return 0.0
	}

	refSet := make(map[string]bool, len(refTokens))
	for _, t := range refTokens {
		refSet[t] = true
	}

	intersection := 0
	for _, q := range queryTokens {
		if refSet[q] {
			intersection++
		}
	}

	union := len(queryTokens) + len(refTokens) - intersection
	if union == 0 {
		return 0.0
	}
	return float64(intersection) / float64(union)
}

// FindMatches ranks dictionary terms against a local hospital item.
func FindMatches(item dictionary.LocalItem, catalog []dictionary.StandardTerm, topN int) []dictionary.MappingMatch {
	if topN <= 0 {
		topN = 3
	}

	queryTokens := Tokenize(item.LocalName)
	type scoredTerm struct {
		term  dictionary.StandardTerm
		score float64
	}

	var candidates []scoredTerm

	for _, st := range catalog {
		if item.Category != "" && st.Category != item.Category {
			continue // Filter by category if specified
		}

		// Exact code match
		if strings.EqualFold(item.LocalCode, st.Code) {
			return []dictionary.MappingMatch{{
				StandardCode: st.Code,
				StandardName: st.DisplayID,
				SystemURI:    st.SystemURI,
				Category:     st.Category,
				Confidence:   100,
				Status:       "AUTO_MATCHED",
			}}
		}

		score := ScoreTokens(queryTokens, st.SearchTerms)
		if score > 0.05 {
			candidates = append(candidates, scoredTerm{term: st, score: score})
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score > candidates[j].score
	})

	var matches []dictionary.MappingMatch
	for i := 0; i < len(candidates) && i < topN; i++ {
		c := candidates[i]
		// Convert score to 0-100 percentage
		confidence := int(c.score * 100 * 1.5) // boost overlapping clinical tokens
		if confidence > 98 {
			confidence = 98
		}
		if confidence < 10 {
			confidence = 10
		}

		status := "NEEDS_REVIEW"
		if confidence >= 75 {
			status = "AUTO_MATCHED"
		} else if confidence < 40 {
			status = "UNMAPPED"
		}

		matches = append(matches, dictionary.MappingMatch{
			StandardCode: c.term.Code,
			StandardName: c.term.DisplayID,
			SystemURI:    c.term.SystemURI,
			Category:     c.term.Category,
			Confidence:   confidence,
			Status:       status,
		})
	}

	return matches
}
