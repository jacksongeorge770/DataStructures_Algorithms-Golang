https://leetcode.com/problems/palindrome-linked-list/?envType=problem-list-v2&envId=9gvhzx23
func isPalindrome(head *ListNode) bool {
  
  
  mid:=findmid(head)
  
  secondHalf:=reverse(mid)

  firstHalf:=head
  
  temp:=secondHalf
  
  for firstHalf!=nil&&secondHalf!=nil{
      
      if firstHalf.Val!=secondHalf.Val{
        return false
      }

      firstHalf=firstHalf.Next
      secondHalf=secondHalf.Next
  }
   reverse(temp) 
  return true

}


func findmid(head *ListNode) *ListNode{
     
     fast:=head
     slow:=head

     for fast!=nil && fast.Next!=nil{

        fast=fast.Next.Next
        slow=slow.Next
     }

     return slow
   
}

func reverse(head *ListNode)*ListNode{

   dummy:=&ListNode{}

   cur:=head
   prev:=dummy
   
   for cur!=nil{
    next:=cur.Next
    cur.Next=prev.Next
    prev.Next=cur
    cur=next
   }

   return dummy.Next


}

//stack based apprch


func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true // Empty or single-node list is always a palindrome
    }

    // Step 1: Use a stack to store values
    stack := []int{}
    current := head

    // Push all node values onto the stack
    for current != nil {
        stack = append(stack, current.Val)
        current = current.Next
    }

    // Step 2: Traverse the list again and compare with the stack
    current = head
    for len(stack) > 0 {
        val := stack[len(stack)-1] // Get the top element from the stack
        stack = stack[:len(stack)-1] // Pop the element
        if current.Val != val {
            return false // If values don't match, return false
        }
        current = current.Next
    }

    return true
}

