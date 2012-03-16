package quizongo

import (
	"testing"
)

func TestUsername(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/catchdave", dummyUserResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("apikey")
	user, err := session.User("catchdave")

	if err != nil {
		t.Error(err.Error())
	}
	if user == nil {
		t.Error("user is nil")
		return
	}
	if len(user.Sets) != 1 {
		t.Error("Number of Sets wrong.")
	}
	if len(user.Favorite_sets) != 1 {
		t.Error("Number of Favorites wrong.")
	}
	if len(user.Groups) != 1 {
		t.Error("Number of Groups wrong.")
	}
	if user.Sets[0].Id != 6009523 {
		t.Error("Set Id Not Valid")
	}
	if user.Username != "catchdave" {
		t.Error("Username does not match")
	}
	if user.Groups[0].Id != 103561 {
		t.Error("Group Id wrong")
	}
}

func TestUserSets(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/catchdave/sets", dummyUserSetsResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("apikey")
	sets, err := session.UserSets("catchdave", nil)

	if err != nil {
		t.Error(err.Error())
	}
	if sets == nil {
		t.Error("sets is nil")
		return
	}
	if len(*sets) != 1 {
		t.Error("User Set length wrong")
	}
	if (*sets)[0].Id != 6009523 {
		t.Error("Id don't match.")
	}
	if (*sets)[0].Subjects[1] != "animals" {
		t.Error("Subject 1 don't match.")
	}
}

func TestUserFavorites(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/catchdave/favorites", dummyUserSetsResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("apikey")
	favorites, err := session.UserFavorites("catchdave", nil)

	if err != nil {
		t.Error(err.Error())
	}
	if favorites == nil {
		t.Error("favorites is nil")
		return
	}
	if len(*favorites) != 1 {
		t.Error("Length don't match.")
	}

}

func TestUserGroups(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/catchdave/groups", dummyUserSetsResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("apikey")
	groups, err := session.UserGroups("catchdave")

	if err != nil {
		t.Error(err.Error())
		return
	}
	if groups == nil {
		t.Error("groups is nil")
		return
	}
	if len(*groups) != 1 {
		t.Error("Length don't match.")
	}

}

func TestUserStudied(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/users/catchdave/studied", dummyUserSetsResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("apikey")
	studied, err := session.UserStudied("catchdave")

	if err != nil {
		t.Error(err.Error())
	}
	if studied == nil {
		t.Error("studied is nil")
		return
	}
	if len(*studied) != 1 {
		t.Error("Length don't match")
	}

}

func TestUserMarkFavorite(t *testing.T) {
	t.Error("UserMarkFavorite not implemented yet.")
}

func TestUserUnmarkFavorite(t *testing.T) {
	t.Error("UserUnmarkFavorite not implemented yet.")
}

var dummyUserSetsResponse = `
[
    {
	  "id": 6009523,
	  "url": "http:\/\/quizlet.com\/6009523\/french-animals-30-flash-cards\/",
	  "title": "french animals 3.0",
	  "created_by": "catchdave",
	  "term_count": 13,
	  "created_date": 1310497102,
	  "modified_date": 1310497363,
	  "has_images": true,
	  "subjects": [
	    "french",
	    "animals"
	  ]
 	}
  ]
`

var dummyUserResponse = `
{
  "username": "catchdave",
  "account_type": "plus",
  "sign_up_date": 1300591868,
  "statistics": {
    "study_session_count": 6,
    "message_count": 1,
    "total_answer_count": 15,
    "public_sets_created": 19,
    "public_terms_entered": 59,
    "total_sets_created": 19,
    "total_terms_entered": 59
  },
  "sets": [
    {
	  "id": 6009523,
	  "url": "http:\/\/quizlet.com\/6009523\/french-animals-30-flash-cards\/",
	  "title": "french animals 3.0",
	  "created_by": "catchdave",
	  "term_count": 13,
	  "created_date": 1310497102,
	  "modified_date": 1310497363,
	  "has_images": true,
	  "subjects": [
	    "french",
	    "animals"
	  ]
 	}
  ],
  "favorite_sets": [
    {
	  "id": 6009523,
	  "url": "http:\/\/quizlet.com\/6009523\/french-animals-30-flash-cards\/",
	  "title": "french animals 3.0",
	  "created_by": "catchdave",
	  "term_count": 13,
	  "created_date": 1310497102,
	  "modified_date": 1310497363,
	  "has_images": true,
	  "subjects": [
	    "french",
	    "animals"
	  ]
  	}
  ],
  "studied": [
    {
      "mode": "speller",
      "last_studied_date": 1308719871,
      "finished_date": 0,
      "score": 1674,
      "set": {
        "id": 415,
        "url": "http:\/\/quizlet.com\/415\/us-state-capitals-flash-cards\/",
        "title": "U.S. State Capitals",
        "created_by": "jalenack",
        "term_count": 50,
        "created_date": 1144296408,
        "modified_date": 1308066716,
        "has_images": false,
        "subjects": [
          "unitedstates",
          "states",
          "geography",
          "capitals"
        ]
      }
    }
  ],
  "groups": [
    {
      "id": 103561,
      "name": "¡Old School learns español!",
      "set_count": 1,
      "user_count": 8,
      "created_date": 1310425248,
      "is_public": true
    }
  ]
}
`

var dummyUserStudiedResponse = `
[
    {
      "mode": "speller",
      "last_studied_date": 1308719871,
      "finished_date": 0,
      "score": 1674,
      "set": {
        "id": 415,
        "url": "http:\/\/quizlet.com\/415\/us-state-capitals-flash-cards\/",
        "title": "U.S. State Capitals",
        "created_by": "jalenack",
        "term_count": 50,
        "created_date": 1144296408,
        "modified_date": 1308066716,
        "has_images": false,
        "subjects": [
          "unitedstates",
          "states",
          "geography",
          "capitals"
        ]
      }
    }
  ]
`

var dummyUserGroupsResponse = `
[
    {
      "id": 103561,
      "name": "¡Old School learns español!",
      "set_count": 1,
      "user_count": 8,
      "created_date": 1310425248,
      "is_public": true
    }
  ]
`
