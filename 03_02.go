package main

import "fmt"
//import "math"

// ДЗ 03.02 Простые числа
// Попробую реализовать Эратосфена
// Для остальных методов не хватает базы математической

func main() {
    var a int
    fmt.Println("Введите число!")
    fmt.Scanln(&a)
	nums:= [] int {1} 
	
	for i:=2; i<=a; i++ {
			nums = append(nums, i)
//			fmt.Println(nums[i-1])
		}
//	fmt.Print(nums)
//	fmt.Print(nums[0])
	for i:=1; i<a-1; i++ {
//		fmt.Println(nums[i])
		if nums[i] != 1 {
			for j:=i+1; j<a; j++ {
//				fmt.Print(nums[j])
//				fmt.Print(" ")
				if (nums[j] % nums[i] == 0) {
					nums[j] = 1
//					fmt.Print("yes ")
				}
			}
//			fmt.Println("")
//			fmt.Print(nums)
		}
	}
	
	fmt.Println("Простые числа:")
	for i:=1; i<a-1; i++ {
		if nums[i] != 1 {
			fmt.Println(nums[i])
		}
	}
	
}

