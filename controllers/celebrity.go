package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api.biomatch.com/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// CelebrityController represents the controller for operating on the Celebrity resource
	CelebrityController struct {
		session *mgo.Session
	}
)

// NewCelebrityController provides a reference to a CelebrityController with provided mongo session
func NewCelebrityController(s *mgo.Session) *CelebrityController {
	return &CelebrityController{s}
}

// GetCelebrity retrieves an individual user resource
func (cc CelebrityController) GetCelebrity(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	u := models.Celebrity{}

	// Fetch user
	if err := cc.session.DB("go_rest_tutorial").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateCelebrity creates a new user resource
func (cc CelebrityController) CreateCelebrity(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.Celebrity{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	// Add an Id
	u.Id = bson.NewObjectId()

	// Write the user to mongo
	cc.session.DB("go_rest_tutorial").C("users").Insert(u)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveCelebrity removes an existing user resource
func (cc CelebrityController) RemoveCelebrity(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := cc.session.DB("go_rest_tutorial").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}
