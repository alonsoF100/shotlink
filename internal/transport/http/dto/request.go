package dto

type CreateShortURLRequest struct {
	URL       string `json:"url" binding:"required,url"`
	ShortCode string `json:"short_code" binding:"required,alphanum,min=3,max=10"`
}

type RedirectRequest struct {
	ShortCode string `uri:"short_code" binding:"required,alphanum,min=3,max=10"`
}
