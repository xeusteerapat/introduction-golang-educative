package main

import (
	"log"
	"os"
)

type Job struct {
	Command     string
	*log.Logger // you don't need Logger *log.Logger
}

func main() {
	job := &Job{"demo", log.New(os.Stdout, "Job: ", log.Ldate)}
	// same as
	// job := &Job{Command: "demo",
	//            Logger: log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Print("test")
}
