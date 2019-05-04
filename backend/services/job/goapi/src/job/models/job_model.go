package models
import "gopkg.in/mgo.v2/bson"

type Job struct {
	ID  bson.ObjectId `json:"id" bson:"_id"`
	Title string        `bson:"title" json:"title"`
	Company  string        `bson:"company" json:"company"`
	Description  string        `bson:"description" json:"description"`
	Industry  string        `bson:"industry" json:"industry"`
	JobType  string        `bson:"jobType" json:"jobType"`
	Function  string        `bson:"function" json:"function"`
	Logo  string        `bson:"logo" json:"logo"`

}
