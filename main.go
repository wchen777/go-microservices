package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// func testPath() {
// 	gopath := os.Getenv("GOPATH")
// 	if gopath == "" {
// 		log.Fatal("Your GOPATH has not been set!")
// 	}

// 	path := os.Getenv("PATH")
// 	gobin := filepath.Join(gopath, "bin")
// 	if !strings.Contains(path, gobin) {
// 		log.Fatalf("Your PATH does not contain %s", gobin)
// 	}

// 	log.Println("Success!")
// }

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("hello")
		d, err := ioutil.ReadAll((r.Body))

		// err check w/ response code
		if err != nil {
			http.Error(rw, "error caught", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("error caught"))
			return
		}

		log.Printf("Data %s\n", d)
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye")
	})

	// start server
	http.ListenAndServe(":9090", nil)

}
