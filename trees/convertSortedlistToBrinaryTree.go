
func sortedListToBST(head *ListNode) *TreeNode {
     
     if head==nil{
        return nil
     }

     mid:=findMid(head)
     
     root:=&TreeNode{Val: mid.Val}
     if head!=mid{
        root.Left=sortedListToBST(head)
     }
    
    root.Right=sortedListToBST(mid.Next)

 return root
}

func findMid(head*ListNode)*ListNode{
    
    f,s:=head,head
    var prev *ListNode =nil
   
    for  f!=nil&&f.Next!=nil {
        f=f.Next.Next
        prev=s
        s=s.Next
    }
      if prev != nil {
        prev.Next = nil
    }
    return s
}
