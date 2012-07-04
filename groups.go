package quizongo

import (
	"errors"
	"strconv"
	"strings"
)

// Optional Parameters
//     modified_since int	; Limits results to sets that have been modified since the specified Unix timestamp.
func (session Session) Group(group int64) (output *Group, error error) {
	output = new(Group)
	request_path := strings.Join([]string{"groups", strconv.FormatInt(group, 10)}, "/")
	error = session.get(request_path, nil, output)
	return
}

// Optional Parameters
//     modified_since int	; Limits results to sets that have been modified since the specified Unix timestamp.
func (session Session) GroupSets(set int64, params Params) (output *[]Set, error error) {
	output = new([]Set)
	request_path := strings.Join([]string{"groups", strconv.FormatInt(set, 10), "sets"}, "/")
	error = session.get(request_path, params, output)
	return
}

// Optional Parameters:
//    is_public boolean		; 0 or 1. Set this to 1 if you want this group to allow anyone to join.
//    password string		; You can optionally set a password on private sets that allows users to join. This field is ignored if the group is public.
//    allow_discussion		; 0 or 1. Flag of whether users are allowed to use the discussion box on this group.
//    allow_member_add_sets	; 0 or 1. Set this flag to 0 to only allow managers to add sets.
func (session Session) AddGroup(name string, description string, params Params) (output *Group, error error) {
	error = errors.New("AddGroup not implemented yet.")
	return
}

// Optional Parameters:
//     name string
//     description string
//     is_public boolean
//     password string
//     allow_discussion boolean
//     allow_member_add_sets boolean
func (session Session) EditGroup(params Params) (error error) {
	error = errors.New("EditGroup not implemented yet.")
	return
}
