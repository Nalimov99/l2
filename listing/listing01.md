Что выведет программа? Объяснить вывод программы.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
Output: [77 78 79]
```
Создаеться слайс b в который включаються элементы начаная<br/>
с индекса 1 (77), заканчивая индексом 3 (79). [1,4)