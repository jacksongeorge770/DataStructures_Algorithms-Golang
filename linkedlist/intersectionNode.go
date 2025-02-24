/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    
  lenA:=getLength(headA)
  lenB:=getLength(headB)

  
  if lenA>lenB{
     difa:=lenA-lenB
     for i:=0;i<difa;i++{
        headA=headA.Next
     }

  }else{
    difb:=lenB-lenA
     for i:=0;i<difb;i++{
        headB=headB.Next
     }

  }

  for headA!=nil && headB!=nil{
     
     if headA==headB{
        return headA
     }
    
    headA=headA.Next
    headB=headB.Next
  }

 return nil
}

func getLength(head *ListNode) int{
    
    length:=0
    for head!=nil{
        length++
        head=head.Next
        
    }
   
   return length

}
