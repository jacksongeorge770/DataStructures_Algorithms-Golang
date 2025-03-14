import (
    "math"
)

func largestValues(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    queue := []*TreeNode{root}
    result := []int{}

    for len(queue) > 0 {
        length := len(queue)
        maxVal := math.MinInt32

        for i := 0; i < length; i++ {
            node := queue[0]
            queue = queue[1:]

            if maxVal < node.Val {
                maxVal = node.Val
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, maxVal)
    }

    return result
}
