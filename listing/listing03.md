Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
``` 

Ответ:
```
Вывод: nil, false
Такой вывод можно объяснить тем, что переменная err имплементирует интерфейс error, но изначально имеет тип *PathError
Интерфейс в языке Go представлен структурой:
type iface struct {
    tab  *itab
    data unsafe.Pointer
}
*itab - указатель на таблицу интерфейса, data - указатель на переменную, имплементирующую интерфейс

Пустой интерфейс представлен структурой:
type eface struct {
    _type *_type
    data  unsafe.Pointer
}
```