package main 

import "fmt"
import "config/plugin"

func main() {
    
    conf := config.GetConf("https://gist.githubusercontent.com/rumyantseva/26bee59a04d416c55e0e7e8155717d59/raw/02d603e96b31158edbfe46e6655db51ad51a2d3c/conf.yaml")
    
    fmt.Printf("%+v", conf)
    confValid := config.ValidUrl(conf)
    fmt.Printf("%+v\n",confValid)
    config.PrintConfig(conf)
    config.PrintConfig(confValid)

}






































// package main

// import (
//     "fmt"
//     "io/ioutil"
//     "log"

//     "gopkg.in/yaml.v3"
// )

// type myData struct {
//     Conf struct {
//         Port      int
//         Hits      int64
//         Time      int64
//         CamelCase string `yaml:"camelCase"`
//     }
// }

// func readConf(filename string) (*myData, error) {
//     buf, err := ioutil.ReadFile(filename)
//     if err != nil {
//         return nil, err
//     }

//     c := &myData{}
//     err = yaml.Unmarshal(buf, c)
//     if err != nil {
//         return nil, fmt.Errorf("in file %q: %v", filename, err)
//     }

//     return c, nil
// }

// func main() {
//     c, err := readConf("https://gist.githubusercontent.com/rumyantseva/26bee59a04d416c55e0e7e8155717d59/raw/02d603e96b31158edbfe46e6655db51ad51a2d3c/conf.yaml")
//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Printf("%v", c)
// }
 
 
 
 // package main

// import (
//     "fmt"
//     "gopkg.in/yaml.v2"
//     "io/ioutil"
//     "log"
// 	"os"
// )

// type conf struct {
//     Port int `yaml:"port"`
//     // Time string `yaml:"time"`
// }

// func (c *conf) getConf() *conf {

//     yamlFile, err := ioutil.ReadFile("https://gist.githubusercontent.com/rumyantseva/26bee59a04d416c55e0e7e8155717d59/raw/02d603e96b31158edbfe46e6655db51ad51a2d3c/conf.yaml")
//     if err != nil {
//         log.Printf("yamlFile.Get err   #%v ", err)
//     }
//     err = yaml.Unmarshal(yamlFile, c)
//     if err != nil {
//         log.Fatalf("Unmarshal: %v", err) 
//     }
	

//     return c
// }

// func main() {
//     var c conf
//     c.getConf()

//     fmt.Println(c)
// }
 
	
 

































// package main
// import "fmt"
// func main() {
// 	var a uint32 = 1
// 	var b uint32 = a // взятие указателя на a. Тип указателя на тип T — *T
// 	fmt.Println("b =", b)
// 	c := 3 *b // взятие значения по указателю
// 	fmt.Println("c =", c)
// }



// package main
// import (
// 	// "bufio"
// 	"fmt"
// 	// "os"
// 	// "strconv"
 	
// )

// var m map[uint]uint64
// func mapper (k uint64, keys uint)uint64{
// 	colorRed := "\033[31m"
// 	colorReset := "\033[0m"
// 	fmt.Print(string(colorRed),keys," ",string(colorReset))
// 	fmt.Print(string(colorReset),k," ")
// 	m[keys] = k
// 	return k	
// }
// func fibo(i uint)uint64{
// 	if i == 0 {return 0}
// 	if i == 1 {return 1}
// 	_, ok := m[i]
// 		if ok {return m[i]}	
// 	return mapper((fibo(i-1) + fibo(i-2)),i)
// }

// func main() {
// 	m = make(map[uint]uint64)
// 	var a uint
// 	for {
// 		fmt.Printf("\nВведите порядковый номер числа в последовательности: ")
//  		n, err := fmt.Scanln(&a)
// 		 if err != nil {
// 			fmt.Println(n, err)
// 			return
// 		}
// 		_, ok := m[a]
// 		if ok {
// 			fmt.Printf("%d найдено в мапе %d\n",a,m[a])
// 			} else {fibo(a)}
				
// 	}
		
// }
















// package main
// import (
// 	"fmt"
// 	"os"
// 	"math"
// )
// func Factorial(n uint)(result uint) {
// 	if (n > 0) {
// 		result = n * Factorial(n-1)
// 		return result
// 	}
// 	return 1
// }
// func main() {
// 	var resint uint
// 	var a, b, res float64
// 	var op string
// 	fmt.Print("Введите первое число: ")
// 	fmt.Scanln(&a)
// 	fmt.Print("Введите второе число: ")
// 	fmt.Scanln(&b)
// 	fmt.Print("Введите арифметическую операцию (+, -, *, /, pow, factorial): ")
// 	fmt.Scanln(&op)
// 	if (b==0 && op == "/") {
// 		fmt.Println("Операция выбрана неверно, на ноль делить нельзя")
// 		os.Exit(1)
// 	}
// 	switch op {
// 		case "+":
// 			res = a + b
// 		case "-":
// 			res = a - b
// 		case "*":
// 			res = a * b
// 		case "/":
// 			res = a / b
// 		case "pow":
// 			res = math.Pow(a, math.Trunc(b))
// 		case "factorial":
// 			resint = Factorial(uint(a))
// 		default:
// 			fmt.Println("Операция выбрана неверно")
// 			os.Exit(1)
// 		}
// 	if op == "factorial" {
// 		fmt.Printf("Факториал из первого аргумента: %d\n", resint)
// 		} else {fmt.Printf("Результат выполнения операции: %.2f\n", res)}
// 	}
	