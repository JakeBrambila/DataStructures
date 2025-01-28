package main

func main() {
	l := LinkedList{}

	Node1 := &Node{Value: 50}
	l.AddToEnd(Node1)

	l.Print()

	Node2 := &Node{Value: 30}
	Node3 := &Node{Value: 20}
	Node4 := &Node{Value: 60}
	Node5 := &Node{Value: 90}

	l.AddToBeginning(Node2)
	l.AddToBeginning(Node3)

	l.Print()

	l.AddToEnd(Node4)
	l.AddToEnd(Node5)

	l.Print()
}
