package upgrade

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/codeskyblue/go-sh"
	"github.com/spf13/cast"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// like https://github.com/daodao97/EasyClash/releases/download/v1.0/EasyClash-amd-v1.0.zip
var _downloadUrlTpl = ""

// Major.Minor.Patch
// 主版本.次版本.修订号
var _currentVersion = "1.0"

func Init(downloadUrl, currentVersion string) {
	_downloadUrlTpl = downloadUrl
	_currentVersion = currentVersion
}

func getNeedCheckVersion() [][]interface{} {
	tmp := func() []int {
		var t []int
		for _, v := range strings.Split(_currentVersion, ".") {
			t = append(t, cast.ToInt(v))
		}
		return t
	}()

	toInterface := func(arr []int) []interface{} {
		var _arr []interface{}
		for _, v := range arr {
			_arr = append(_arr, v)
		}
		return _arr
	}

	var needCheck [][]interface{}

	for i := len(tmp) - 1; i >= 0; i-- {
		t := make([]int, len(tmp))
		copy(t, tmp)
		t[i] = t[i] + 1
		needCheck = append(needCheck, toInterface(t))
	}

	return needCheck
}

type versionInfo struct {
	Url  string
	Size int
}

func CheckNewVersion() (*versionInfo, error) {
	if _currentVersion == "" || _downloadUrlTpl == "" {
		return nil, errors.New("upgrade not init")
	}

	checkList := getNeedCheckVersion()
	for _, v := range checkList {
		url := fmt.Sprintf(_downloadUrlTpl, v...)
		resp, err := http.Head(url)
		if err != nil {
			return nil, err
		}

		headerResp, err := http.Head(url)
		if err != nil {
			return nil, err
		}
		defer headerResp.Body.Close()
		size := cast.ToInt(headerResp.Header.Get("Content-Length"))

		if resp.StatusCode == http.StatusOK {
			return &versionInfo{
				Url:  url,
				Size: size,
			}, nil
		}
	}

	return nil, errors.New("no version")

}

type Process struct {
	Size     int64
	Complate int64
	FileName string
	Error    error
}

func Download(path string, url string) chan Process {
	client := grab.NewClient()
	req, _ := grab.NewRequest(path, url)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	process := make(chan Process)

	go func() {
		t := time.NewTicker(500 * time.Millisecond)
		defer func() {
			t.Stop()
			close(process)
		}()
	Loop:
		for {
			select {
			case <-t.C:
				process <- Process{
					Size:     resp.Size(),
					Complate: resp.BytesComplete(),
				}
			case <-resp.Done:
				process <- Process{
					Size:     resp.Size(),
					Complate: resp.Size(),
					FileName: resp.Filename,
				}
				break Loop
			}
		}

		if err := resp.Err(); err != nil {
			process <- Process{
				Error: fmt.Errorf("Download failed: %v\n", err),
			}
		}
	}()

	return process
}

func Check(ctx context.Context, homeDir string) {
	downUrl, err := CheckNewVersion()
	if err != nil {
		return
	}
	runtime.EventsEmit(ctx, "new_version_download", true)
	runtime.EventsOn(ctx, "new_version_download_action", func(optionalData ...interface{}) {
		process := Download(filepath.Join(homeDir, "Downloads"), downUrl.Url)
		done := false
		file := ""
		for v := range process {
			runtime.EventsEmit(ctx, "new_version_download_process", v)
			if v.Complate == v.Size {
				done = true
				file = v.FileName
			}
		}
		if done {
			_ = sh.Command("open", file).Run()
		}
	})
}
