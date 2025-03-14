func lengthOfLongestSubstring(s string) int {

  //first make hashmap 

  if(len(s)==0) {
    return 0;
  }


  
  hashmap:= make(map[rune]int)
   
  left:=0
  maxlength:=0
  
  for right,char :=range s{

     if previdx,found := hashmap[char] ;found && previdx>=left{

        left=previdx+1
     }

     hashmap[char]=right

     if(maxlength<right-left+1){

        maxlength=right-left+1
     }



  }


  return maxlength




    
}


// brute force approch


func lengthOfLongestSubstring(s string) int {
    maxLength := 0

    // Iterate over all possible starting points for substrings
    for i := 0; i < len(s); i++ {
        seen := make(map[rune]bool) // Hash map to track characters in the current substring
        for j := i; j < len(s); j++ {
            char := rune(s[j])
            // If the character is already in the current substring, break out of the loop
            if seen[char] {
                break
            }
            // Add character to the seen map
            seen[char] = true
            // Update the maxLength if the current substring is longer
            maxLength = max(maxLength, j-i+1)
        }
    }
    return maxLength
}

// Helper function to return the maximum of two numbers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

