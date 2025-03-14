/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
       if root == nil {
		return [][]int{} // Return empty if no nodes
	}

       var result [][]int
       queue:=[]*TreeNode{root}
       for len(queue)>0{
         length:=len(queue)
         var temp []int
         for i:=0;i<length;i++{
              
              
              node:=queue[0]
              queue=queue[1:]
              
              temp=append(temp,node.Val)
              if node.Left!=nil{
                   
                   queue=append(queue,node.Left)
                   
              }

              if node.Right!=nil{
                 
                 queue=append(queue,node.Right)
              }

         
         }
        
        
       result=append(result,temp)
 

       }
       return result
}
