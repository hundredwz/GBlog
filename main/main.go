package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hundredwz/GBlog/config"
	"github.com/hundredwz/GBlog/server"
	"os"
)

var (
	help  bool
	port  int
	debug bool
)

func init() {
	flag.BoolVar(&help, "h", false, "print help")
	flag.IntVar(&port, "p", 701, "the port to run")
	flag.BoolVar(&debug, "d", false, "release mode")
	flag.Usage = usage
}
func usage() {
	fmt.Fprintf(os.Stderr, `GBlog help 
Usage: gblog [-p port] [-d]

Options: 
-h          : print help message
-d          : run as debug mode
-p port     : use custom port; default is 701
`)
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	blog := gin.Default()

	server.InitServer(blog)

	p := fmt.Sprintf(":%v", port)
	config.Addr = p
	blog.Run(p)
}
