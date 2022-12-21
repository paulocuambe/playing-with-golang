package main

import (
	"fmt"
	"os"
)

func mustErr(err error){
	if err != nil {
		panic("happened here" + err.Error())
	}
}

func Cat(path string) {
	f, err := os.Open(path)
	mustErr(err)

	c := make([]byte, 5)
	_, err = f.Read(c)

	mustErr(err)

	for v, k := range c {
		fmt.Println(v, k)
	}

	fmt.Println(c)

	defer f.Close()
}

func main() {
	fmt.Println(os.Args)
	cmdList := os.Args[1:]
	cmd := os.Args[1]
	fmt.Println("cmd: ", cmd, cmdList)

	switch cmd {
	case "cat":
		Cat(cmd)
	case "mv":
		fmt.Println("mv")
	case "rm":
		fmt.Println("rm")
	case "echo":
		fmt.Println("echo")
	case "touch":
		fmt.Println("touch")
	default:
		fmt.Println("The", cmd, "command is not supported")
	}
}
