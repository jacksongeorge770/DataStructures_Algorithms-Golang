
import "math"
func indorder(root *TreeNode,res *[]int){
  if root ==nil{
    return
  }
  indorder(root.Left,res)
  *res=append(*res, root.Val)
  indorder(root.Right,res)
}

func minDiffInBST(root *TreeNode) int {
    
    var result [] int

    indorder(root,&result)
    
    maxL:=math.MaxInt32
    for i:=1;i<len(result);i++{
        
        res:= result[i]-result[i-1]

        if res<maxL{
            maxL=res
        }

     
    }
    

return maxL
}
