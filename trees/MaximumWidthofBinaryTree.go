import(
    "container/list"
)
func widthOfBinaryTree(root *TreeNode) int {
    
    queue:=list.New()

    queue.PushBack([]interface{}{root,0})
    
    
     maxLength:=0
    for queue.Len()>0{
      firstidx:=queue.Front().Value.([]interface{})[1].(int)
      lastidx:=0
      length:=queue.Len()
      
     
      for i:=0;i<length;i++{
             
             element:=queue.Front()
             queue.Remove(element)
             temp:=element.Value.([]interface{})
             node:=temp[0].(*TreeNode)
             index:=temp[1].(int)-firstidx

             lastidx=index

             
             
             if node.Left!=nil{
                queue.PushBack([]interface{}{node.Left,2*index})
             }
              

             if node.Right!=nil{
                queue.PushBack([]interface{}{node.Right,2*index+1})
             }

             

      }
     
    maxLength=max(maxLength,lastidx+1)


    }
   return maxLength


}

func max(a,b int)int{

    if a>b{
        return a
    }else{
        return b
    }
}
