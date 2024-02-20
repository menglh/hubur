# fdmax
小型 Helper 库，可自动增加当前 go 程序的最大文件描述符数。
可以简单地导入如下：

```
package main

import (
	"fmt"

	_ "github.com/projectdiscovery/fdmax/autofdmax"
)

func main() {
	fmt.Println("test")
}
```
