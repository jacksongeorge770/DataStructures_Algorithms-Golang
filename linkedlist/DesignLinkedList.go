type MyLinkedList struct {
    head *ListNode
}

func Constructor() MyLinkedList {
    return MyLinkedList{}
}



func (this *MyLinkedList) Get(index int) int {
  
  cur:=this.head
  i:=0
  for i<index && cur!=nil{
      cur=cur.Next
      i++

  }

  if cur==nil{
    return -1 
  }
 
 return cur.Val

}


func (this *MyLinkedList) AddAtHead(val int)  {
    
  newnode:=&ListNode{Val:val}
  newnode.Next=this.head
  this.head=newnode

}


func (this *MyLinkedList) AddAtTail(val int)  {
   
   newnode:=&ListNode{Val:val}
   if this.head==nil{
          
     this.head=newnode
     return
   } 
   cur:=this.head
   for cur.Next!=nil{
        cur=cur.Next
   }   

   cur.Next=newnode
   
}
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    
    newnode:=&ListNode{Val:val}
    if index == 0 {
        newnode.Next = this.head
        this.head = newnode
        return
    }
    cur:=this.head
    i:=0
    for i<index-1 &&cur!=nil{
        
        cur=cur.Next
        i++
    } 

    if cur!=nil{
        
        newnode.Next=cur.Next
        cur.Next=newnode
    }


}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    
    if index==0{
        if this.head!=nil{
            this.head=this.head.Next
        }
        return
    }
    cur:=this.head
    i:=0
    for i<index-1&&cur!=nil{
         
         cur=cur.Next
         i++
    }

    if cur!=nil && cur.Next!=nil{

        cur.Next=cur.Next.Next
    }

}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
