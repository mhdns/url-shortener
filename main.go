package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type paths struct {
	Path string
	URL  string
}

func main() {
	paths := getPaths("urls.yaml")

	// http.Redirect(w, r, newUrl, http.StatusSeeOther)
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloworld)
	for _, v := range paths {
		// fmt.Println(v.Path)
		mux.HandleFunc(v.Path, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, v.URL, http.StatusSeeOther)
		})
	}
	// fmt.Print(mux)
	http.ListenAndServe(":5000", mux)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func getPaths(filename string) []paths {
	// Read File
	file, err := os.Open("urls.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	// Parse Yaml
	var urls []paths

	err = yaml.Unmarshal(data, &urls)
	if err != nil {
		fmt.Println(err)
	}

	return urls
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
