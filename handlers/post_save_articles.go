package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/paznews/paznews-backend/dto"
	"github.com/paznews/paznews-backend/models"
)

func handlePOSTSaveArticles(w http.ResponseWriter, req *http.Request) {

	saveArticlesReq := dto.SaveArticlesRequest{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&saveArticlesReq)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = models.SaveArticles(saveArticlesReq.Articles)

}
