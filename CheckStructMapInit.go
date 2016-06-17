package main

import (
	"fmt"
	"time"
)

type internalMap struct {
	version string
	state   int
}

type testMap struct {
	updateTime time.Time
	arguments  map[string]interface{}
	comps      map[string]internalMap
	status     int
}

func main() {
	var ret testMap
	ret.updateTime = time.Now()
	ret.status = 1111

	args := make(map[string]interface{})
	args["args"] = "testargs"

	ret.arguments = args

	im1 := internalMap{
		version: "v1.2",
		state:   10,
	}

	im2 := internalMap{
		version: "v1.3",
		state:   11,
	}

	ret.comps = make(map[string]internalMap) //we must use make to init a map

	ret.comps["im1"] = im1
	ret.comps["im2"] = im2

	fmt.Println(ret)
}
