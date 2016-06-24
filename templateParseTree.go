package main

import (
	"fmt"
	"text/template"
	"text/template/parse"
)

func errMsg(line int, err error) {
	fmt.Println("line :", line, "error occured:", err)
}

func main() {
	filename := "template.conf"
	t := template.New(filename)
	t, err := t.ParseFiles(filename)
	if err != nil {
		errMsg(18, err)
	}

	tr := t.Tree
	rn := tr.Root.Nodes

	l := len(rn)
	fmt.Println("length of nodes is:", l)

	for _, v := range t.Root.Nodes {
		if v.Type() == parse.NodeAction {
			//fmt.Println("assert type NodeAction:", v.Type() == parse.NodeAction)
			tmpNode := v.(*parse.ActionNode)
			for kk, vv := range tmpNode.Pipe.Cmds {
				fmt.Printf("   index:%d, value:%v\n", kk, vv)
			}

			/*for kk, vv := range tmpNode.Pipe.Decl {
				fmt.Printf("   index:%d, value:%v\n", kk, vv)
			}*/
		}
	}

	fmt.Println("t.ParseName:", t.ParseName)
	fmt.Printf("t.Tree.Name:%s, t.Tree.ParseName:%s", t.Tree.Name, t.Tree.ParseName)

}
