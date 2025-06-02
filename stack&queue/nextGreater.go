func nextGreaterElement(nums1 []int, nums2 []int) []int {
 
 stack:=[]int{}
 Next:=make(map[int]int)

 for _,num:=range nums2{
    
  for len(stack)>0 && stack[len(stack)-1]<num{
    top:=stack[len(stack)-1]
    stack=stack[:len(stack)-1]
    Next[top]=num
  }
 stack=append(stack, num)
  }

result:=make([]int,len(nums1))

for i,num:= range nums1{
  
  if val,found:=Next[num];found{
    
    result[i]=val
  }else{
  result[i]=-1
  }
}
return result

}
