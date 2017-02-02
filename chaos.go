package chaos

import (

)

type (
	Procedure interface{
		Name() string
		Runnable() bool
		Run() error
		Success() bool
	}
)

var procedures = []Procedure{}

// Register a new procedure you would like to run
func Register(proc Procedure) {
	procedures = append(Procedure, proc)
}

// start running procedures in a random order and run until failure
func Start() error {
	for {
		for _, proc := range procedures {
			if proc.Runnable() {
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
}

