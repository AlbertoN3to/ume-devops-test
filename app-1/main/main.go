package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App2Domain string `yaml:"app-2-domain"`
}

var cli http.Client
var config Config

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello!!!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
			fmt.Printf("%v: %v\n", name, h)
		}
	}
}

func ping(w http.ResponseWriter, req *http.Request) {
	msg := req.URL.Query().Get("msg")

	url := fmt.Sprintf("http://%s/api/v1/pong", config.App2Domain)
	pongResp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching response from Pong service", http.StatusInternalServerError)
		return
	}
	defer pongResp.Body.Close()

	pongBody, err := ioutil.ReadAll(pongResp.Body)
	if err != nil {
		http.Error(w, "Error reading response from Pong service", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Received message: %s, Pong response: %s\n", msg, pongBody)
}

func main() {
	var configFile string
	if len(os.Args) == 2 {
		configFile = os.Args[1]
	} else {
		configFile = "/mnt/config/config.yaml"
	}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("Error reading YAML file: %v", err))
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(fmt.Errorf("Error unmarshaling YAML data: %v", err))
	}

	cli = http.Client{Timeout: time.Duration(10) * time.Second}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/api/v1/ping", ping)

	fmt.Printf("running server on port 8080\n")
	http.ListenAndServe(":8080", nil)
}
