func twoSum(nums []int, target int) []int {
    
     hash:=make(map[int]int)   //this is hash maker
    
    for i,integer:= range nums{
        
        
        ans:=target-integer
        
        if j,exist:= hash[ans];exist{
            
            
            return []int{j,i}
        }
        
        
        hash[integer]=i //putting key and value to hash map
    }

 return nil
}
