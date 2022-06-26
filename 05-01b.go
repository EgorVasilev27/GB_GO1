package main

import "fmt"

// ДЗ 05.01 Числа Фибоначчи с сохранением в словарь
func main () {
	
	numsArr := map[int]int {
		1 : 0,
		2 : 1,
	}
	num := 0
	userResponse := "y"

	for  ;;  {
		fmt.Println ("Введите порядковый номер числа")
		fmt.Scanln (&num)
		_, indexExists := numsArr[num]
	
		if indexExists {
			fmt.Println("Есть такой, это", numsArr[num])
		} else {
			fmt.Println("Нет такого, сейчас посчитаем")
			numsArr[num] = getFibo(num)
			fmt.Println("Посчитали, это", numsArr[num])
		}
		
		fmt.Println(numsArr)		
		fmt.Println ("Ещё разок? (y/n)")
		fmt.Scanln (&userResponse)

		if userResponse != "y" {
			break
		}
	}

}

//Функция расчёта числа Фибоначчи
func getFibo(num int) int {
	var a int
	if num > 2 {
		a = getFibo(num-1)+getFibo(num-2)
	} 
	if num == 2 {
		a = 1
	}
	if num == 1 {
		a = 0
	}
	return a
}