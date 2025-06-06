func romanToInt(s string) int {
    // Define the mapping of Roman numeral characters to their integer values
    hashmap := map[byte]int{
        'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
    }
    
    total := 0
    // Loop through the string
    for i := 0; i < len(s); i++ {
        // Check if the current numeral is smaller than the next numeral
        if i+1 < len(s) && hashmap[s[i]] < hashmap[s[i+1]] {
            // If smaller, subtract the current numeral from total
            total -= hashmap[s[i]]
        } else {
            // Otherwise, add the current numeral to total
            total += hashmap[s[i]]
        }
    }
    
    return total
}
