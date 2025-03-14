func maxArea(height []int) int {
    
    
    if(len(height)==0){
        
        return 0;
    }
    
    
    left:=0
    right:=len(height)-1
    maxarea:=0
    
    
    for left<right{
        
        minHeight:= min(height[left],height[right])
        
        width:=right-left
      
        area:=minHeight*width                  // height*width = area
       if maxarea<area{
            
            maxarea=area
        }
        
        
        if(height[left]<=height[right]){
            
            left++
        }else{
            
            right--
        }
        
        
        
        
    }
    
    
    return maxarea
    
}


func min(a,b int) int{
    
    if(a<b){
        
        return a
    }
    
    return b
    
    
}
