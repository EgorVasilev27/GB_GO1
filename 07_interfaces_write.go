package main
import "fmt"
import "io"

//		type
type buffer []byte
//		funcs

/*func (b buffer) Read (buf []byte) (n int, err error) {
	n = len(buf)
	err = nil
	return
}*/

func (b buffer) Write (buf []byte) (n int, err error) {
	for i:=0; i<len(buf); i++ {
		b[i]=buf[i]
	}
	n = len(buf)
	b[n] = '~'							//в качестве "EOF" использую тильду,
	fmt.Println(b)						//чтобы в Reader'е читать до неё
	err = nil
	return
}

//		interfaces
type Reader interface {
	Read ([]byte) (int, error)
}
type Writer interface {
	Write([]byte) (int, error)
}
//		main
func main () {
var b buffer 							//буфер
var bufSize int							//размер буфера
var s string							//вводимая строка
fmt.Print("Buffer size:")	
fmt.Scanln(&bufSize)	
b = make([]byte,bufSize+1,bufSize+1)	//делаем буфер с ячейкой для "EOF"	
b1 := b									//проверочный буфер
fmt.Println("buffer:", b)
fmt.Print("Enter the string:")
fmt.Scanln(&s)
fmt.Println("Input:", s)
data := []byte(s)
fmt.Println("byte:", data)
b.Write(data)
fmt.Println("buffer filled directly:", b)
n, err := io.WriteString(b1,s)
fmt.Println(n, err)
fmt.Println("buffer filled with WriteString:",b1)
}