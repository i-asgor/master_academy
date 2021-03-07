package main

import "fmt"

func main() {
    var n, sum, num int
    
    fmt.Scan(&n)
    
    for i := 0; i < n; i++ {
        fmt.Scan(&num)
        sum += num
    }
    
    fmt.Println(sum)
}