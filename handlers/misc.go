package handlers

import (
	"fmt"
	"github.com/danryan/hal"
	"time"
)

var (
	// PingHandler responds to !ping with PONG
	PingHandler = hal.Hear(`^!ping$`, func(res *hal.Response) error {
		return res.Send(fmt.Sprintf("PONG: %s\n", time.Now()))
	})
)
