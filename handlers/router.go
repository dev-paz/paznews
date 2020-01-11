package handler

import "net/http"

//HandleRequests routes incoming requests
func HandleRequests() {
	http.HandleFunc("/top_headlines", handleGetTopArticles)

}
