package main

import (
	"changeme/pkg/appmod"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/daodao97/fly"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"changeme/pkg/notify"
	"changeme/pkg/stdpath"
	"changeme/pkg/utils"
	"changeme/plugin/upgrade"
	goruntime "runtime"
)

var appPath = "./"
var appName = "devbox"

func init() {
	// upgrade.Init("", "1.0")
	if !appmod.IsDev() {
		path, err := stdpath.AppDataLocation(appName)
		if err != nil {
			_ = notify.Notification(appName, "", err.Error(), "")
			return
		}
		appPath = path

		err = utils.SaveDir(data, appPath, false)
		if err != nil {
			_ = notify.Notification(appName, "", err.Error(), "")
			return
		}

	}

	err := fly.Init(map[string]*fly.Config{
		"default": {
			DSN:    filepath.Join(appPath, "data", "request.db"),
			Driver: "sqlite3",
		},
	})
	if err != nil {
		_ = notify.Notification(appName, "", err.Error(), "")
		return
	}
}

type Boot struct {
	homeDir string
	ctx     context.Context
}

func NewBoot() *Boot {
	home, _ := os.UserHomeDir()
	return &Boot{
		homeDir: home,
	}
}

func (a *Boot) onStartup(ctx context.Context) {
	a.ctx = ctx
	a.menu(ctx)
	//runtime.WindowMaximise(ctx)
}

func (a Boot) onDomReady(ctx context.Context) {
	home, _ := os.UserHomeDir()
	go upgrade.Check(ctx, home)
}

func (a *Boot) onShutdown(ctx context.Context) {

}

func (a *Boot) onBeforeClose(ctx context.Context) bool {
	return true
}

func (a *Boot) menu(ctx context.Context) {
	myMenu := menu.NewMenuFromItems(
		menu.SubMenu("File", menu.NewMenuFromItems(
			menu.Separator(),
			menu.Text("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
				os.Exit(0)
			}),
			menu.Text("Restart", keys.CmdOrCtrl("r"), func(_ *menu.CallbackData) {
				if runtime.Environment(a.ctx).Platform == goruntime.GOOS {
					restartApp(true)
				}
			}),
			menu.Text("Hide", keys.CmdOrCtrl("w"), func(_ *menu.CallbackData) {
				runtime.WindowHide(a.ctx)
			}),
		)),
		menu.EditMenu(),
	)
	runtime.MenuSetApplicationMenu(ctx, myMenu)
}

func restartApp(admin bool) {
	withAdmin := ""
	if admin {
		withAdmin = " with prompt \"开启增强模式\" with administrator privileges"
	}
	str := fmt.Sprintf(`
	do shell script "/Applications/%s.app/Contents/MacOS/%s >/dev/null 2>&1 &" %s
	`, appName, appName, withAdmin)
	cmd := exec.Command("osascript", "-e", str)
	_, err := cmd.Output()
	if err != nil {
		return
	}
	os.Exit(0)
}
