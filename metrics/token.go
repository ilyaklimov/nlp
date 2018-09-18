package metrics

import (
	"errors"
)

// TF calculate the frequency (without normalizing) of token in tokens.
// token = (uniterms | biterms | n-terms ...) || (unirunes | birunes | n-runes ...)
func TF(token string, tokens []string) (float64, error) {
    if token == "" {
        return 0.0, errors.New("token is empty")
    }
    if len(tokens) == 0 {
        return 0.0, errors.New("token list is empty")
    }
    match := 0
    for _, value := range tokens {
    	if token == value {
    		match++
    	}
    }
    return (float64(match) / float64(len(tokens))), nil
}