package handlers

import (
	"fmt"
	"github.com/danryan/hal"
	"github.com/davecgh/go-spew/spew"
)

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
	return res.Send(fmt.Sprintf("get: %s=%s", key, string(val)))
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
	return res.Send(fmt.Sprintf("set: %s=%s", key, val))
})

// DeleteHandler nukes keys in the Store
var DeleteHandler = hal.Hear(`^forget (.+)$`, func(res *hal.Response) error {
	key := res.Match[0][1]

	if err := res.Robot.Store.Delete(key); err != nil {
		res.Send(err.Error())
		return err
	}
	return res.Send(fmt.Sprintf("delete: %s", key))
})

// UsersHandler provides insite into the Robots user storage
var UsersHandler = hal.Hear(`^!show users$`, func(res *hal.Response) error {
	line := spew.Sdump("%#v\n", res.Robot.Users)
	return res.Send(line)
})
