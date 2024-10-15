package main

import ("fmt"
    "example/greetings"
)

func main() {

    fmt.Println("Hello, World!")

    message := greetings.Greet("murugan")

    fmt.Println(message)
}
