package main

import "fmt"

// Node structure for the tree
type Node struct {
    Value int
    Left  *Node
    Right *Node
}

// LevelOrder function for breadth-first traversal
func LevelOrder(root *Node) {
    if root == nil {
        return // If the tree is empty, return
    }

    // Create a queue to hold nodes for level order traversal
    queue := []*Node{}
    queue = append(queue, root) // Enqueue the root node

    // While the queue is not empty, continue processing nodes
    for len(queue) > 0 {
        // Dequeue the first node from the queue
        current := queue[0]
        queue = queue[1:]

        // Process the current node (print its value)
        fmt.Print(current.Value, " -> ")

        // Enqueue left child if exists
        if current.Left != nil {
            queue = append(queue, current.Left)
        }

        // Enqueue right child if exists
        if current.Right != nil {
            queue = append(queue, current.Right)
        }
    }
}

func main() {
    // Creating a sample tree
    root := &Node{Value: 5}
    root.Left = &Node{Value: 3}
    root.Right = &Node{Value: 8}
    root.Left.Left = &Node{Value: 2}
    root.Left.Right = &Node{Value: 4}
    root.Right.Right = &Node{Value: 7}

    // Call the LevelOrder function to print the nodes in level order
    LevelOrder(root)
}

