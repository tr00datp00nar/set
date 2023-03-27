package set

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/tr00datp00nar/set/wallpapers"
)

var Cmd = &Z.Cmd{
	Name:        `set`,
	Usage:       `COMMAND`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     `Tool to set various system settings.`,
	Description: ``,
	Commands:    []*Z.Cmd{help.Cmd, wallpapers.Cmd},
}
