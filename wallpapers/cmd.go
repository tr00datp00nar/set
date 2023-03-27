package wallpapers

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:      `wallpaper`,
	Usage:     `[help]`,
	Version:   `v0.1.0`,
	Copyright: `Copyright Micah Nadler 2023`,
	License:   `Apache-2.0`,
	Summary:   `Tools for managing desktop backgrounds.`,
	Commands:  []*Z.Cmd{help.Cmd, toggleWaifus},
}

var toggleWaifus = &Z.Cmd{
	Name:      `waifus`,
	Usage:     `[help]`,
	Version:   `v0.1.0`,
	Copyright: `Copyright Micah Nadler 2023`,
	License:   `Apache-2.0`,
	Summary:   `Toggles the waifu desktop backgrounds.`,
	Commands: []*Z.Cmd{
		help.Cmd,
	},
	Description: `
		Checks for the current 'wallpaper-path' of the Gnome Wallpaper Switcher Extension
		and toggles between '~/pictures/wallpapers/waifus' and '/usr/share/backgrounds/wallpapers'.
	`,

	Call: func(_ *Z.Cmd, _ ...string) error {
		waifus()
		return nil
	}}
