package domain

import (
	"fmt"
	"strings"
	"unicode"
)

// fieldLimit is the default length cap for free-form text the schema declares
// with a finite size (VARCHAR columns).
const fieldLimit = 160

// rejectInvalidText validates a user-provided text against the rules the
// workshop accepts: no control characters, a length within the schema column
// and, when structured is true, no angle brackets that could carry markup.
func rejectInvalidText(value, field string, limit int, structured bool) error {
	if hasControlCharacter(value) {
		return fmt.Errorf("%w: %s contains control characters", ErrInvalidInput, field)
	}
	if len(value) > limit {
		return fmt.Errorf("%w: %s exceeds the maximum length of %d", ErrInvalidInput, field, limit)
	}
	if structured && (strings.Contains(value, "<") || strings.Contains(value, ">")) {
		return fmt.Errorf("%w: %s contains markup characters", ErrInvalidInput, field)
	}
	return nil
}

// rejectControlCharacter refuses a text column (TEXT) that carries a control
// character. Markup is allowed in free prose but control characters never are.
func rejectControlCharacter(value, field string) error {
	if hasControlCharacter(value) {
		return fmt.Errorf("%w: %s contains control characters", ErrInvalidInput, field)
	}
	return nil
}

func hasControlCharacter(value string) bool {
	for _, r := range value {
		if unicode.IsControl(r) {
			return true
		}
	}
	return false
}
