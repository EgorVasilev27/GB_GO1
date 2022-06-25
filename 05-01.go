package main

import "fmt"

// ДЗ 05.01 Числа Фибоначчи с сохранением в словарь
func main () {
	numsArr := map[int]int {
		1 : 0,
		2 : 1,
	}

	num := 0
	maxInd := 2
	var userResponse string

	for  ;;  {
		fmt.Println ("Введите порядковый номер числа")
		fmt.Scanln (&num)
		
		if num <= maxInd {
			fmt.Println("Есть такой, это", numsArr[num])
		} else {
			fmt.Println("Нет такого, сейчас посчитаем")
			for i := maxInd + 1 ; i <= num ; i++ {
				numsArr[i] = numsArr[i-1] + numsArr [i-2]
			}
			maxInd = num
			fmt.Println("Посчитали, это", numsArr[maxInd])
		}
		fmt.Println(numsArr)

		
		fmt.Println ("Ещё разок? (y/n)")
		fmt.Scanln (&userResponse)

		if userResponse != "y" {
			break
		}
	}

}