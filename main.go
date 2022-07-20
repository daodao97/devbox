package main

import (
	"changeme/pkg/appmod"
	"embed"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"

	"changeme/app"
	request "changeme/plugin/http"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed data
var data embed.FS

var log = logger.NewDefaultLogger()

func fileServer(w http.ResponseWriter, request *http.Request) {
	filePath := filepath.Join(appPath, "data", strings.ReplaceAll(strings.Split(request.RequestURI, "?")[0], "wails://wails/", ""))
	bt, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error(fmt.Sprintf("read plugin file %s fail %v", request.RequestURI, err))
		return
	}

	fileExtension := filepath.Ext(filePath)
	w.Header().Add("Content-Type", mime.TypeByExtension(fileExtension))

	_, err = w.Write(bt)
	if err != nil {
		log.Error(fmt.Sprintf("response plugin file %s fail %v", request.RequestURI, err))
		return
	}
}

func main() {
	boot := NewBoot()
	newApp := app.NewApp()

	log.Info(fmt.Sprintf("app mod %v", appmod.IsDev()))

	err := wails.Run(&options.App{
		Width:         1024,
		Height:        768,
		Assets:        assets,
		AssetsHandler: http.HandlerFunc(fileServer),
		OnStartup:     boot.onStartup,
		OnDomReady:    boot.onDomReady,
		OnShutdown:    boot.onShutdown,
		// OnBeforeClose: boot.onBeforeClose,
		Bind: []interface{}{
			newApp,
			&request.Http{},
		},
		Mac: &mac.Options{
			TitleBar: mac.TitleBarHidden(),
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
