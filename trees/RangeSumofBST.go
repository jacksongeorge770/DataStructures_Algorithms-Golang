//we can use BFS by using simple queue...  but here we are using dfs
func rangeSumBST(root *TreeNode, low int, high int) int {
    
   if root ==nil{
    return 0
   }
   sum:=0
   if root.Val>=low && root.Val<=high{
    sum+=root.Val
   }

   if root.Val>low{
    sum+=rangeSumBST(root.Left,low,high)
   }
    
    if root.Val<high{
    sum+=rangeSumBST(root.Right,low,high)
   }
  
  return sum
}
