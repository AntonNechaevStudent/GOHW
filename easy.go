package main
import (
	"fmt"
	"math/big"
)
func main() {
	var n,i int64
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&n)
	for i=1; i<=n; i++ {
		if big.NewInt(i).ProbablyPrime(0) {
			fmt.Println(i)
		}
	}
	
}
	