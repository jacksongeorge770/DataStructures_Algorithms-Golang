func dailyTemperatures(temperatures []int) []int {

n:=len(temperatures)
result:=make([]int,n)

stack:=[]int{}

for CurrDay,CurTemp:=range temperatures{
  for len(stack)>0 && temperatures[stack[len(stack)-1]]<CurTemp{
    prev:=stack[len(stack)-1]
    stack=stack[:len(stack)-1]
    result[prev]= CurrDay-prev
  }
  stack=append(stack, CurrDay)
}


return result
}
