package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//go:embed "input.txt"
var input string

var (
	tmpl = template.Must(template.ParseFiles("index.html"))
)

func main() {
	var logs []string
	input = strings.TrimSuffix(input, ",")
	lines := strings.Split(input, ",")

	logs = append(logs, trimToSec(lines, logs)...)
	fmt.Println(logs)

	data := []byte(strings.Join(logs, "\n"))

	err := ioutil.WriteFile("data.txt", data, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}

func trimToSec(num []string, log []string) []string {
	for _, n := range num {
		pin, _ := strconv.Atoi(n)
		if pin <= 60 {
			log = append(log, "A"+strconv.Itoa(pin+1))
		} else if pin > 60 && pin <= 120 {
			log = append(log, "B"+strconv.Itoa(pin-60+1))
		} else if pin > 120 && pin <= 180 {
			log = append(log, "C"+strconv.Itoa(pin-120+1))
		} else if pin > 180 && pin <= 240 {
			log = append(log, "D"+strconv.Itoa(pin-180+1))
		}
	}
	return log
}
