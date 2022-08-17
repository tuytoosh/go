# Learn Go language

Hello world program in Go

```
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

### Get user input
The `Scan` function gets user input and we have to send a `pointer` of the variable to it.
```
fmt.Println("Enter ticket count: ")
fmt.Scan(&ticketCount)
```

### Arrays and Slices
Arrays in Go have fixed size and if we try to append elements to the end of that we have to know exact position of new spot.
```
var users = [3] string {"Hamid", "Amir"}
users[3] = "Majid"
```

Slices are better because they have dynamic size and we can `append` any number of elements to them.
```
var users[]string
users = append(users, "Hamid")
users = append(users, "Amir")
```

### Conditions and Loops

#### while

The simplest loop in go is `for{}` which is eqivalant to `while(){}` in other languages.

```

var users []string

for {
    var name string
    fmt.Println("Enter user or `exit`: ")
    fmt.Scan(&name)
    if (name == "exit") {
        break
    }
    users = append(users, name)
}

fmt.Printf("%v\n", users)
```
#### foreach
```
for index, name := range users {
    fmt.Print(index, name)
}
```
You can simply left the `index` blank buy puttig `_` in its location:

```
for _, name := range users {
    fmt.Print(name)
}
```

## Functions
Functions are very similar to other languages:
```
func bye(name string) {
	fmt.Printf("Bye %v\n!", name)
}
```
And for calling them:
```
bye("Hamid")
```
