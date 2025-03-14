func removeZeroSumSublists(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head} // Dummy node to handle edge cases
    prefixsum := make(map[int]*ListNode) // Stores prefix sum and corresponding node

    sum := 0
    cur:=dummy

    for cur!=nil{
        sum+=cur.Val
        prefixsum[sum]=cur
        cur=cur.Next
    }

    sum=0
    cur=dummy

    for cur!=nil{
        sum+=cur.Val
        cur.Next=prefixsum[sum].Next
        cur=cur.Next

    }

    return dummy.Next


}
