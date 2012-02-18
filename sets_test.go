package quizongo

import (
	"testing"
)

func TestSets(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/sets", dummySetsResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("apikey")
	sets, err := session.Sets(nil, 415,6009523)
	if err != nil {
		t.Error(err.Error())
	}
	if sets == nil {
		t.Error("set is nil")
		return
	}
	if len(*sets) != 2 {
		t.Error("Sets not right size")
	}
}

func TestSet(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/sets/6009523", dummySetResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("apikey")
	set, err := session.Set(6009523, nil)
	if err != nil {
		t.Error(err.Error())
	}
	if set == nil {
		t.Error("set is nil")
		return
	}
	if set.Id != 415 {
		t.Error("Set Id does not match.")
	}


}

func TestSetTerms(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/sets/6009523/terms", dummyTermsResponse, t)
	defer closeDummyServer(dummy_server)

	session := NewSession("apikey")
	terms, err := session.SetTerms(6009523, nil)
	if err != nil {
		t.Error(err.Error())
	}
	if terms == nil {
		t.Error("set is nil")
		return
	}
	if len(*terms) != 1{
		t.Error("Term length is wrong.")
	}
	if (*terms)[0].Id != 15407 {
		t.Error("Id does not match.")
	}
	if (*terms)[0].Term != "Montgomery" {
		t.Error("Term does not match.")
	}
	if (*terms)[0].Definition != "Alabama" {
		t.Error("Definition does not match.")
	}
}

var dummyTermsResponse = `
[
  {
    "id": 15407,
    "term": "Montgomery",
    "definition": "Alabama",
    "image": null
  }
]
`

var dummySetResponse = `
{
  "id": 415,
  "url": "http:\/\/quizlet.com\/415\/four-one-five-flash-cards\/",
  "title": "four one five",
  "created_by": "jalenack",
  "term_count": 50,
  "created_date": 1144296408,
  "modified_date": 1316216163,
  "has_images": false,
  "subjects": [
    "unitedstates",
    "states",
    "geography",
    "capitals"
  ],
  "visibility": "only_me",
  "editable": "only_me",
  "has_access": true,
  "has_discussion": true,
  "description": "List of the U.S. states and their capitals",
  "terms": [
    {
      "id": 15407,
      "term": "Montgomery",
      "definition": "Alabama",
      "image": null
    }
  ]
}
`

var dummySetsResponse = `
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
	  ],
	  "visibility": "public",
	  "editable": "only_me",
	  "has_access": true,
	  "has_discussion": true,
	  "description": "",
	  "terms": [
	    {
	      "id": 190713940,
	      "term": "une grenouille",
	      "definition": "a frog",
	      "image": {
	        "url": "http:\/\/farm1.static.flickr.com\/74\/179791578_70992b1821.jpg",
	        "width": 180,
	        "height": 240
	      }
	    }
	  ]
	},
	{
	  "id": 415,
	  "url": "http:\/\/quizlet.com\/415\/four-one-five-flash-cards\/",
	  "title": "four one five",
	  "created_by": "jalenack",
	  "term_count": 50,
	  "created_date": 1144296408,
	  "modified_date": 1316216163,
	  "has_images": false,
	  "subjects": [
	    "unitedstates",
	    "states",
	    "geography",
	    "capitals"
	  ],
	  "visibility": "only_me",
	  "editable": "only_me",
	  "has_access": true,
	  "has_discussion": true,
	  "description": "List of the U.S. states and their capitals",
	  "terms": [
	    {
	      "id": 15407,
	      "term": "Montgomery",
	      "definition": "Alabama",
	      "image": null
	    }
	  ]
	}
]
`
