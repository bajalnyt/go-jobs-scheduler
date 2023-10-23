package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	IN_PROGRESS = "in_progress"
	COMPLETED   = "completed"
	FAILED      = "failed"
	ABORTED     = "aborted"
)

type job struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

var jobs = []job{
	{ID: "1", Name: "Job1", Status: IN_PROGRESS},
	{ID: "2", Name: "Job2", Status: IN_PROGRESS},
}

func getJobs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jobs)
}

func runJob(jobId string) {
	time.Sleep(8 * time.Second)

	// Printed after sleep is over
	fmt.Printf("Job %s completed", jobId)

}

// postJob adds a job.
func postJob(c *gin.Context) {
	var newJob job
	// Always set a new job to IN_PROGRESS
	newJob.Status = IN_PROGRESS

	if err := c.BindJSON(&newJob); err != nil {
		return
	}

	// Add the new album to the slice.
	jobs = append(jobs, newJob)
	c.IndentedJSON(http.StatusCreated, jobs)
	go runJob(newJob.ID)
}

// getJobById
func getJobById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of jobs, looking for
	// a job whose ID value matches the parameter.
	for _, a := range jobs {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "job not found"})
}

// udpateJobStatusByID
func udpateJobStatusByID(c *gin.Context) {
	id := c.Param("id")
	var updatedJob job

	if err := c.BindJSON(&updatedJob); err != nil {
		return
	}
	// Loop over the list of jobs, looking for
	// a job whose ID value matches the parameter.
	for i, a := range jobs {
		if a.ID == id {
			// replace the element to delete with the one at the end of the slice
			jobs[i].Status = updatedJob.Status
			c.IndentedJSON(http.StatusCreated, jobs)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "job not found"})
}

func main() {
	router := gin.Default()
	router.GET("/jobs", getJobs)
	router.POST("/jobs", postJob)
	router.GET("/job/:id", getJobById)
	router.PUT("/job/:id", udpateJobStatusByID)

	router.Run("localhost:8080")
}
