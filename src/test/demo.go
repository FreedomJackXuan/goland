package test
import (
	"errors"
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func Demo() {
	flag.Parse()
	greeting, err := hello(name)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Println(greeting, it())
}

// hello 用于生成问候内容。
func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}

// introduce 用于生成介绍内容。
func it() string {
	return "Welcome to my Golang column."
}