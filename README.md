# CharactorGen
This is a character generator.Character is the one you image what you should be. It stay in your pc or mobile locally and use basic service from network.

The first function is chatroom based on a simple tcp framework.

# Install
You should install [Go](http://godoc.golangtc.com/doc/install) first and compile the src by yourself.

If you have installed Go, you can easily install the chatroom using the following commands.
```
cd /path/to/chatroom
export GOPATH=$(pwd):$GOPATH
go install example
```

# Usage
If you have installed the chatroom successfully, you will see the binary file `example` in
`/path/to/chatroom/bin`, you can use it as the following:
```
./bin/example server
#open a new bash
./bin/example client
> chatroom join 1
> chatroom send hello
```
