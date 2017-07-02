Golang for API development
==========

This image builds a development environment to work with the programming language Go and the dependencies needed to create RESTFUL APIs.

Build
-----
```
    docker build -t mbonell/go-dev-apis .
```

Or just pull it from Docker Hub
------------------------------
```
    docker pull mbonell/go-dev-apis
```

Run the Golang APIs container
----------------------------
```
    docker run -it mbonell/go-dev-apis
```

Or run the container using a host directory as a data volume
----------------------------
```
    docker run -it -v <host_directory>:/root/workspace mbonell/go-dev-apis
```

Extra configurations for the vim-go IDE
----------------------------
The container uses [vim-go](https://github.com/farazdagi/vim-go-ide) as IDE for Golang. Use the alias vimgo to open it.
```
vimgo
```

To make sure that all extra Go tools (godep, gocode etc) are present on the container,  run the following command from within Vim:
```
:GoInstallBinaries
```
