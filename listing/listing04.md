Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Вывод: 0 1 ... 9, затем deadlock
for ... range заблокирует выполнение main, так как канал ch не был закрыт
```