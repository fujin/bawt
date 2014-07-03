package main

import (
	"github.com/danryan/hal"
	"log"
	"os"
)

var pingHandler = hal.Hear(`ping`, func(res *hal.Response) error {
	return res.Send("PONG")
})

var openDoorsHandler = hal.Respond(`open the pod bay doors`, func(res *hal.Response) error {
	return res.Reply("I'm sorry, Dave. I can't do that.")
})

func main() {
	os.Exit(Run())
}

// Run launches the bot, returns standard exit code
func Run() int {
	robot, newErr := hal.NewRobot()
	if newErr != nil {
		log.Println(newErr)
		return 1
	}

	robot.Handle(
		pingHandler,
		openDoorsHandler,
	)

	runErr := robot.Run()
	if runErr != nil {
		log.Println(runErr)
		return 1
	}
	return 0
}
