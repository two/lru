package lru

import (
	"reflect"
	"testing"
)

func TestList_AddFirst(t *testing.T) {
	tests := map[string]struct {
		nodes      []*Node
		resultNext []string
		resultPrev []string
	}{
		"normal": {
			nodes: []*Node{
				&Node{val: "node1"},
				&Node{val: "node2"},
				&Node{val: "node3"},
				&Node{val: "node4"},
				&Node{val: "node5"},
			},
			resultNext: []string{"node5", "node4", "node3", "node2", "node1"},
			resultPrev: []string{"node1", "node2", "node3", "node4", "node5"},
		},

		"no element": {
			nodes:      []*Node{},
			resultNext: []string{},
			resultPrev: []string{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			list := List{}
			var got = []string{}
			for _, node := range tc.nodes {
				list.AddFirst(node)
			}

			// 从前往后查找
			head := list.head
			for head != nil {
				got = append(got, head.val)
				head = head.next
			}
			if !reflect.DeepEqual(got, tc.resultNext) {
				t.Fatalf("got: %v, want: %v", got, tc.resultNext)
			}

			// 从后往前查找
			tail := list.tail
			got = []string{}
			for tail != nil {
				got = append(got, tail.val)
				tail = tail.prev
			}
			if !reflect.DeepEqual(got, tc.resultPrev) {
				t.Fatalf("got: %v, want: %v", got, tc.resultPrev)
			}

		})
	}

}

func TestList_Remove(t *testing.T) {
	// empty list
	list := List{}
	list.Remove(list.head)
	if list.head != nil {
		t.Fatalf("got: %v, want: %v", list.head, nil)
	}

	// one element
	list = List{}
	nodes := []*Node{&Node{val: "node1"}}
	for _, node := range nodes {
		list.AddFirst(node)
	}
	list.Remove(list.head)
	if list.head != nil {
		t.Fatalf("got: %v, want: %v", list.head, nil)
	}

	// head element
	list = List{}
	nodes = []*Node{&Node{val: "node1"}, &Node{val: "node2"}}
	result := []string{"node1"}
	var got []string
	for _, node := range nodes {
		list.AddFirst(node)
	}
	list.Remove(list.head)
	head := list.head
	for head != nil {
		got = append(got, head.val)
		head = head.next
	}
	if !reflect.DeepEqual(got, result) {
		t.Fatalf("got: %v, want: %v", got, result)
	}

	// middle element
	list = List{}
	nodes = []*Node{&Node{val: "node1"}, &Node{val: "node2"}, &Node{val: "node3"}}
	result = []string{"node3", "node1"}
	for _, node := range nodes {
		list.AddFirst(node)
	}
	list.Remove(list.head.next)
	head = list.head
	got = []string{}
	for head != nil {
		got = append(got, head.val)
		head = head.next
	}
	if !reflect.DeepEqual(got, result) {
		t.Fatalf("got: %v, want: %v", got, result)
	}

	// tail element
	list = List{}
	nodes = []*Node{&Node{val: "node1"}, &Node{val: "node2"}, &Node{val: "node3"}}
	result = []string{"node3", "node2"}
	for _, node := range nodes {
		list.AddFirst(node)
	}
	list.Remove(list.head.next.next)
	head = list.head
	got = []string{}
	for head != nil {
		got = append(got, head.val)
		head = head.next
	}
	if !reflect.DeepEqual(got, result) {
		t.Fatalf("got: %v, want: %v", got, result)
	}
}

func TestList_Clear(t *testing.T) {
	list := List{}
	nodes := []*Node{&Node{val: "node1"}, &Node{val: "node2"}}
	for _, node := range nodes {
		list.AddFirst(node)
	}
}
