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
