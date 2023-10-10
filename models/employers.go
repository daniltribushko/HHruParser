package models

type EmployerJson struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Url string `json:"alternate_url"`
	Logo struct{
		Url string `json:"original"`
	} `json:"logo_urls"`
}