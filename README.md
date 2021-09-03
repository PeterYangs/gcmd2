#gcmd2

安装
```shell
go get github.com/PeterYangs/gcmd2
```

阻塞获取输出
```go
package main

import (
	"context"
	"fmt"
	"github.com/PeterYangs/gcmd2"

	"log"
)

func main() {

	cmd := gcmd2.NewCommand("dir", context.TODO())

	out, err := cmd.CombinedOutput()

	if err != nil {

		log.Println(err)

		return

	}

	fmt.Println(string(out))

}

```

阻塞获取实时输出
```go
package main

import (
	"context"
	"github.com/PeterYangs/gcmd2"
	"log"
)

func main() {

	cmd := gcmd2.NewCommand("php index.php", context.TODO())

	err := cmd.Start()

	if err != nil {

		log.Println(err)

		return

	}
}

```

非阻塞（后台运行模式）
```go
package main

import (
	"context"
	"github.com/PeterYangs/gcmd2"

	"log"
)

func main() {

	cmd := gcmd2.NewCommand("php index.php", context.TODO())

	err := cmd.StartNoWait()

	if err != nil {

		log.Println(err)

		return

	}
}

```


