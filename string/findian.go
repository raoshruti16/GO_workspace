package main
import( "fmt"
	"strings"
	)
func main(){
 fmt.Printf("Enter the string")
 var s string
 s=strings.ToLower(s)
 fmt.Scan(&s)
 if strings.HasPrefix(s,"i"){
	if strings.HasSuffix(s,"n"){
		if strings.Contains(s,"a"){
			fmt.Printf("Found")
		}else { 
fmt.Printf("Not Found") 
}
           }else { 
fmt.Printf("Not Found") 
}
    }else { 
fmt.Printf("Not Found") 
}
		
		
 }