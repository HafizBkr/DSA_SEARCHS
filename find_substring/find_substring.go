package main

import "fmt"

func findSubstring(str, substr string) bool {
    n := len(str)
    m := len(substr)

    for i := 0; i <= n-m; i++ {
        j := 0
        for j < m {
            if str[i+j] != substr[j] {
                break
            }
            j++
        }
        if j == m {
            return true
        }
    }
    return false 
}

func main() {
    str := "Hello, World!"
    substr := "World"
    found := findSubstring(str, substr)
    if found {
        fmt.Println("Substring found in string")
    } else {
        fmt.Println("Substring not found in string")
    }
}
