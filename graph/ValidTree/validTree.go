func validTree(n int, edges [][]int) bool {
   
   if len(edges)!=n-1{
    return false
   }
    
   // take adjusent 
   
   adj:=make(map[int][]int)

  for _,edge:=range edges{
    a,b:=edge[0],edge[1]

    adj[a]=append(adj[a],b)
    adj[b]=append(adj[b],a)
  }

  visited:=make(map[int]bool)

  if !dfs(0,-1,visited,adj){
    return false
  }
    return n==len(visited)
}

func dfs(node int,parent int ,visited map[int]bool,adj map[int][]int ) bool{
   
   if visited[node]{   //detect acyclic
    return false
   }

   visited[node]=true

   for _,neigh:=range adj[node]{
     
     if neigh==parent{
        continue
     }
   
    if !dfs(neigh,node,visited,adj){
        return false
    }


   }

   return true

}
