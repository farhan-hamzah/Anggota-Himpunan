package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i, j int
	var valid bool
	valid = true
	fmt.Scan(&n)

	for i =0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n; i++{
		for j= i+1; j < n; j++{
			if A[i] == A[j]{
				valid = false
			}
		}
	}
	if valid == true{
		fmt.Print("Adalah valid")
	}else{
		fmt.Print("Tidak valid")
	}
}



