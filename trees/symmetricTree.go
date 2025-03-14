/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    
 if root==nil{
    return true
 }

 return mirrorCheck(root.Left,root.Right) 

}


func mirrorCheck(root1*TreeNode,root2*TreeNode)bool{
    if root1==nil&& root2==nil{
        return true
    }

    if root1==nil || root2==nil{
        return false
    }

    if root1.Val!=root2.Val{
        return false
    }

    return mirrorCheck(root1.Left,root2.Right)&& mirrorCheck(root1.Right,root2.Left)
}
