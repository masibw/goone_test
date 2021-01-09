package db

import (
	"database/sql"
	"log"
)

type Person struct {
	Name  string
	JobID int
}

type Job struct {
	JobID int
	Name  string
}

func GetJob2(cnn *sql.DB,person Person) Job {
	var job Job
	// This is N+1 Query
	err := cnn.QueryRow("SELECT job_id, name FROM Jobs WHERE job_id = ?", person.JobID).Scan(&job.JobID, &job.Name)
	if err != nil {
		log.Fatal(err)
	}
	return job
}
