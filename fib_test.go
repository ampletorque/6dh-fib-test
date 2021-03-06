package main

import (
	"testing"
//	"net/http"
//	"net/http/httptest"
//	"github.com/gorilla/mux"
)

//slice comparison function, as slices can't be compared with standard operators
func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

//test Fib - After failing to figure out how to compare the value of the enclosed
// function output, I then tried to resort to simply comparing the output of two
// identically called functions in order to at least test consistent output, but
// am still getting errors.

//func TestFib(t *testing.T) {
//	var result = Fib()
//	var comparison = Fib()
//	if result == nil {
//		t.Error("Expected 0, got ", result)
//	}
//}

//test IterateFib with custom slice comparison function. This test works.

func TestIterateFib(t *testing.T) {
	var digits int = 5	
	var comparison = []int{1, 1, 2, 3, 5}
	var result = IterateFib(digits)
	if !(Equal(result, comparison)) {
		t.Error("Expected", comparison, ", received", result)
	}
}

//test HandleCall -  This requires spinning up a test router and then sending it
// a sample API call and reading the results. I've been studying several approaches
// that require 3rd party libraries, and haven't finished this in time.

//func TestHandleCall(t *testing.T) {
//	router := mux.NewRouter().StrictSlash(true)
//        router.HandleFunc("/api/fibonacci/{digits}", HandleCall)
//	req := httptest.NewRequest("GET", "localhost:8080/api/fibonacci/5", nil)
//	w := httptest.NewRecorder()
//}

//new try at testing HandleCall, this one nearly works, but still not in time

//func executeRequest(req *http.Request) *httptest.ResponseRecorder {
//	router := mux.NewRouter()
//	rr := httptest.NewRecorder()
//	router.ServeHTTP(rr, req)
//
//	return rr
//}

//func checkResponseCode(t *testing.T, expected, actual int) {
//	if expected != actual {
//		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
//    }
//}

//func TestHandleCall(t *testing.T) {
//	var comparison = "[1,1,2,3,5]"
//	req, _ := http.NewRequest("GET", "/api/fibonacci/5", nil)
//	response := executeRequest(req)
//
//	checkResponseCode(t, http.StatusOK, response.Code)
//
//	if body := response.Body.String(); body != comparison {
//		t.Errorf("Error, did not receive expected output")
////		t.Errorf("Expected %s", comparison, ", got", body)
//	}
//}
