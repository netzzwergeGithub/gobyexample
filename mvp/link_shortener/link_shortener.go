package main

import (
	"fmt"
	"net/http"
	"strings"
)

type storage interface {
	Query(string) (string, error)
}

type configH struct {
	stg     storage
	address string
}

func main() {
	var sstg = SimpleStorage{
		data: map[string]string{
			"go":          "https://golang.de",
			"python":      "https://www.python.org/",
			"pytorch":     "https://pytorch.org/",
			"huggingface": "https://huggingface.co/",
		},
	}
	Config := configH{
		address: "127.0.0.1:8080",
		stg:     sstg,
	}
	run(Config)
}

func run(Config configH) {

	handleRequest := func(writer http.ResponseWriter, request *http.Request) {
		splitted := strings.Split(request.URL.Path, "/")
		// fmt.Println(splitted, len(splitted))
		if len(splitted) == 3 && splitted[1] == "r" {
			link, err := Resolve(splitted[2], Config.stg)

			if err != nil {
				fmt.Println(err)
				writer.WriteHeader(http.StatusBadRequest)
			} else {
				fmt.Fprintln(writer, link)
			}

		} else {
			writer.WriteHeader(http.StatusNotFound)

		}

	}
	handler := http.HandlerFunc(handleRequest)
	// create an http-Server
	http.ListenAndServe(Config.address, handler)

}

func Resolve(index string, storage storage) (string, error) {
	return storage.Query(index)
}

type SimpleStorage struct {
	data map[string]string
}

func (storage SimpleStorage) Query(index string) (string, error) {
	link, ok := storage.data[index]
	if !ok {
		return link, fmt.Errorf("storage: no no data for index: %s", index)
	}
	return link, nil
}
