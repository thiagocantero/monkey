package main
import (
"fmt"
"os"
"os/user"
"monkeylang/repl"
)
func main() {

	user, err := user.Current()

	if err != nil {

		panic(err)

	}

	fmt.Printf("Olá %s! Esta é Linguagem de Programação Monkey!\n",
	user.Username)

	fmt.Printf("Sinta-se à vontade para digitar comandos\n")

	repl.Start(os.Stdin, os.Stdout)
}