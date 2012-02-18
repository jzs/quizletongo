package quizongo

import (
	"strings"
	"strconv"
)

func (session Session) Group(group int64) (output *Group, error error) {
	output = new(Group)
	request_path := strings.Join([]string{"groups", strconv.FormatInt(group,10)}, "/")
	error = session.get(request_path, nil, output)
	return
}

func (session Session) GroupSets(set int64, params Params) (output *[]Set, error error) {
	output = new([]Set)
	request_path := strings.Join([]string{"groups", strconv.FormatInt(set,10), "sets"}, "/")
	error = session.get(request_path, params, output)
	return
}
