package main

import (
	"github.com/hopex-hk/go_sdk/example/accountclientexample"
	"github.com/hopex-hk/go_sdk/example/homeclientexample"
	"github.com/hopex-hk/go_sdk/example/marketclientexample"
	"github.com/hopex-hk/go_sdk/example/tradeclientexample"
)

func main() {
	run()
}

func run() {
	homeclientexample.RunAllExample()
	accountclientexample.RunAllExample()
	marketclientexample.RunAllExample()
	tradeclientexample.RunAllExample()
}
