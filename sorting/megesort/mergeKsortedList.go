/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
if len(lists) == 0 {
		return nil
	}
  for len(lists)>1{
    
    var mergeSorted []*ListNode

    for i:=0;i<len(lists);i+=2{
          
      if i+1<len(lists){
        mergeSorted=append(mergeSorted,merge(lists[i],lists[i+1]))
      }else{
        mergeSorted=append(mergeSorted,lists[i])
      }


     

    }
    lists=mergeSorted

  }

  return lists[0]

}


func merge(l1,l2 *ListNode)*ListNode{
    
   dummy:=&ListNode{}
   cur:=dummy
   for l1!=nil && l2!=nil{
        
        if l1.Val<l2.Val{
            cur.Next=l1
            l1=l1.Next
        }else{
           cur.Next=l2
           l2= l2.Next
        }
        cur=cur.Next

   }

   if l1!=nil{
    cur.Next=l1
    
   }else{
    cur.Next = l2
   }

 return dummy.Next

}
