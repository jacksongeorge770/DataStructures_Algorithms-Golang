import "math"
func isValidBST(root *TreeNode) bool {
     max:= math.MaxInt64
     min:= math.MinInt64
   return validate(root,max,min)    
  
}

func validate(root*TreeNode,max int,min int)bool{
     
     if root==nil{
      return true
     }

     if root.Val<=min||root.Val>=max{
            return false
     }


  return validate(root.Left,root.Val,min)&& validate(root.Right,max,root.Val)    

}
