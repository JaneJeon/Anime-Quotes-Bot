package quote

import "time"

type Provider interface {
	GetRandomQuote() (Quote, error)
	refresh(string) error
}

type quotes struct {
	topics []Quote
	used []Quote
	unused []Quote
	checkedAt time.Time
}

type Quote struct {
	Content string
	Author string
}

func appendSet(needle Quote, haystack []Quote) []Quote {
	for _, q := range haystack {
		if needle == q {
			return haystack
		}
	}
	
	return append(haystack, needle)
}