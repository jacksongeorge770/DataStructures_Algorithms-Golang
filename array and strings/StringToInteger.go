package main

import (
	"math"
)

func myAtoi(s string) int {
    i := 0
    for i < len(s) && s[i] == ' ' {
        i++
    }

    if i == len(s) {
        return 0
    }

    sign := 1
    if s[i] == '-' {
        sign = -1
        i++
    } else if s[i] == '+' {
        i++
    }

    result := 0
    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        digit := int(s[i] - '0')
        result = result*10 + digit

        if result*sign > math.MaxInt32 {
            return math.MaxInt32
        }
        if result*sign < math.MinInt32 {
            return math.MinInt32
        }

        i++
    }

    return sign * result
}

