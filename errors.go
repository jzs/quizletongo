package quizongo

import (
	"fmt"
)

type QuizletError struct {
	http_code int
	error string
	error_title string
	error_description string
}

func (e QuizletError) Error() string {
	return fmt.Sprintf("%v", e.error_title)
}
