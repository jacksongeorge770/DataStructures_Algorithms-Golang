/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
     
   if root==nil{
    return 0
   }
   mDepth:=0
   for _,node:=range root.Children{
       
       childDepth:=maxDepth(node)
       if childDepth>mDepth{
        mDepth=childDepth
       }

   }

  return mDepth+1
}
