package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

const (
	jsonRpcVersion = "2.0"
)

var (
	args []string // Массив аргументов программы
)

type config struct {
	Requests []struct {
		Name   string      `yaml:"name"`
		URL    string      `yaml:"url"`
		Method string      `yaml:"method"`
		Params interface{} `yaml:"params"`
		ID     int         `yaml:"id,omitempty"`
	} `yaml:"requests"`
}

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

// Запуск CLI версии
func runCli() {
	args = os.Args[1:]

	command := args[0]

	args = spliceByIndex(0, args)

	isVerbose := false

	findValueMap(func(v string) bool {
		if strings.Contains(v, "-v") || strings.Contains(v, "-verbose") {
			isVerbose = true

			return true
		}

		return false
	})

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

	if params == "" {
		params = "{}"
	}

	var headers string

	findValueMap(func(v string) bool {
		if strings.Contains(v, "-headers") {
			pieces := strings.Split(v, "=")

			if len(pieces) == 2 {
				headers = pieces[1]

				return true
			}
		}

		return false
	})

	if headers == "" {
		headers = "{}"
	}

	var pParams map[string]interface{}

	if err := json.Unmarshal([]byte(params), &pParams); err != nil {
		pParams = map[string]interface{}{}

		log.Println("Cant understand params", params, err)
	}

	var pHeaders map[string]interface{}

	if err := json.Unmarshal([]byte(headers), &pHeaders); err != nil {
		pHeaders = map[string]interface{}{}

		log.Println("Cant understand headers", headers, err)
	}

	switch command {
	case "request":
		jsonRpc := map[string]interface{}{
			"jsonrpc": jsonRpcVersion,
			"method":  method,
			"params":  pParams,
		}

		if id > 0 {
			jsonRpc["id"] = id
		}

		f := colorjson.NewFormatter()
		f.Indent = 2
		f.KeyColor = color.New(color.FgCyan)

		coloredRequest, err := f.Marshal(jsonRpc)
		coloredHeaders, err := f.Marshal(pHeaders)
		request, _ := json.Marshal(jsonRpc)

		if err != nil {
			println(err.Error())
			panic("Cant generate colored json")
		}

		println("--->")

		if isVerbose == true {
			println(string(coloredHeaders))
			println("---")
		}

		println(string(coloredRequest))

		httpClient := &http.Client{}
		hRequest, err := http.NewRequest("POST", url, strings.NewReader(string(request)))

		if err != nil {
			panic("Cant create HttpClient!")
		}

		hRequest.Header.Add("Content-Type", "application/json")

		for k, v := range pHeaders {
			hRequest.Header.Add(k, v.(string))
		}

		start := time.Now()
		response, _ := httpClient.Do(hRequest)
		end := time.Now()
		timeDiff := end.Sub(start)

		milliseconds := int(timeDiff / time.Millisecond)

		println("<---")

		if response == nil {
			log.Fatalln("Response is null. Server is down?")
		}

		defer response.Body.Close()

		if isVerbose == true {
			fmt.Printf("Time: %dms\n\n", milliseconds)

			for i, v := range response.Header {
				fmt.Printf("[%s] %s\n", i, strings.Join(v, ","))
			}

			println("---")
		}

		if strings.Contains(response.Header.Get("Content-Type"), "application/json") {
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

		break

	case "curl":
		println("Run curl command")
		println("Not supported yet")

		break

	default:
		log.Fatalln("Unknown command!\nTry request or curl!")
	}
}

// Запуск файловой версии
func runCmd(cConfig config) {
	args := os.Args[1:]
	command := args[0]

	if strings.Contains(command, "-") {
		log.Fatalln("Wrong command name")
	}

	var commands []string

	for _, v := range cConfig.Requests {
		commands = append(commands, v.Name)
	}
}

func main() {
	pConfigFile := flag.String("file", "<none>", "Set file config")

	flag.Parse()

	if *pConfigFile != "<none>" {
		data, err := ioutil.ReadFile(*pConfigFile)

		if err != nil {
			log.Fatalf("Config file %s not found", *pConfigFile)
		}

		cConfig := config{}

		if err := yaml.Unmarshal(data, &cConfig); err != nil {
			log.Fatalln("Yaml file has incorrect data. Error: ", err)
		}

		runCmd(cConfig)
	} else {
		runCli()
	}
}
