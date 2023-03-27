package wallpapers

import (
	"fmt"
	"regexp"

	"github.com/charmbracelet/log"

	scriptish "github.com/ganbarodigital/go_scriptish"
)

func waifus() {
	checkPath, _ := scriptish.NewPipeline(
		scriptish.Exec("gsettings", "get", "org.gnome.shell.extensions.WallpaperSwitcher", "wallpaper-path"),
	).Exec().TrimmedString()

	alreadyOn, err := regexp.MatchString(`(waifu)`, checkPath)
	if err != nil {
		log.Warn(err)
	}

	if alreadyOn == true {
		scriptish.NewPipeline(
			scriptish.Exec("gsettings", "set", "org.gnome.shell.extensions.WallpaperSwitcher", "wallpaper-path", "/usr/share/backgrounds/wallpapers/"),
		).Exec()
		fmt.Println("Disabled waifus")
	} else {
		scriptish.NewPipeline(
			scriptish.Exec("gsettings", "set", "org.gnome.shell.extensions.WallpaperSwitcher", "wallpaper-path", "~/pictures/wallpapers/waifus"),
		).Exec()
		fmt.Println("Enabled waifus")
	}
}
