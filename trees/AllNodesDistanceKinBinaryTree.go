//cretae graph
//do bfs
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
  
  graph:=make(map[int][]int)
  creategraph(root,nil,graph)
  return bfs(graph,target.Val,k)
}


func creategraph(root*TreeNode,parent *TreeNode,graph map[int][]int){
     if root==nil{
        return
     }
     if parent!=nil{
        graph[root.Val]=append(graph[root.Val], parent.Val)
        graph[parent.Val]=append(graph[parent.Val], root.Val)
     }

     creategraph(root.Left,root,graph)
     creategraph(root.Right,root,graph)
     
     
}


func bfs(graph map[int][]int, targetVal int, k int)[]int{
   
   queue:=[]int{targetVal}
   isVisited:=map[int]bool{targetVal: true}
   
   for i:=0;i<k;i++{
    size:=len(queue)
    for j:=0;j<size;j++{
        node:=queue[0]
        queue=queue[1:]

        for _ ,n:=range graph[node]{
            if !isVisited[n]{
              isVisited[n]=true
              queue=append(queue, n)
            }
        }
    }
   }
 return queue
}
