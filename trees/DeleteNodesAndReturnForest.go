/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    delete_map:=make(map[int]bool)

    for _,val:=range to_delete{
      
      delete_map[val]=true
    }

   
  var result []*TreeNode
    dfs(root, true,  delete_map, &result)
    return result


}

func dfs(node*TreeNode,isRoot bool,d_map map[int]bool ,result *[]*TreeNode)*TreeNode{
   
   if node==nil{
    return nil
   }

   if isRoot && !d_map[node.Val]{
      *result=append(*result, node)
       
   }

   node.Left=dfs(node.Left,d_map[node.Val],d_map,result)
   node.Right=dfs(node.Right,d_map[node.Val],d_map,result)
   
   if d_map[node.Val]{
    return nil
   }

   return node

    
 
}
