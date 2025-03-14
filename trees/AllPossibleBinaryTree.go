func allPossibleFBT(n int) []*TreeNode {
    

    if n%2==0{
        return []*TreeNode{}
    }

    if n==1{
        return []*TreeNode{&TreeNode{Val:0}}
    }
     
    var result[]*TreeNode
    
    
    for leftNode:=1;leftNode<n;leftNode+=2{
          rightNode:=n-1-leftNode

          Ltree:=allPossibleFBT(leftNode)
          Rtree:=allPossibleFBT(rightNode)

          for _,left:=range Ltree{
            for _,right:=range Rtree{

                node:=&TreeNode{Val: 0,Left:left,Right:right}

                result=append(result, node)
            }
          }
    }
 return result

}
