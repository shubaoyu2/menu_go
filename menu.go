package main
 
import "fmt"
 
func main() {
	for {
		var cmd string
		fmt.Println("cmd:")
		fmt.Scanln(&cmd)
		if "help" == cmd {
			fmt.Println("this is a help function ")
		} else if "exit" == cmd {
			fmt.Println("exit")
			break
		} else {
			fmt.Println("wrong command")
		}
	}
}