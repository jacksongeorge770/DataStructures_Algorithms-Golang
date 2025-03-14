import(
    "container/list"
)
func maxDepth(root *TreeNode) int {
   if root==nil{
    return 0
   }
   queue:=list.New()

   queue.PushBack(root)
   depth:=0
   for  queue.Len()>0{

    length:=queue.Len()

    for i:=0;i<length;i++{
          
          temp:=queue.Front()
          node:=temp.Value.(*TreeNode)
          queue.Remove(temp)
          if node.Left!=nil{
                queue.PushBack(node.Left)
          }
          if node.Right!=nil{
            queue.PushBack(node.Right)
       
          }
       }
    
    depth++
   
   }
 return depth
}












