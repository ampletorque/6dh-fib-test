package main

//in order to run this package, you will need to install the gorilla url router
//using this command: go get github.com/gorilla/mux

import (
	"fmt"
	"strconv"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

// Fib returns a function that returns fibonacci #s
func Fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// IterateFib calls Fib as many times as is specified in API call
func IterateFib(n int) []int {
	var f = Fib()
	var output []int
        for i := 0; i < n; i++ {
                output = append(output, f())
        }
	return output
}

// HandleCall handles the API call, using URL variable 
// to call IterateFib and return the sequence
func HandleCall(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	digits, _ := strconv.Atoi(vars["digits"])
	output := IterateFib(digits)
	outJ, _ := json.Marshal(output)

// I chose to cap this at 92, as the requirements PDF document specified
// that type int is preferred due to being more performant, and with type
// int the upper bound for values that don't fail on my system was 92

	if digits > 92 || digits < 1 {
	fmt.Fprintln(w, "[ \"Error: Integer between 1 and 92 required, you used", digits, "\" ]") 
	} else {
	fmt.Fprintln(w, string(outJ))
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/fibonacci/{digits}", HandleCall)
	log.Fatal(http.ListenAndServe(":8080", router))
}
