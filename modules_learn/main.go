/*
	root directory files should always have package main
	Initialize a module using go mod init name
	go run -> will execute a file or program not all files when we importing from another files, so use build
	Build the module using go build -o name

	User -> public access
	user -> private access, but public in own package
*/

package main

import (
	"fmt"
	"modules_learn/types"
	"modules_learn/util"
)

func main() {

	user := types.User{
		Name: util.GetName(),
		Age:  util.GetAge(),
	}
	fmt.Println(user)
}
