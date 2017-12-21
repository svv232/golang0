package main

import "fmt"
import "sort"


func main(){
    nums := []int{9,2,4,39,34,32,2}
    sort.Ints(nums)
    fmt.Println("Ints:  ", nums)
    s := sort.IntsAreSorted(nums)
    fmt.Println("Sorted:", s)
}
