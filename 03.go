package main

import "fmt"

// ДЗ 02.03 Определение разрядов трёхзначного числа

func main() {
    var num int
    fmt.Println("Введите трёхзначное число")
    fmt.Scanln(&num)
    fmt.Println("Количество сотен:", num / 100)
	num = num - num / 100 * 100
    fmt.Println("Количество десятков:", num / 10)
	num = num - num / 10 * 10
	fmt.Println("Количество единиц:", num)
}