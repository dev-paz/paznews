package dto

//SaveArticlesRequest struct
type SaveArticlesRequest struct {
	Articles []*Article
}

//ArticlesResponse struct
type ArticlesResponse struct {
	Status       string
	TotalResults int64
	Articles     []*Article
}

//SourcesResponse struct
type SourcesResponse struct {
	Status  string
	Sources []*Source
}

//Article struct
type Article struct {
	ID          string
	Author      string
	Title       string
	Source      Source
	Description string
	URL         string
	URLToImage  string
	PublishedAt string
	Content     string
}

//Source struct
type Source struct {
	ID          string
	Name        string
	Description string
	URL         string
	Category    string
	Language    string
	Country     string
}
