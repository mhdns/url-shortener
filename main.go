package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {

	// type url struct {
	// 	URL  string `yaml:"path"`
	// 	Path string `yaml:"url"`
	// }

	// type urls struct {
	// 	URLList []url `yaml:"PathLists"`
	// }
	// res := urls{}
	file, err := os.Open("urls.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println(string(data))

	// type URLPair struct {
	// 	Path string
	// 	URL  string
	// }

	// type URLList struct {
	// 	List []URLPair
	// }

	// var urls URLList

	var urls []interface{}

	err = yaml.Unmarshal(data, &urls)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", urls)

	for i, v := range urls {
		fmt.Println(i, v)
		fmt.Printf("%T\n", v)
	}

	// http.HandleFunc("/", helloworld)
	// http.ListenAndServe(":5000", nil)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func exampleYaml() {
	type N struct {
		N int
		C int
	}

	type T struct {
		F []N `yaml:"a,omitempty"`
		B int
	}

	var t T
	fmt.Println("a: \n  - n: 1\n    m: 2\nb: 2")
	yaml.Unmarshal([]byte("a: \n  - n: 1\n    m: 2\nb: 2"), &t)

	fmt.Println(t)
}
