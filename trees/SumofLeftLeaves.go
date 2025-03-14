/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    // we can use bfs or level order travese
    if root==nil{
        return 0
    }

     queue:=[]*TreeNode{root}
      sum:=0
     for len(queue)>0{
          
          length:=len(queue)

          for i:=0;i<length;i++{

            node:=queue[0]
            queue=queue[1:]

            if node.Left!=nil && node.Left.Left==nil && node.Left.Right==nil{
                sum+=node.Left.Val
            }

            if node.Left!=nil{
                queue=append(queue, node.Left)
            }

            if node.Right!=nil{
                queue=append(queue, node.Right)
            }
          }
       
     }

  return sum
}
