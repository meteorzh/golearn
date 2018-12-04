package fetch

import (
	"io"
	"net/http"
	"os"
	"path"
)

// Fetch 下载 URL 并返回本地文件的名字和长度
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// 关闭文件，并保留错误信息
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
