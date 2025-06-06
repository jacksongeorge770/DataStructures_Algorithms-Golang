func cloneGraph(node *Node) *Node {
   
   if node==nil{
    return nil
   }
   visited:=make(map[*Node]*Node)
   return dfs(node,visited)

}

func dfs(node *Node, visited map[*Node]*Node) *Node{
    
    if clone, found:=visited[node];found{
        return clone
    }

    clone:=&Node{Val: node.Val}
    visited[node]=clone

    for _ ,neighbour:=range node.Neighbors{
        clone.Neighbors=append(clone.Neighbors, dfs(neighbour,visited))
    }

return clone

}
