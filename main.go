package main

import (
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/irc"
	"github.com/fujin/bawt/handlers"
	"log"
	"os"
)

// Bawt is an infobot for #go-chef/#chef, etc
type Bawt struct {
	*hal.Robot
}

// NewBawt makes bawt or returns an error
func NewBawt() (bawt Bawt, err error) {
	bawt.Robot, err = hal.NewRobot()
	if err != nil {
		log.Fatal(err)
	}
	return bawt, nil
}

func main() {
	bawt, err := NewBawt()
	if err != nil {
		log.Fatal("unexpected error during bawt creation:", err)
	}

	bawt.Handle(
		handlers.PingHandler,
		handlers.GetHandler,
		handlers.SetHandler,
		handlers.DeleteHandler,
		handlers.UsersHandler,
	)
	bawt.Run()

	os.Exit(0)
}
