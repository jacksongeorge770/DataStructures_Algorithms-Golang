func removeDuplicateLetters(s string) string {
    
    stack := []byte{}
    last:=make(map[byte]int)
    for  i:=0;i< len(s);i++{
       last[s[i]]=i
    }
    seen:=make(map[byte]bool)
    for i:=0;i<len(s);i++{
        c:=s[i]

        if seen[c] {
            continue
        }

        for len(stack)>0&& c < stack[len(stack)-1] && last[stack[len(stack)-1]]>i {
                 
                top:=stack[len(stack)-1]
                seen[top]=false
                stack=stack[:len(stack)-1]

        }

        stack = append(stack, c)
		seen[c] = true
                                    

    }

return string(stack)
}
