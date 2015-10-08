# CharactorGen
This is a character generator.Character is the one you image what you should be. It stay in your pc or mobile locally and use basic service from network.

The first function is chatroom based on a simple tcp framework.

# Install
#### Go installation
You should install [Go](http://golang.org) first.
- Ubuntu & Debian based Linux
```shell
sudo apt-get install go-lang
```

#### Chatroom installation
If you have installed Go, you can easily install the chatroom using the following commands.
```shell
cd /path/to/chatroom
export GOPATH=$GOPATH:`pwd`
go install example
```

# Usage
If you have installed the chatroom successfully, you will see the binary file `example` in
`/path/to/chatroom/bin`, you can use it as the following:
```shell
./bin/example server
# open a new bash
./bin/example client
> chatroom join 1
> chatroom send hello
```
