// Задание 1
// package main

// import "fmt"

// func main() {
// 	var a, b float64
// 	fmt.Print("Введите первую сторону: ")
// 	fmt.Scanln(&a)
// 	fmt.Print("Введите вторую сторону: ")
// 	fmt.Scanln(&b)
// 	fmt.Printf("Площадь прямоугольника: %f\n", a * b)
// }

// Задание 2

// package main

// import ("fmt"
// 		"math")

// func main()  {
// 	var a float64

// 	fmt.Print("Введите площадь круга: ")
// 	fmt.Scanln(&a)

// 	r := math.Sqrt(a / math.Pi)
// 	fmt.Printf("Диаметр окружности: %f\n", 2*r)
// 	fmt.Printf("Длина окружности: %f\n", (math.Pi * 2 * r))

// }

//Задание3

package main

import "fmt"
		

func main()  {
	var a int

	fmt.Print("Введите трехзначное целое число: ")
	fmt.Scanln(&a)
	b := a / 100
	c := (a % 100)/10
	d := a % 10 % 10
	fmt.Printf("Сотен: %d, Десятков: %d, Едениц: %d.",b,c,d)

}