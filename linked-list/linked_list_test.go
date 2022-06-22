package linkedlist

import (
	"testing"
)

var (
	firstNode  = NewNode("Once")
	secondNode = NewNode("Upon")
	thirdNode  = NewNode("A")
	fouthNode  = NewNode("Time")
)

func TestRead(t *testing.T) {
	firstNode.nextNode = secondNode
	secondNode.nextNode = thirdNode
	thirdNode.nextNode = fouthNode

	tests := []struct {
		index    int
		expected string
	}{
		{index: 0, expected: "Once"},
		{index: 1, expected: "Upon"},
		{index: 2, expected: "A"},
		{index: 3, expected: "Time"},
		{index: 4, expected: ""},
	}

	linkedList := NewLinkedList(firstNode)

	for _, test := range tests {
		result := linkedList.Read(test.index)
		if result != test.expected {
			t.Fatalf("Expectation doesn't match with actual result expect %s got %s\n", test.expected, result)
		}
	}
}

func TestDelete(t *testing.T) {
	firstNode.nextNode = secondNode
	secondNode.nextNode = thirdNode
	thirdNode.nextNode = fouthNode

	linkedList := NewLinkedList(firstNode)
	linkedList.Delete(2)
	var result string

	for i := 0; i < 3; i++ {
		result += linkedList.Read(i)
	}

	if result != "OnceUponTime" {
		t.Fatalf("Expectation doesn't match with actual result expect %s got %s", "OnceUponTime", result)
	}
}

func TestSearch(t *testing.T) {
	firstNode.nextNode = secondNode
	secondNode.nextNode = thirdNode
	thirdNode.nextNode = fouthNode

	linkedList := NewLinkedList(firstNode)

	tests := []struct {
		input    string
		expected int
	}{
		{input: "Once", expected: 0},
		{input: "Upon", expected: 1},
		{input: "A", expected: 2},
		{input: "Time", expected: 3},
		{input: "Hello", expected: -1},
	}

	for _, test := range tests {
		result := linkedList.Search(test.input)
		if result != test.expected {
			t.Fatalf("Expectation doesn't match with actual result expect %d got %d", test.expected, result)
		}
	}

}
