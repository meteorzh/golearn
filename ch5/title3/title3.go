package title3

import (
	"fmt"

	"golang.org/x/net/html"
)

// soleTitle 返回文档中第一个非空标题元素
// 如果没有标题则返回错误
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// 没有宕机
		case bailout{}:
			// "预期的"宕机
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // 未预期的宕机，继续宕机过程
		}
	}()

	// 如果发现多余一个非空标题，退出递归
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // 多个标题元素
			}
			title = n.FirstChild.Data
		}
	}, nil)

	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
