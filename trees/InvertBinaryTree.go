import(
    "container/list"
)
func invertTree(root *TreeNode) *TreeNode {
    
    if root==nil{
        return nil
    }
    queue:=list.New()
    queue.PushBack(root)
     
     for queue.Len()>0{
        length:=queue.Len()

        for i:=0;i<length;i++{
             
             front:=queue.Front()
             temp:= queue.Remove(front)
             node:=temp.(*TreeNode)
             
             node.Left,node.Right=node.Right,node.Left

             if node.Left!=nil{
                queue.PushBack(node.Left)
             }


             if node.Right!=nil{
                queue.PushBack(node.Right)
             }
        
        
        }
     }
   return root
}
