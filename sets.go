package quizongo

import (
	"strings"
	"strconv"
	"errors"
)

// Optional Parameters:
//     modified_since int64
func (session Session) Sets(params Params, sets ... int64) (output *[]Set, error error) {
	output = new([]Set)
	error = session.get("sets", params, output)
	return
}

// Optional Parameters:
//     modified_since int64
func (session Session) Set(set int64, params Params) (output *Set, error error) {
	output = new(Set)
	request_path := strings.Join([]string{"sets", strconv.FormatInt(set,10)}, "/")
	error = session.get(request_path, params, output)
	return
}

func (session Session) SetTerms(set int64, params Params) (output *[]Term, error error) {
	output = new([]Term)
	request_path := strings.Join([]string{"sets", strconv.FormatInt(set,10), "terms"}, "/")
	error = session.get(request_path, params, output)
	return
}

func (session Session) AddSet(title string, terms []Term, definitions []string, lang_terms string, lang_definitions string, params Params) (output *SetSuccess, error error) {
	error = errors.New("AddSet not implemented yet.")
	return
}

func (session Session) EditSet() {}
func (session Session) AddSetTerm() {}
func (session Session) EditSetTerm() {}
func (session Session) DeleteSetTerm() {}
