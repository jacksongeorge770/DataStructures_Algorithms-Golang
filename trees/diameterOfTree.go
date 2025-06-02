
func diameterOfBinaryTree(root *TreeNode) int {
diameter:=0

dfs(root,&diameter)
return diameter

}

func dfs(root*TreeNode,diameter *int)int{

    if root==nil{
        return 0
    }

    leftH:=dfs(root.Left,diameter)
    rightH:=dfs(root.Right,diameter)

    *diameter=max(*diameter,leftH+rightH)

    return 1+ max(leftH,rightH)
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
