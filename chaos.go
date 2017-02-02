package chaos

import (
	"fmt"
	"math/rand"
)

type Procedure interface {
	Name() string
	Runnable() bool
	Run() error
	Success() bool
}

var procedures = []Procedure{}

// Register a new procedure you would like to run
func Register(pr Procedure) {
	procedures = append(procedures, pr)
}

// start running procedures in a random order and run until failure
func Start() error {
	for {
		proc := procedures[rand.Intn(len(procedures))]
		if proc.Runnable() {
			fmt.Println("\n\n\nNow running:%s\n\n", proc.Name())

			err := proc.Run()
			if err != nil {
				return err
			}

			if !proc.Success() {
				return fmt.Errorf("process (%s) ran but is not successful", proc.Name())
			}
		}
	}
}
