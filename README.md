# CharactorGen
This is a character generator.Character is the one you image what you should be. It stay in your pc or mobile locally and use basic service from network.

The first function is chatroom based on a simple tcp framework.

# Recent activity
Use new protocol, eg. `chatroom.send|msg:haha`
2015-11-6 21:41
Use the new feature of [hazel](http://github.com/sctlee/hazel): Delegation
Use Message to send msg to clients, including boardcast.

More details seen in [hazel](http://github.com/sctlee/hazel)

# Install
#### Go installation
You should install [Go](http://golang.org) first.
- Ubuntu & Debian based Linux
```shell
sudo apt-get install go-lang
```

#### Dependence
This project is based on [hazel](http://github.com/sctlee/hazel) and [utils](http://github.com/sctlee/utils). You can get it by the command:
```shell
go get github.com/sctlee/hazel
go get github.com/sctlee/utils
```
Be careful, You may want to `cd /path/to/example` and `export GOPATH=$(pwd):$GOPATH`, the exec the above commands.

#### Example installation
If you have installed Go and dependences, you can easily install the example using the following commands.
```shell
cd /path/to/example
export GOPATH=$(pwd):$GOPATH
go install example
```

# Usage
#### Quick Start
If you have installed the example successfully, you will see the binary file `example` in
`/path/to/example/bin`, you can use it as the following:
```shell
./bin/example server
# open a new bash
./bin/example client
> chatroom.join|ctName:1
> chatroom.send|msg:hello
```

#### Chatroom
Chatroom feature has three command: list, join, send
```
> chatroom.list|
> chatroom.join|ctName:1
> chatroom.send|msg:hello
```

#### User
User feature has three command: login, logout, setName
```
> user.login|username:***;password:***
> user.setName|name:hc
> user.logout|
```

Because db has not been supported, user's name is fake. It means you can't sign up and the name you set can't be saved.
