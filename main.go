package main

import (
	"embed"
	"os/exec"
	"time"

	"github.com/getlantern/systray"
)

const (
	ICON_CHANGE_TIMEOUT = 2
)

func main() {
	systray.Run(onReady, onExit)
}

//go:embed assets/*
var assets embed.FS

func onReady() {

	emptyIcon, _ := assets.ReadFile("assets/trash_empty.ico") // lucide.dev/icons/trash
	fullIcon, _ := assets.ReadFile("assets/trash_full.ico")

	systray.SetTitle("Mini Bin")
	systray.SetTooltip("Mini Bin")

	go func() {
		for {
			empty, _ := IsRecycleBinEmpty()
			if empty {
				systray.SetIcon(emptyIcon)
			} else {
				systray.SetIcon(fullIcon)
			}

			time.Sleep(time.Second * ICON_CHANGE_TIMEOUT)
		}
	}()

	mOpen := systray.AddMenuItem("Open", "Open Bin folder")
	mEmpty := systray.AddMenuItem("Empty", "Empty Bin")
	mQuit := systray.AddMenuItem("Quit", "Quit app")

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	go func() {
		<-mOpen.ClickedCh
		exec.Command("explorer", "shell:RecycleBinFolder").Start()
	}()

	go func() {
		<-mEmpty.ClickedCh
		emptyBin()
	}()
}

func onExit() {}
