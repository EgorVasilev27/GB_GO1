package main

import "fmt"

// ДЗ 02.01 Расчёт площади прямоугольника

func main() {
    var a int
    var b int
    fmt.Println("Введите стороны прямоугольника через пробел")
    fmt.Scanln(&a, &b)
    fmt.Println("Площадь:", a * b)
}