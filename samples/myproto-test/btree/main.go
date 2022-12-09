// 用chatGPT生成的BTree代码

package main

import (
	"fmt"
	"math/rand"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node *TreeNode) Insert(val int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: val}
	}
	if val < node.Value {
		node.Left = node.Left.Insert(val)
	} else {
		node.Right = node.Right.Insert(val)
	}
	return node
}

func (node *TreeNode) String() string {
	if node == nil {
		return ""
	}
	return fmt.Sprintf("%v %v %v", node.Left, node.Value, node.Right)
}

func main() {
	var root *TreeNode
	for i := 0; i < 10; i++ {
		val := rand.Int()
		fmt.Printf("Inserting %v into tree\n", val)
		root = root.Insert(val)
		fmt.Println(root)
	}
}
