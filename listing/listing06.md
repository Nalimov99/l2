Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
Output: [3 2 3]
```
Слайс это структура с указателем на массив.<br/>
append возвращает новый слайс, но если capacity слайса переполнена, аллоцируется <br/>
новый массив. В данном случае в строчке i[0] = "3", происходит изменение в массиве, <br/>
а после i = append(i, "4"), аллоцируется новый массив и слайс i указывает на него, поэтому <br/>
остальные изменение игнорируются.
