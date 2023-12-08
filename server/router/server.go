package router

import (
	"fmt"
	"net/http"
)

type Router struct {
	S    *http.ServeMux
	port int
}

func NewRouter() *Router {
	mux := http.NewServeMux()
	var port int
	fmt.Scanf("%d", &port)

	mux.HandleFunc("/AddSubscription", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// user := &Person{}
		// err := json.NewDecoder(r.Body).Decode(user)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }

		// fmt.Println("got user:", user)
		// w.WriteHeader(http.StatusCreated)
	})

	return &Router{S: mux, port: port}
}

func (r *Router) RunServer() {
	http.ListenAndServe(fmt.Sprintf(":%d", r.port), r.S)
}
