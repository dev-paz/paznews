package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/paznews/paznews-backend/dto"
)

var apiKey = "66cb91b07c064b8b8cce78a0f0c9c8f3"

func fetchArticles(
	endpoint string,
	category string,
	sources string,
	query string,
) (articlesResp dto.ArticlesResponse, err error) {

	queryString := fmt.Sprintf("https://newsapi.org/v2/%s?country=gb&category=%s&sources=%s&q=%s&apiKey=%s",
		endpoint,
		category,
		sources,
		query,
		apiKey,
	)

	rsp, err := http.Get(queryString)
	if err != nil {
		fmt.Println("error fetching")
		return dto.ArticlesResponse{}, err
	}

	articlesResponse := dto.ArticlesResponse{}
	err = json.NewDecoder(rsp.Body).Decode(&articlesResponse)
	if err != nil {
		fmt.Println("error decoding")
		fmt.Println(err.Error())
		return articlesResponse, err
	}

	return articlesResponse, nil
}
