func findTheWinner(n int, k int) int {
     if k == 1 {
        return n // Special case: winner is always the n-th person
    }
    head:=CL(n)
    cur:=head
    for cur.Next!=cur{
    for i:=1;i<k-1;i++{
    
        cur=cur.Next
    
      }
      cur.Next= cur.Next.Next
      cur=cur.Next
    
    }

    return cur.Val


    
}

func CL(n int)*ListNode{
   if(n==0){
    return nil
   }
   head:=&ListNode{Val:1}
   cur:=head
   for  i:=2;i<=n;i++{
     cur.Next= &ListNode{Val:i}
     cur=cur.Next
   }

   cur.Next=head

return head
}
