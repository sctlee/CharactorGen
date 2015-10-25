package task

import (
	"strings"

	"github.com/sctlee/tcpx"
)

func Publish(client *tcpx.Client, paramString string) {
	params := strings.Fields(paramString)
	if len(params) != 1 {
		client.PutOutgoing("You can only input one param")
		return
	}

}

func Build(client *tcpx.Client, paramString string) {
	params := strings.Fields(paramString)
	if len(params) != 1 {
		client.PutOutgoing("You can only input one param")
		return
	}

}

func Accept(client *tcpx.Client, paramString string) {
	params := strings.Fields(paramString)
	if len(params) != 1 {
		client.PutOutgoing("You can only input one param")
		return
	}

}
