
package models


type Job struct {
	title string        `bson:"title" json:"title"`
	company  string        `bson:"company" json:"company"`
	description  string        `bson:"description" json:"description"`
	industry  string        `bson:"industry" json:"industry"`
	jobType  string        `bson:"jobType" json:"jobType"`
	function  string        `bson:"function" json:"function"`
	logo  string        `bson:"logo" json:"logo"`

}
