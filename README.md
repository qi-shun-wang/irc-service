# IRCService

Intelligent Remote Control Service.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them

```
Give examples
```

### Installing

A step by step series of examples that tell you have to get a development env running

Build an executable file named IRCService for Android arm7 target :

```

CC=$ANDROID_NDK_ROOT/bin/arm-linux-androideabi/bin/arm-linux-androideabi-gcc GOOS="android" GOARCH="arm" CGO_ENABLED="1" go build -v -o IRCService /Users/shun/go/src/IRCService/Run/main.go


```
 
## Running the tests
 
Test muticast as a Listenner:

```
go run .../src/IRCService/tester.go
```
 
 
## Built With
* [Golang](https://golang.org) - go 1.9.3
* [go-CoAP](https://godoc.org/github.com/dustin/go-coap) - Package coap provides a CoAP client and server.
* [Multicast](https://github.com/dmichael/go-multicast) - Experiments in UDP Multicasting

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Shun,Billy** - *Initial work* - [IRCervice](https://bitbucket.org/ising99fullstack/ircservice)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc
