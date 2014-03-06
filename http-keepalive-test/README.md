request.Close属性在golang 1.2没有作用，做了个测试telnet，后来换成Header的Connection close好了，1.3还没测试，出了后再实验。  
===========

```
package main

import (
        "fmt"
        "net/http"
        "runtime"
)

type Hello struct{}

func (h Hello) ServeHTTP(
        w http.ResponseWriter,
        r *http.Request) {
        r.Close = true // 这行代码写不写都一样
        w.Header().Set("Connection", "close")
        fmt.Fprint(w, "hello!\r\n")
}

func main() {
        runtime.GOMAXPROCS(runtime.NumCPU())
        var h Hello
        http.ListenAndServe("localhost:4000", h)
}
```

测试方法
=======
telnet localhost 4000  
然后，注意要有回车换行
```
GET / HTTP/1.1
User-Agent: curl/7.30.0
Host: localhost:4000
Accept: */*


```