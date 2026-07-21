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

	emptyIcon, err := assets.ReadFile("assets/trash_empty.ico") // lucide.dev/icons/trash
	if err != nil {
		panic(err)
	}
	fullIcon, err := assets.ReadFile("assets/trash_full.ico")
	if err != nil {
		panic(err)
	}

	systray.SetTitle("Mini Bin")
	systray.SetTooltip("Mini Bin")

	go func() {
		for {
			empty, err := IsRecycleBinEmpty()
			if err == nil {
				if empty {
					systray.SetIcon(emptyIcon)
				} else {
					systray.SetIcon(fullIcon)
				}
			}
			time.Sleep(time.Second * ICON_CHANGE_TIMEOUT)
		}
	}()

	mOpen := systray.AddMenuItem("Open", "Open Bin folder")
	mEmpty := systray.AddMenuItem("Empty", "Empty Bin")
	mQuit := systray.AddMenuItem("Quit", "Quit app")

	go func() {
		for {
			select {
			case <-mOpen.ClickedCh:
				exec.Command("explorer", "shell:RecycleBinFolder").Start()
			case <-mEmpty.ClickedCh:
				emptyBin()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {}
