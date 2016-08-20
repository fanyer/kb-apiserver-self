package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// The main structure for database

type Study struct {
	Id            bson.ObjectId `json:"id" bson:"_id,omitempty"` // Mongodb _id
	Title         string        `json:"Title,omitempty" bson:"Title,omitempty"`
	UniversalID   string        `json:"UniversalID,omitempty" bson:"UniversalID,omitempty"`
	Abstract      string        `json:"Abstract,omitempty" bson:"Abstract,omitempty"`
	StudyType     string        `json:"StudyType,omitempty" bson:"StudyType,omitempty"`
	Description   string        `json:"Description,omitempty" bson:"Description,omitempty"`
	CenterName    string        `json:"CenterName,omitempty" bson:"CenterName,omitempty"`
	Organism      string        `json:"Organism,omitempty" bson:"Organism,omitempty"`
	Owner         string        `json:"Owner,omitempty" bson:"Owner,omitempty"`
	DeleteFlag    bool          `json:"DeleteFlag,omitempty" bson:"DeleteFlag,omitempty"`
	CreateTime    time.Time     `json:"CreateTime,omitempty" bson:"CreateTime,omitempty"`
	CreateTimeRaw string        `json:"CreateTimeRaw,omitempty" bson:"CreateTimeRaw,omitempty"`
}

func main() {
	// Home page route
	http.HandleFunc("/", simpleHandler)
	// Route for API /studys
	http.HandleFunc("/kbp/", apiHandler)

	bind := fmt.Sprintf("%s:%s", "localhost", 8081)
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		panic(err)
	}
}

func apiHandler(res http.ResponseWriter, req *http.Request) {

	// Mongo db configurations
	session, err := mgo.Dial("localhost:3001")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Db and collection name configs
	c := session.DB("knowledgebase").C("study")

	var result []Study
	// Get id form URL
	id := strings.Replace(req.URL.Path, "/kbp/", "", -1)

	//set mime type to JSON, Its JSON REST API
	res.Header().Set("Content-type", "application/json")

	// Handle the methods and behave accordingly
	switch req.Method {
	case "GET":
		// If no id passed in url, show them all out Pandas

		if id != "" {

			err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).All(&result)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err = c.Find(nil).All(&result)
			fmt.Println(result[0])

		}

	// case "POST":
	// 	// Read POST body from request
	// 	body, err := ioutil.ReadAll(req.Body)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// 	// Convert body json to struct data
	// 	var study Study
	// 	err = json.Unmarshal(body, &study)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// 	// We need a new mongodb _id to insert record, We are doing this becuase mongodb doesnt return last inserted record info
	// 	i := bson.NewObjectId()
	// 	study.Id = i
	//
	// 	// Insert study into farm.studys
	// 	err = c.Insert(study)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	// Get details about just inserted row
	// 	err = c.Find(bson.M{"_id": i}).All(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// case "PUT":
	// 	// Read POST body
	// 	body, err := ioutil.ReadAll(req.Body)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// 	// Convert request json to struct
	// 	var study Study
	// 	err = json.Unmarshal(body, &study)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	// We need a new mongoDb _id
	// 	i := bson.ObjectIdHex(id)
	// 	study.Id = i
	//
	// 	// Update
	// 	err = c.Update(bson.M{"_id": i}, study)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	// Get info about just inserted document
	// 	err = c.Find(bson.M{"_id": i}).All(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	case "DELETE":
		// When a study leaves :(, Delete from database
		err = c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
		if err != nil {
			log.Fatal(err)
		}
	default:
	}

	// Convert result struct to JSON
	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the result JSON to the client.
	fmt.Fprintf(res, "%v", string(json))
}
func simpleHandler(res http.ResponseWriter, req *http.Request) {
	//set mime type to HTML
	res.Header().Set("Content-type", "text/html")
	// Guide them
	fmt.Fprintf(res, "Sir you are at wrong place!<br />Pandas are at <a href='/kbp/'>/studys</a>")
}
