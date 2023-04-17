package wallpapers

import (
	"fmt"
	"regexp"
	"runtime"

	"github.com/charmbracelet/log"
	"github.com/tr00datp00nar/fn"
)

func waifus() {
	runtimeOs := runtime.GOOS
	if runtimeOs == "linux" {
		checkPath, _, _ := fn.ExecBash("gsettings get org.gnome.shell.extensions.WallpaperSwitcher 'wallpaper-path'")

		alreadyOn, err := regexp.MatchString(`(waifus)`, checkPath)
		if err != nil {
			log.Warn(err)
		}

		if alreadyOn == true {
			fn.ExecBash("gsettings set org.gnome.shell.extensions.WallpaperSwitcher wallpaper-path '/usr/share/backgrounds/wallpapers/'")
			fmt.Println("Disabled waifus")
		} else {
			fn.ExecBash("gsettings set org.gnome.shell.extensions.WallpaperSwitcher wallpaper-path '~/pictures/wallpapers/waifus'")
			fmt.Println("Enabled waifus")
		}
	} else {
		log.Fatal("Unsupported OS...Exiting...")
	}
}
