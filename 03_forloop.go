package main
import "fmt"
func main() {

idx:=1

for idx<=10{
	
	fmt.Println(idx)
	idx++
} 

fmt.Println(idx)

if idx>10{
	fmt.Println("End of Loop")
} else {
	fmt.Println("Exited before the Loop")
}


for idx=1;idx<10;idx++{
	fmt.Println(idx)
}

}