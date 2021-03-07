package main

import "fmt"

func main(){
    var n,sum int    
    fmt.Print("Enter a positive integer : ")
    fmt.Scan(&n)
    for i:=1; i<=n; i++ { 
        sum += (i*i*i)
    }
    fmt.Print("Sum : ",sum)
}