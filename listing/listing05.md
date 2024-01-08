Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод: error
Из функции test() возвращается указатель *customError со значением nil, затем приводится к интерфейсу error.
При сравнении с nil не равен ему, так как не является nil-интерфейсом
```