/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    
    dummy:=&ListNode{Next:nil}
 
    
    cur:=head
    
    for cur!=nil{
         prev:=dummy
         next:=cur.Next
         for prev.Next!=nil && prev.Next.Val < cur.Val { 
              prev=prev.Next
         }
         cur.Next=prev.Next
         prev.Next=cur

         cur=next

    }

return dummy.Next

}
