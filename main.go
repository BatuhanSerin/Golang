package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// type movie struct{
// 	name string
// }

func main() {

	names := readFile()

	if os.Args[1] == "list" {

		listCommand(names)

	} else if os.Args[1] == "search" {

		searchCommand(names)

	}

}

//readFile() function reads json file is named "data.json", checks for errors then returns data is read from json file
func readFile() map[string]interface{} {

	data := map[string]interface{}{}
	contents, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(contents, &data); err != nil {
		panic(err)
	}

	return data["psychological thriller movies"].(map[string]interface{})
}

//listCommand() function is called after "list" command, prints movie list.
func listCommand(names map[string]interface{}) {
	fmt.Println("**************Movie List**************")
	for _, movieNames := range names {

		movieNames := fmt.Sprintf("%v", movieNames)

		fmt.Println(movieNames)
	}
}

//searchCommand() function is called after "search" command, checks and prints if the searched movie is in the list.
func searchCommand(names map[string]interface{}) {
	args := ""

	for i := 2; i < len(os.Args); i++ {
		args = args + " " + string(os.Args[i])
	}
	args = strings.Title(strings.ToLower(args))[1:]

	flag := []bool{true}
	for _, movieNames := range names {

		movieNames := fmt.Sprintf("%v", movieNames)

		if movieNames == args {
			fmt.Printf("The movie is found: %s", movieNames)
			flag[0] = false
			break
		}

	}
	if flag[0] {
		fmt.Println("The movie is not found!")
	}
}
