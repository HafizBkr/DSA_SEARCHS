package linearsearch

import "fmt"

func LinearSearch(arr []int, target int) int {
    for i, num := range arr {
        if num == target {
            return i
        }
    }
    return -1 
}

func main() {
    arr := []int{12, 45, 67, 89, 34}
    target := 12
    result := LinearSearch(arr, target)
    if result != -1 {
        fmt.Println("Element found at index:", result)
    } else {
        fmt.Println("Element not found in array")
    }
}
