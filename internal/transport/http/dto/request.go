package dto

type CreateShortURLRequest struct {
	URL        string  `json:"url" binding:"required,url"`
	CustomCode *string `json:"custom_code,omitempty" binding:"omitempty,alphanum,min=3,max=10"`
}

type RedirectRequest struct {
	CustomCode string `uri:"code" binding:"required,alphanum,min=3,max=10"`
}

type GetLinkInfoRequest struct {
	CustomCode string `uri:"code" binding:"required,alphanum,min=3,max=10"`
}
