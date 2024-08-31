package main

import (
	"fmt"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Listagem de notas")

	// tem como configurar cabeçalho ...Seção3-19
}
func noteView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Exibindo uma nota")

	// tem como pegar parâmetros da url ...Seção3-20
}
func noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		// rejeitar
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return

		// not use - security
		//w.Header().Set("Allow", http.MethodPost)
	}
	fmt.Fprint(w, "Critando nota")
}

func main() {
	fmt.Println("servidor porta 5000")
	mux := http.NewServeMux()

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":5000", mux)
}

// go run main.go
// go mod init
