package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func getCO2() string {
	out, err := exec.Command("sudo", "python3", "/usr/local/bin/get_co2.py").Output()
	if err != nil {
		log.Fatal(err)
	}
	CO2Value := string(out)
	return CO2Value
}

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		co2Value := getCO2()

		l1 := "# HELP co2_value CO2 Concentration"
		l2 := "# TYPE co2_value gauge"
		l3 := "co2_value " + co2Value

		msg := fmt.Sprintf("%v\n%v\n%v\n", l1, l2, l3)
		io.WriteString(w, msg)
	}
	http.HandleFunc("/metrics", h1)

	fmt.Println("server start.\nPort: 4102")
	log.Fatal(http.ListenAndServe(":4102", nil))
}
