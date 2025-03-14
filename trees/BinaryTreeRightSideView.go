/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root==nil{
        return []int{}
    }
   var result []int
   queue:=[]*TreeNode{root}

   for len(queue)>0{
    length:=len(queue)
    for i:=0;i<length;i++{
        
        node:=queue[0]
        queue=queue[1:]
        
        if i==length-1{
            result=append(result,node.Val)
        }
        if node.Left!=nil{
            queue=append(queue,node.Left)
        }
        if node.Right!=nil{
            queue=append(queue,node.Right)
        }

    }
    
   }
    

return result

}
