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
Output: error
```
err является интерфейсным типом, и при возвращении nil<br/>
с функции test error получает метаданные о customError, который реализует интерфейс error. <br/>
В таком случае выражение err != nil будет всегда возвращать true.
