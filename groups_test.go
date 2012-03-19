package quizongo

import (
	"testing"
)

func TestGroup(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/groups/102269", dummyGroupResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("apikey")
	group, err := session.Group(102269)
	if err != nil {
		t.Error(err.Error())
	}
	if group == nil {
		t.Error("group is nil")
		return
	}
}

func TestGroupSets(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/groups/102269/sets", dummyGroupSetsResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("apikey")
	sets, err := session.GroupSets(102269, nil)
	if err != nil {
		t.Error(err.Error())
	}
	if sets == nil {
		t.Error("sets is nil")
		return
	}
}

/*
func TestAddGroup(t *testing.T) {
	return
}

func TestEditGroup(t *testing.T) {
	return
}

func TestDeleteGroup(t *testing.T) {
	return
}

func TestJoinGroup(t *testing.T) {
	return
}

func TestLeaveGroup(t *testing.T) {
	return
}

func TestAddSetToGroup(t *testing.T) {
	return
}

func TestRemoveSetFromGroup(t *testing.T) {
	return
}
*/


var dummyGroupSetsResponse = `
[
    {
      "id": 6081999,
      "url": "http:\/\/quizlet.com\/6081999\/lesson-4-with-catsdogs-flash-cards\/",
      "title": "Lesson 4 (with cats+dogs)",
      "created_by": "wsvincent",
      "term_count": 33,
      "created_date": 1313790421,
      "modified_date": 1314044315,
      "has_images": true,
      "subjects": [
        "spanish",
        "cats",
        "dogs"
      ]
    }
  ]
`

var dummyGroupResponse = `
{
  "id": 102269,
  "name": "Learn Spanish with Cats!",
  "set_count": 3,
  "user_count": 8,
  "created_date": 1314037860,
  "is_public": true,
  "has_password": false,
  "member_add_sets": false,
  "description": "This set is exclusively for Spanish flashcard sets with relevant cat images as the set definitions.",
  "sets": [
    {
      "id": 6081999,
      "url": "http:\/\/quizlet.com\/6081999\/lesson-4-with-catsdogs-flash-cards\/",
      "title": "Lesson 4 (with cats+dogs)",
      "created_by": "wsvincent",
      "term_count": 33,
      "created_date": 1313790421,
      "modified_date": 1314044315,
      "has_images": true,
      "subjects": [
        "spanish",
        "cats",
        "dogs"
      ]
    }
  ],
  "members": [
    {
      "username": "philfreo",
      "role": "creator",
      "email_notification": true
    },
    {
      "username": "jeffchan",
      "role": "member",
      "email_notification": true
    },
    {
      "username": "jalenack",
      "role": "member",
      "email_notification": true
    }
  ]
}
`
