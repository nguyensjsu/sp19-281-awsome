package DAO

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	. "job/models"
)


type JobDatabase struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "jobs"
)

func (j *JobDatabase) Connect() {
	session, err := mgo.Dial(j.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(j.Database)
}

// Find list of jobs
func (m *JobDatabase) FindAll() ([]Job, error) {
	var jobs []Job
	err := db.C(COLLECTION).Find(bson.M{}).All(&jobs)
	return jobs, err
}

// Create a Job
func (m *JobDatabase) CreateJob(job Job) (Job, error) {
	fmt.Print("Creating Job: ")
	fmt.Print(job)
	err := db.C(COLLECTION).Insert(job)
	return job, err
}
