
import (
	"math"
)
func largestRectangleArea(heights []int) int {
    
   stack:=[]int{}
   MaxArea:=0

   for i:=0;i<len(heights);i++{
      
      for len(stack)>0 && heights[stack[len(stack)-1]]>heights[i]{
      top:=heights[stack[len(stack)-1]]
      stack=stack[:len(stack)-1] 
      width:=i
      if len(stack)>0{
           width=i-stack[len(stack)-1]- 1
         
         }
         area:=width*top
         MaxArea=int(math.Max(float64(MaxArea),float64(area)))
               }
      stack=append(stack,i)
   }
 for len(stack)>0{

        height:=heights[stack[len(stack)-1]]
        stack=stack[:len(stack)-1]
        width:=len(heights)
        
        if len(stack)>0{
            width=len(heights)-stack[len(stack)-1]-1
        }
       MaxArea=int(math.Max(float64(MaxArea),float64(width*height)))
      }
   return MaxArea

}
