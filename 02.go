package main

import "fmt"
import "math"

// ДЗ 02.02 Расчёт диаметра и длины окружности через площадь круга

func main() {
    var s float64
    fmt.Println("Введите площадь круга")
    fmt.Scanln(&s)
    fmt.Println("Диаметр:", 2 * math.Sqrt (s/math.Pi))
	fmt.Println("Длина окружности:", 2 * math.Sqrt (s/math.Pi) * math.Pi)
}