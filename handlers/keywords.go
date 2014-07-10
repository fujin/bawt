package handlers

import (
	"github.com/danryan/hal"
	"github.com/davecgh/go-spew/spew"
)

const okStr = "OK"

// GetHandler pulls a key out of the store, when addressed like so:
// foobar?
// bawt> get: foobar=baz
var GetHandler = hal.Hear(`^(.+)?$`, func(res *hal.Response) error {
	key := res.Match[0][1]
	val, err := res.Robot.Store.Get(key)
	if err != nil {
		res.Send(err.Error())
		return err
	}
	return res.Send(string(val))
})

// SetHandler does a simple match/set into the Store
var SetHandler = hal.Hear(`^(.+) is (.+)$`, func(res *hal.Response) error {
	key := res.Match[0][1]
	val := res.Match[0][2]
	err := res.Robot.Store.Set(key, []byte(val))
	if err != nil {
		res.Send(err.Error())
		return err
	}
	return res.Send(okStr)
})

// DeleteHandler nukes keys in the Store
var DeleteHandler = hal.Hear(`^forget (.+)$`, func(res *hal.Response) error {
	key := res.Match[0][1]

	if err := res.Robot.Store.Delete(key); err != nil {
		res.Send(err.Error())
		return err
	}
	return res.Send(okStr)
})

// UsersHandler provides insite into the Robots user storage
var UsersHandler = hal.Hear(`^!show users$`, func(res *hal.Response) error {
	lines := []string{}
	for _, user := range res.Robot.Users.All() {
		lines = append(lines, spew.Sdump(user))
	}
	return res.Send(lines...)
})

// UserHandler is the singular form of UsersHandler
var UserHandler = hal.Hear(`^!show user (.+)$`, func(res *hal.Response) error {
	id := res.Match[0][1]
	user, _ := res.Robot.Users.Get(id)
	line := spew.Sdump(user)
	return res.Send(line)
})
