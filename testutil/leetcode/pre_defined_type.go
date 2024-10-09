package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (o *TreeNode) toRawString() string {
	nodes := []*TreeNode{}
	queue := []*TreeNode{o}
	for len(queue) > 0 {
		o, queue = queue[0], queue[1:]
		nodes = append(nodes, o)
		if o != nil {
			queue = append(queue, o.Left, o.Right)
		}
	}

	for len(nodes) > 0 && nodes[len(nodes)-1] == nil {
		nodes = nodes[:len(nodes)-1]
	}

	sb := &strings.Builder{}
	sb.WriteByte('[')
	for _, node := range nodes {
		if sb.Len() > 1 {
			sb.WriteByte(',')
		}
		if node != nil {
			sb.WriteString(strconv.Itoa(node.Val))
		} else {
			sb.WriteString("null")
		}
	}
	sb.WriteByte(']')
	return sb.String()
}

func buildTreeNode(rawArray string) (root *TreeNode, err error) {
	if len(rawArray) < 2 || rawArray[0] != '[' || rawArray[len(rawArray)-1] != ']' {
		return nil, fmt.Errorf("invalid test data %s", rawArray)
	}

	rawArray = rawArray[1 : len(rawArray)-1]
	if rawArray == "" {
		return
	}
	splits := strings.Split(rawArray, ",")

	nodes := make([]*TreeNode, len(splits))
	for i, s := range splits {
		s = strings.TrimSpace(s)
		if s != "null" {
			nodes[i] = &TreeNode{}
			var er error
			nodes[i].Val, er = strconv.Atoi(s)
			if er != nil {
				return nil, er
			}
			//nodes[i].Val = s[1 : len(s)-1]
		}
	}
	kids := nodes

	root, kids = kids[0], kids[1:]
	for _, node := range nodes {
		if node != nil {
			if len(kids) > 0 {
				node.Left, kids = kids[0], kids[1:]
			}
			if len(kids) > 0 {
				node.Right, kids = kids[0], kids[1:]
			}
		}
	}
	return
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (o *ListNode) toRawString() string {
	sb := &strings.Builder{}
	sb.WriteByte('[')
	for ; o != nil; o = o.Next {
		if sb.Len() > 1 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.Itoa(o.Val))
	}
	sb.WriteByte(']')
	return sb.String()
}

func buildListNode(rawArray string) (head *ListNode, err error) {
	if len(rawArray) < 2 || rawArray[0] != '[' || rawArray[len(rawArray)-1] != ']' {
		return nil, fmt.Errorf("invalid test data %s", rawArray)
	}

	rawArray = rawArray[1 : len(rawArray)-1]
	if rawArray == "" {
		return
	}
	splits := strings.Split(rawArray, ",")

	nodes := make([]*ListNode, len(splits))
	for i, s := range splits {
		s = strings.TrimSpace(s)
		nodes[i] = &ListNode{}
		var er error
		nodes[i].Val, er = strconv.Atoi(s)
		if er != nil {
			return nil, er
		}
		//nodes[i].Val = s[1 : len(s)-1]
	}

	head = nodes[0]
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	return
}

type Point struct {
	X int
	Y int
}

func (p *Point) toRawString() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func buildPoint(rawPoint string) (p *Point, err error) {
	if len(rawPoint) < 5 || rawPoint[0] != '(' || rawPoint[len(rawPoint)-1] != ')' {
		return nil, fmt.Errorf("invalid test data %s", rawPoint)
	}
	rawPoint = rawPoint[1 : len(rawPoint)-1]
	splits := strings.Split(rawPoint, ",")
	x, err := strconv.Atoi(strings.TrimSpace(splits[0]))
	if err != nil {
		return
	}
	y, err := strconv.Atoi(strings.TrimSpace(splits[1]))
	if err != nil {
		return
	}
	return &Point{x, y}, nil
}

//

type Interval struct {
	Start int
	End   int
}

func (p *Interval) toRawString() string {
	return fmt.Sprintf("[%d,%d]", p.Start, p.End)
}

func buildInterval(rawInterval string) (p *Interval, err error) {
	if len(rawInterval) < 5 || rawInterval[0] != '[' || rawInterval[len(rawInterval)-1] != ']' {
		return nil, fmt.Errorf("invalid test data %s", rawInterval)
	}
	rawInterval = rawInterval[1 : len(rawInterval)-1]
	splits := strings.Split(rawInterval, ",")
	x, err := strconv.Atoi(strings.TrimSpace(splits[0]))
	if err != nil {
		return
	}
	y, err := strconv.Atoi(strings.TrimSpace(splits[1]))
	if err != nil {
		return
	}
	return &Interval{x, y}, nil
}
