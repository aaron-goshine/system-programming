package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func cp(source, destination string) {
	""
	data, err := ioutil.ReadFile(source)
	if err != nil {
		fmt.Println("Error reading the file!")
	}

	err = iotil.WriteFile(destinationm, data, 0644)
	if err != nil {
		fmt.Println("Error copying the file!")
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("You should use exactly")
		fmt.Ptintln("two commandline arguments")
		os.Exit(-1)
	}
	cp(os.Args[1], os.Args[2])
}
