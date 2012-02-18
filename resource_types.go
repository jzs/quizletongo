package quizongo

type Set struct {
	Id uint64
	Title string
	Url string
	Created_by string
	Created_date uint64
	Modified_date uint64
	Term_count int
	Has_images bool
	Visibility string
	Editable string
	Has_access bool
	Subjects []string
	Description string
	Has_discussion bool
	Lang_terms string
	Lang_definitions string
	Terms []Term
}

type Term struct {
	Id uint64
	Term string
	Definition string
	Image *Image
}

type Image struct {
	Url string
	Width int
	Height int
}

type Group struct {
	Id uint64
}

type User struct {
	Username string
	Account_type string
	Sign_up_date uint64
	Statistics Statistics
	Sets []Set
	Favorite_sets []Set
	Studied []StudiedSession
	Groups []Group
}

type Statistics struct {
	Study_session_count int
	Message_count int
	Total_answer_count int
	Public_sets_created int
	Public_terms_entered int
	Total_sets_created int
	Total_terms_entered int
}

type StudiedSession struct {
	Mode string
	Start_date uint64
	Finish_date uint64
	Formatted_score string
	Set Set
}