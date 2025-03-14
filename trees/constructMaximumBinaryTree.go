
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {

    if len(nums)==0{
        return nil
    }


    maxid:=findMax(nums)
    maxele:=nums[maxid]

    root:=&TreeNode{Val:maxele}

    root.Left=constructMaximumBinaryTree(nums[:maxid])
    root.Right=constructMaximumBinaryTree(nums[maxid+1:])


   return root

}



func findMax(nums []int)int{

    maxid:=0
    for i:=0;i<len(nums);i++{
        if nums[i]>nums[maxid]{
            maxid=i
        }


    }
    return maxid
}
