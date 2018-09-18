package tokenize

import (
	"strings"
	"regexp"
)

// Tokenize string to uniterms. 
func ToUniterms(str string) ([]string, error) {
	patterns := []string{
		`[\n|\n\r|\r]`, 			// new lines
		" -|- ", 					// hyphen = dash
		"--",						// brackets
		`[^-A-Za-zА-Яа-я0-9Ëё]`, 	// all except letters and numbers
		` {2,}`, 					// spaces
	}

	for _, pattern := range patterns {
		str = regexp.MustCompile(pattern).ReplaceAllString(str, " ")
	}

	contractions := map[string]string{
		`(?i)\b(can|don|doesn|hasn|couldn|wouldn|wasn|won|didn|isn|shouldn|haven|weren|aren|hadn|shan|mustn|mayn|mightn) t\b`: "$1't",
		`(?i)\b(it|here|what|there|let) s\b`: "$1's",
		`(?i)\b(you|we|they) (ll|d|ve|re)\b`: "$1'$2",
		`(?i)\b(he|she) (s|ll|d)\b`: "$1'$2",
		`(?i)\b(i) (ll|d|ve|m)\b`: "$1'$2",
		`(?i)\b(that) (s|ll)\b`: "$1'$2",
	}

	for pattern, substitution := range contractions {
		str = regexp.MustCompile(pattern).ReplaceAllString(str, substitution)
	}

	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)

	return strings.Split(str, " "), nil	
}