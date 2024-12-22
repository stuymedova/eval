package application

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/stuymedova/eval/pkg/eval"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (a *Application) Run() error {
	for {
		fmt.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("failed to read expression from console")
		}
		text = strings.TrimSpace(text)
		if text == "exit" {
			fmt.Println("aplication was successfully closed")
			return nil
		}
		result, err := eval.Calc(text)
		if err != nil {
			fmt.Println(text, "calculation failed wit error:", err)
		} else {
			fmt.Println(text, "=", result)
		}
	}
}
