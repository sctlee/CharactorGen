package task

import (
	"strings"

	"github.com/sctlee/hazel"
)

func Route(url string, client *hazel.Client) {
	var action string
	url = strings.TrimSpace(url)
	i := strings.Index(url, " ")
	if i == -1 {
		action = url[:]
	} else {
		action = url[:i]
	}
	switch action {
	case "publish":
	case "accept":
	case "build":
	}
}
