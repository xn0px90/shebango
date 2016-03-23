# shebango

System administrator will  write shell scripts to accomplish certain task.
Many times we encounter scripting languages have limited capabilities.
For example establishing secure transport layer or upload files to a remote location securely.
Using Go to replace shell scripts can be achieved easily .

# Example: ###To execute Go program as a shell script.

```
/*Put this line at the top of your code*/
//usr/bin/env go run $0 $@ ; exit

package main
import "fmt"
func main() {
    fmt.Printf("hello, world\n")
}
```

# Make it an executable

```
chmod +x hello.go
./hello.go
hello, world
```
#There you go !
