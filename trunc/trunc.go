package main
import "fmt"
func main(){
 fmt.Printf("Enter the float number")
 var flnum float32
 fmt.Scan(&flnum)
 var inum int
 inum=int(flnum)
 fmt.Print(inum)
 }