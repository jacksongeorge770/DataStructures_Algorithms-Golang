func longestValidParentheses(s string) int {
    
  stack:=[]int{-1}
  maxLen:=0
  for i,ch:=range s{
    
    if ch =='('{
        stack=append(stack,i)

  }else{
    stack=stack[:len(stack)-1]
    if len(stack)==0{
        stack=append(stack, i)
    }else{
   curLen:=i-stack[len(stack)-1]
   
   if curLen>maxLen{
    maxLen=curLen
   }
   
   


    }
   
  }
  }
return maxLen
}
