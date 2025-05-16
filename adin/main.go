package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Weight int
	Next   *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) AddNode(weight int) {
	newNode := &Node{Weight: weight}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (ll *LinkedList) ToSlice() []int {
	var weights []int
	current := ll.Head
	for current != nil {
		weights = append(weights, current.Weight)
		current = current.Next
	}
	return weights
}

func (ll *LinkedList) FromSlice(weights []int) {
	ll.Head = nil
	for _, weight := range weights {
		ll.AddNode(weight)
	}
}

func (ll *LinkedList) Sort(ascending bool) {
	weights := ll.ToSlice()
	if ascending {
		sort.Ints(weights)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(weights)))
	}
	ll.FromSlice(weights)
}

func (ll *LinkedList) RemoveFirst() *Node {
	if ll.Head == nil {
		return nil
	}
	removed := ll.Head
	ll.Head = ll.Head.Next
	return removed
}

func (ll *LinkedList) RemoveLast() *Node {
	if ll.Head == nil {
		return nil
	}
	if ll.Head.Next == nil {
		removed := ll.Head
		ll.Head = nil
		return removed
	}

	current := ll.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	removed := current.Next
	current.Next = nil
	return removed
}

func (ll *LinkedList) GetHeaviest() *Node {
	if ll.Head == nil {
		return nil
	}
	ll.Sort(false)
	return ll.Head
}

func (ll *LinkedList) GetLightest() *Node {
	if ll.Head == nil {
		return nil
	}
	ll.Sort(true)
	return ll.Head
}

func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Weight, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	ll := &LinkedList{}
	ll.AddNode(5)
	ll.AddNode(3)
	ll.AddNode(8)
	ll.AddNode(1)

	fmt.Println("Original list:")
	ll.Display()

	ll.Sort(true)
	fmt.Println("Sorted list (ascending):")
	ll.Display()

	ll.Sort(false)
	fmt.Println("Sorted list (descending):")
	ll.Display()

	removed := ll.RemoveFirst()
	fmt.Printf("Removed first potato: %d\n", removed.Weight)
	ll.Display()

	removed = ll.RemoveLast()
	fmt.Printf("Removed last potato: %d\n", removed.Weight)
	ll.Display()

	heaviest := ll.GetHeaviest()
	fmt.Printf("Heaviest potato: %d\n", heaviest.Weight)

	lightest := ll.GetLightest()
	fmt.Printf("Lightest potato: %d\n", lightest.Weight)
}
