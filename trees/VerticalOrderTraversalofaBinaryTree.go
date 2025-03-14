/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import(
    "container/list"
    "sort"
)

func verticalTraversal(root *TreeNode) [][]int {

   queue:=list.New()
   queue.PushBack([]interface{}{root,0})

   hashMap:=make(map[int][]int)

   min,max:=0,0
   
   for queue.Len()>0{

    length:=queue.Len()
    tempMap:=make(map[int][]int)
    
    for i:=0;i<length;i++{

        element:=queue.Front()
        queue.Remove(element)
         
        temp:=element.Value.([]interface{})
        node:=temp[0].(*TreeNode)
        level:=temp[1].(int)

        tempMap[level]=append(tempMap[level],node.Val)

        if max<level{
            max=level
        }

        if min>level{
            min=level
        }


        if node.Left!=nil{
            queue.PushBack([]interface{}{node.Left,level-1})
        }
        if node.Right!=nil{
            queue.PushBack([]interface{}{node.Right,level+1})
        }

    }
     
    for level,nodes:=range tempMap{
     sort.Ints(nodes)
     hashMap[level]=append(hashMap[level],nodes...)

    }

     

   }

    var result [][]int

    for i:=min;i<=max;i++{
        
        result=append(result,hashMap[i])
    }

return result
}
