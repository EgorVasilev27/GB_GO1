package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

// ДЗ 04.01 Сортировка вставкой


/* Ввод массива */
func main() {
	fmt.Println("Введите массив целых чисел через пробел!")
	var stringLine string

in := bufio.NewScanner(os.Stdin)
in.Scan()

stringLine = in.Text()
//fmt.Println(stringLine)

stringArray := strings.Split(stringLine, " ")
//fmt.Println(stringArray)

var intsArray []int
for _, s := range stringArray {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("atata!")
 		} 
  intsArray = append(intsArray, num)
	}

/* Собственно сортировка */
for i:=1; i<len(intsArray); i++ {
	index := intsArray[i]	
	flag := 0
	for j:=i-1; j>=0; j-- {
		if intsArray[j]>index {
			intsArray[j], intsArray[j+1] = intsArray[j+1], intsArray[j]
			flag = 1
			}
		if flag == 0 {
//			fmt.Println("Break!")
			break
			}
		}
	fmt.Println(intsArray)
	}
}