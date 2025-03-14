/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
      

    leave1:=[]int{}
    leave2:=[]int{}

    dfs(root1,&leave1)
    dfs(root2,&leave2)

    return isequal(leave1,leave2)
      
         
}

func dfs(root *TreeNode,leaves *[]int){
     if root == nil {
		return
	}
      if root.Left==nil && root.Right==nil{
        *leaves=append(*leaves,root.Val)
      }

      dfs(root.Left,leaves)
      dfs(root.Right,leaves)

      
}

func isequal(a []int, b []int) bool{
     if len(a)!=len(b){
        return false
     }
    for i:=0;i<len(a);i++{
         
         if a[i]!=b[i]{
            return false
         }

    }

    return true
}
