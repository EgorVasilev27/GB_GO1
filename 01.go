package main

import "fmt"
import "math"

// ДЗ 03.01 Калькулятор с валидацией

func main() {
	fmt.Println("Введите первое число! (Целое)")
	var a1 = getNum ()
	fmt.Println("Введите второе число! (Целое)")
	var a2 = getNum ()
	fmt.Println("Введите операцию:")	
	fmt.Println("+: сумма")
	fmt.Println("-: разность")
	fmt.Println("*: произведение")
	fmt.Println("/: частное")
	fmt.Println("**: а в степени b")
	var oper string
	fmt.Scanln(&oper)
	fmt.Println("Результат:")
	switch oper {
		case "+":
			fmt.Println(a1 + a2)
		case "-":
			fmt.Println(a1 - a2)
		case "*":
			fmt.Println(a1 * a2)
		case "/":
			fmt.Println(a1 / a2)
		case "**":
			fmt.Println(math.Pow(a1,a2))						
		default:
			fmt.Println("Вы ввели что-то не то")
	}
}

//Функция проверяет, что пользователь вводит целое число
func intValidator(num string) (float64 , int) {
    var nums =  "0123456789"
    var isFlag = 0
    var numFloat = 0.0
    numsInt:= [10]float64{0,1,2,3,4,5,6,7,8,9}

    for i:=0; i<len(num); i++ {
        isFlag = 0
        for j:=0; j<10; j++ {
            if num[i] == nums[j] {
                isFlag = 1
                numFloat = numFloat + numsInt[j]*math.Pow(10,float64(len(num) - i - 1))
            }
        }
        if isFlag == 0 {
            isFlag = 3
            break
        }
    }
    return numFloat, isFlag
}

//Функция организует ввод числа
func getNum() float64 {
    var a string
    var aFloat = 0.0
    var isFlag = 0
	for i:=0; i<2; {
    	aFloat = 0
    	fmt.Scanln(&a)
    	aFloat, isFlag = intValidator(a)
		if isFlag == 3 {
			fmt.Println("Вы ввели что-то не то, повторите")
		}
		if isFlag == 1 {break}
	}
	return aFloat
}



