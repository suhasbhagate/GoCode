package main

import(
	"net/http"
	//"github.com/gorilla/mux"
)

func main() {
	mux := http.NewServeMux() mux.Handle("/", index)
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
}

func index(w http.ResponseWriter, r *http.Request){ 
		fmt.Fprintf(w, "Welcome to Go Web") 
	}