package main

import (
	"fmt"
)

type NodeAttr struct {
	Attr string
}

type Node struct {
	NodeId string
	Attr   *NodeAttr
}

func main() {
	tmp := &Node{
		NodeId: "111",
	}

	fmt.Println("line 21:", tmp)
	fmt.Println("line 22", tmp.NodeId, tmp.Attr.Attr)
}
