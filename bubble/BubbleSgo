package main
import (
    "fmt"  
)
func main() {
    var k int
    fmt.Printf("Enter a number:")
    fmt.Scan(&k)
    kappa := make([]int,0, 10)
    var n int
    i:=0
    for i<k{
        if(i==10){
            break
        }
        fmt.Scan(&n)
        kappa = append(kappa, n)
        
        i++
    }
    bubbleSort(kappa)
        fmt.Print(kappa)
     
}
func swap(kaim []int, i int, k int){
    temp:=kaim[i]
    kaim[i]=kaim[k]
    kaim[k]=temp
}
func bubbleSort(kaim []int){
    x:= len(kaim)
    for i:=0;i<x;i++{
        for k:=0;k<x-i-1;k++{
            if(kaim[k]>kaim[k+1]){
                swap(kaim,k,k+1)
            }
        }
    }
}