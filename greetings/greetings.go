package greetings
import "fmt"

func Greet(name string) string {
	message  := fmt.Sprintf("Hi %v welcome", name)
	return message
}

