package models

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/paznews/paznews-backend/dto"
)

// func ReadTopHeadlines() (*[]dto.Donation, error) {
// 	var d dto.Donation
// 	var donations []dto.Donation
// 	rows, err := db.Query(`SELECT donation_id, user_name, user_id, amount, donation_created_timestamp FROM donations`)
// 	if err != nil {
// 		fmt.Println("error querying")
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&d.ID, &d.UserName, &d.UserId, &d.Amount, &d.DonationCreatedTimestamp)
// 		if err != nil {
// 			// handle this error
// 			panic(err)
// 		}
// 		donations = append(donations, d)
// 	}
// 	return &donations, nil
// }

func SaveArticles(articles []*dto.Article) error {
	sqlStatement := `
	INSERT INTO top_headlines (id, author, title, description, url, url_to_image, published_at, content)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id`
	idstr := "123"
	for _, a := range articles {
		a.ID = idstr
		fmt.Println(a.ID)
		id := 0
		err := db.QueryRow(sqlStatement, a.ID, a.Author, a.Title, a.Description, a.URL, a.URLToImage, a.PublishedAt, a.Content).Scan(&id)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		idstr = idstr + "1"
	}

	return nil
}
