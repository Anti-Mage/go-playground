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

	fmap := template.FuncMap{
		"minus": func(a, b int) int {
			return a - b
		},
	}

	t = t.Funcs(fmap)

	t, err := t.ParseFiles(filename)
	if err != nil {
		errMsg(18, err)
	}

	tr := t.Tree
	rn := tr.Root.Nodes

	l := len(rn)
	fmt.Println("length of nodes is:", l)

	for _, v := range t.Root.Nodes {
		if v.Type() != parse.NodeText {
			fmt.Printf("NodeName:%s NodeType:%v\n", v.String(), v.Type())

			if v.Type() == parse.NodeAction {
				tmpNode := v.(*parse.ActionNode)
				for kk, vv := range tmpNode.Pipe.Cmds {
					fmt.Printf("	index:%d, value:%v\n", kk, vv)
					for kkk, vvv := range vv.Args {
						fmt.Printf("		index:%d, value:%v\n", kkk, vvv)
					}
				}
			} else if v.Type() == parse.NodeRange {
				tmpNode := v.(*parse.RangeNode)
				for kk, vv := range tmpNode.Pipe.Cmds {
					fmt.Printf("	index:%d, value:%v\n", kk, vv)
					for kkk, vvv := range vv.Args {
						fmt.Printf("		index:%d, value:%v\n", kkk, vvv)
					}
				}
			}
		}
	}

	//fmt.Println("t.ParseName:", t.ParseName)
	//fmt.Printf("t.Tree.Name:%s, t.Tree.ParseName:%s", t.Tree.Name, t.Tree.ParseName)

}
