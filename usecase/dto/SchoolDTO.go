package dto

type SchoolDTO struct {
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	ProviderNumber string `json:"provider_number"`
	Email          string `json:"email"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	ZipCode        string `json:"zip_code"`
	StateId        uint   `json:"state_id"`
	Active         bool   `json:"active"`
}

type SchoolResponseDTO struct {
	Id             uint   `json:"id"`
	Url            string `json:"url"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	ZipCode        string `json:"zip_code"`
	ProviderNumber string `json:"provider_number"`
	StateId        uint   `json:"state_id"`
	Active         bool   `json:"active"`
}
