package main

import (
	"github.com/bingbaba/tool/color"
	"github.com/bingbaba/util/logs"
	"github.com/xyalan/migateway"
	"time"
)

var (
	LOGGER = logs.GetBlogger()
)

func init() {
	logs.SetDebug(true)
}

func main() {
	manager, err := migateway.NewAqaraManager(nil)
	if err != nil {
		panic(err)
	}
	manager.SetAESKey("t7ew6r4y612eml0f")

	gateway := manager.GateWay
	for _, color := range color.COLOR_ALL {
		err = gateway.ChangeColor(color)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}

	err = gateway.Flashing(color.COLOR_RED)
	if err != nil {
		panic(err)
	}

	//do something...
	time.Sleep(10 * time.Second)
}
