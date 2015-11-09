package auth

import "github.com/sctlee/tcpx"

func Permission(client *tcpx.Client, f tcpx.RouteFun) tcpx.RouteFun {
	return f
}
