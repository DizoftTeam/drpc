package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
)

const (
	jsonRpcVersion = "2.0"
)

var (
	args []string // Массив аргументов программы
)

// Вырезает элемент из массива
func spliceByIndex(index int, arr []string) []string {
	copy(arr[index:], arr[index+1:])
	arr[len(arr)-1] = ""

	return arr[:len(arr)-1]
}

// Поиск элемента в массиве строк через callback
// Если функция возвращает false - значит поиск завершен
func findValueMap(callback func(v string) bool) {
	for k, v := range args {
		if callback(v) {
			args = spliceByIndex(k, args)

			break
		}
	}
}

func main() {
	args = os.Args[1:]

	command := args[0]

	args = spliceByIndex(0, args)

	switch command {
	case "request":
		println("Run request command")

		os.Exit(0)

		break

	case "curl":
		println("Run curl command")

		os.Exit(0)

		break

	default:
		println("Unkkown command!\nTry request or curl!")

		os.Exit(0)

		break
	}

	var url string

	findValueMap(func(v string) bool {
		if strings.Contains(v, "http") {
			url = v

			return true
		}

		return false
	})

	var id int

	findValueMap(func(v string) bool {
		if intVal, err := strconv.Atoi(v); err == nil {
			id = intVal

			return true
		}

		return false
	})

	var method string

	findValueMap(func(v string) bool {
		if strings.Contains(v, "-method") {
			pieces := strings.Split(v, "=")

			if len(pieces) == 2 {
				method = pieces[1]

				return true
			}
		}

		return false
	})

	var params string

	findValueMap(func(v string) bool {
		if strings.Contains(v, "-params") {
			pieces := strings.Split(v, "=")

			if len(pieces) == 2 {
				params = pieces[1]

				return true
			}
		}

		return false
	})

	var headers string

	findValueMap(func(v string) bool {
		if strings.Contains(v, "-headers") {
			pieces := strings.Split(v, "=")

			if len(pieces) == 2 {

			}
		}

		return false
	})

	println("headers", headers)

	var pParams interface{}

	if err := json.Unmarshal([]byte(params), &pParams); err != nil {
		pParams = map[string]string{}

		log.Println("Cant understand params", params, err)
	}

	jsonRpc := map[string]interface{}{}

	jsonRpc["jsonrpc"] = jsonRpcVersion
	jsonRpc["method"] = method
	jsonRpc["params"] = pParams

	if id > 0 {
		jsonRpc["id"] = id
	}

	f := colorjson.NewFormatter()
	f.Indent = 2
	f.KeyColor = color.New(color.FgCyan)

	fRequest, err := f.Marshal(jsonRpc)
	request, _ := json.Marshal(jsonRpc)

	if err != nil {
		println(err.Error())
		panic("Cant generate colored json")
	}

	println("--->")
	println(string(fRequest))

	response, _ := http.Post(url, "application/json", strings.NewReader(string(request)))

	defer response.Body.Close()

	println("<---")

	if response.Header.Get("Content-Type") == "application/json" {
		var s interface{}

		if err = json.NewDecoder(response.Body).Decode(&s); err != nil {
			println("JsonResponse could not be parsed", err.Error())
		}

		fr, _ := f.Marshal(s)

		println(string(fr))
	} else {
		body, _ := ioutil.ReadAll(response.Body)

		println(string(body))
	}
}
