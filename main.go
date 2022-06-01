package main
import (
	"fmt"
	"os"
	"math"
)
func Factorial(n uint)(result uint) {
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
func main() {
	var resint uint
	var a, b, res float64
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, pow, factorial): ")
	fmt.Scanln(&op)
	if (b==0 && op == "/") {
		fmt.Println("Операция выбрана неверно, на ноль делить нельзя")
		os.Exit(1)
	}
	switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			res = a / b
		case "pow":
			res = math.Pow(a, math.Trunc(b))
		case "factorial":
			resint = Factorial(uint(a))
		default:
			fmt.Println("Операция выбрана неверно")
			os.Exit(1)
		}
	if op == "factorial" {
		fmt.Printf("Факториал из первого аргумента: %d\n", resint)
		} else {fmt.Printf("Результат выполнения операции: %.2f\n", res)}
	}
	