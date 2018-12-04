// fetch输出从url获取的内容
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	httpPrefix = "http://"
)

func main() {
	for _, url := range os.Args[1:] {
		// 练习：自动判断是否有协议前缀
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", b)

		// 练习：输出状态码
		fmt.Println(resp.Status)
		// 练习：使用io.Copy将数据复制到os.Stdout
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %v\n", err)
			os.Exit(1)
		}
	}
}
