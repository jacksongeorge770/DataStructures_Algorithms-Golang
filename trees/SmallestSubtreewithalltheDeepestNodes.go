
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
      
      _,node:=dfs(root)

      return node

}
func dfs( root *TreeNode)(int,*TreeNode){
    
    if root==nil{
        return 0,nil
    }

    LeftDepth,LeftNode:=dfs(root.Left)
    RightDepth,RightNode:=dfs(root.Right)
    
    if LeftDepth==RightDepth{
         
         return  max(LeftDepth,RightDepth)+1,root
    }

    if LeftDepth>RightDepth{

        return max(LeftDepth,RightDepth)+1,LeftNode
    }
    return max(LeftDepth,RightDepth)+1,RightNode
}


func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
