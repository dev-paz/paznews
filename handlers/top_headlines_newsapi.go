package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleGetTopArticles(w http.ResponseWriter, req *http.Request) {
	category := ""
	sources := ""
	query := ""

	c, _ := req.URL.Query()["category"]
	s, _ := req.URL.Query()["sources"]
	q, _ := req.URL.Query()["q"]

	if len(c) > 0 {
		category = c[0]
	}

	if len(s) > 0 {
		sources = s[0]
	}

	if len(q) > 0 {
		query = q[0]
	}

	articles, err := fetchArticles("top-headlines", category, sources, query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := json.Marshal(articles)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
