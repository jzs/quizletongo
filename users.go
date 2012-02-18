package quizongo

import (
	"strings"
)

func (session Session) User(username string) (output *User, error error) {
	output = new(User)
	request_path := strings.Join([]string{"users", username}, "/")
	error = session.get(request_path, nil, output)
	return
}

func (session Session) UserSets(username string, params Params) (output *[]Set, error error) {
	output = new([]Set)
	request_path := strings.Join([]string{"users", username, "sets"}, "/")
	error = session.get(request_path, params, output)
	return
}

func (session Session) UserFavorites(username string, params Params) (output *[]Set, error error) {
	output = new([]Set)
	request_path := strings.Join([]string{"users", username, "favorites"}, "/")
	error = session.get(request_path, params, output)
	return
}

func (session Session) UserGroups(username string) (output *[]Group, error error) {
	output = new([]Group)
	request_path := strings.Join([]string{"users", username, "groups"}, "/")
	error = session.get(request_path, nil, output)
	return
}

func (session Session) UserStudied(username string) (output *[]StudiedSession, error error) {
	output = new([]StudiedSession)
	request_path := strings.Join([]string{"users", username, "studied"}, "/")
	error = session.get(request_path, nil, output)
	return
}