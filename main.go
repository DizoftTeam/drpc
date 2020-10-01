package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	nullFlagString = "<nil>"
	nullFlagInt    = -1

	jsonRpcVersion = "2.0"
)

func main() {
	// URL до сервера
	var url string

	// Имя метода JSONRpc
	var method string

	// Параметры метода
	var params interface{}

	// Идентификатор запроса
	var id int

	urlPtr := flag.String("url", nullFlagString, "URL to server")
	methodPrt := flag.String("method", nullFlagString, "JsonRPC method name")
	idPtr := flag.Int("id", nullFlagInt, "JsonRPC id (if not notify)")
	paramsPtr := flag.String("params", "{}", "JsonRPC params block")

	flag.Parse()

	if *urlPtr != nullFlagString {
		url = *urlPtr
	} else {
		panic("Url must be specified!")
	}

	if *methodPrt != nullFlagString {
		method = *methodPrt
	} else {
		panic("Method must be specified!")
	}

	if *idPtr != nullFlagInt {
		id = *idPtr
	}

	if err := json.Unmarshal([]byte(*paramsPtr), &params); err != nil {
		params = nil

		log.Println("Cant understand params", *paramsPtr)
	}

	jsonRpc := map[string]interface{}{}

	jsonRpc["jsonrpc"] = jsonRpcVersion
	jsonRpc["method"] = method
	jsonRpc["params"] = params

	if id > 0 {
		jsonRpc["id"] = id
	}

	request, err := json.Marshal(jsonRpc)

	if err != nil {
		fmt.Println(err)

		panic("Cant generate request json")
	}

	response, _ := http.Post(url, "application/json", strings.NewReader(string(request)))

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
}
