package main

import "fmt"

type Node struct {
	Value int
}

// 用于构建结构体切片为最小堆，需要调用down函数
func Init(nodes []Node) {
	for i := len(nodes)/2 - 1; i >= 0; i-- {
		down(nodes, i, len(nodes))
	}
}

// 需要down（下沉）的元素在切片中的索引为i，n为heap的长度，将该元素下沉到该元素对应的子树合适的位置，从而满足该子树为最小堆的要求
func down(nodes []Node, i, n int) {
	parent := i
	for leftSon, rightSon := 2*parent+1, 2*parent+2; leftSon < n || rightSon < n; leftSon, rightSon = 2*parent+1, 2*parent+2 {
		next := parent
		if leftSon < n && nodes[parent].Value > nodes[leftSon].Value {
			next = leftSon
		}
		if rightSon < n && nodes[parent].Value > nodes[rightSon].Value {
			if nodes[rightSon].Value < nodes[next].Value {
				next = rightSon
			}
		}
		if next == parent {
			break
		}
		nodes[parent].Value, nodes[next].Value = nodes[next].Value, nodes[parent].Value
		parent = next
	}
}

// 用于保证插入新元素(j为元素的索引,切片末尾插入，堆底插入)的结构体切片之后仍然是一个最小堆
func up(nodes []Node, j int) {
	son := j
	for parent := (son - 1) / 2; parent >= 0; {
		if nodes[parent].Value <= nodes[son].Value {
			break
		} else {
			nodes[parent].Value, nodes[son].Value = nodes[son].Value, nodes[parent].Value
			son = parent
			parent = (son - 1) / 2
		}
	}
}

// 弹出最小元素，并保证弹出后的结构体切片仍然是一个最小堆，第一个返回值是弹出的节点的信息，第二个参数是Pop操作后得到的新的结构体切片
func Pop(nodes []Node) (Node, []Node) {
	min := nodes[0]
	nodes[0].Value = nodes[len(nodes)-1].Value
	down(nodes[0:len(nodes)-1], 0, len(nodes)-1)
	return min, nodes[0 : len(nodes)-1]
}

// 保证插入新元素时，结构体切片仍然是一个最小堆，需要调用up函数
func Push(node Node, nodes []Node) []Node {
	nodes = append(nodes, node)
	up(nodes, len(nodes)-1)
	return nodes
}

// 移除切片中指定索引的元素，保证移除后结构体切片仍然是一个最小堆
func Remove(nodes []Node, node Node) []Node {
	for i, value := range nodes {
		if value.Value == node.Value {
			nodes[i].Value, nodes[len(nodes)-1].Value = nodes[len(nodes)-1].Value, nodes[i].Value
			down(nodes[:len(nodes)-1], i, len(nodes)-1)
			break
		}
	}
	return nodes[0 : len(nodes)-1]
}

func main() {
	test := []Node{{18}, {32}, {5}, {4}, {9}, {26}}
	Init(test)
	test = Push(Node{0}, test)
	test = Push(Node{88}, test)
	test = Remove(test, Node{5})
	temp := test
	a := Node{0}
	for i := 0; i < len(test); i++ {
		a, temp = Pop(temp)
		fmt.Println(a)
	}
}