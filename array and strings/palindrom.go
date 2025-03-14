func isPalindrome(x int) bool {
    
orginal:=x
reversed:=0

for x>0{
  
  digit:=x%10
  reversed=10*reversed+digit
  x=x/10

}

if orginal==reversed{
    return true
}

return false
}
