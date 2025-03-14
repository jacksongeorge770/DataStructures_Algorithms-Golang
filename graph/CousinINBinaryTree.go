
func isCousins(root *TreeNode, x int, y int) bool {
    
     if root == nil {
        return false
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
       length := len(queue)
       var parentX, parentY *TreeNode
       var foundX, foundY bool
  
         for i := 0; i < length; i++ {
            node := queue[0]
            queue = queue[1:]
            
            if node.Left!=nil{
                 if node.Left.Val==x{
                    parentX=node
                    foundX=true
                 }
                 if node.Left.Val==y{
                    parentY=node
                    foundY=true
                 }
                 queue=append(queue, node.Left)
            }
            if node.Right!=nil{
                 if node.Right.Val==x{
                    parentX=node
                    foundX=true
                 }
                 
                 if node.Right.Val==y{
                    parentY=node
                    foundY=true
                 }
                 queue=append(queue, node.Right)
            }

         }

         if foundX&&foundY{
            return parentX!=parentY
         }
         

    }

return false
}
