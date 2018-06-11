package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/rishavs/buffapi/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
