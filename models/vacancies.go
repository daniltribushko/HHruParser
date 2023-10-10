package models

type VacancyJsonIdResponse struct {
	Ids []VacancyJsonId `json:"items"`
}

type VacancyJsonResponse struct {
	Vacancies []VacancyJson `json:"items"`
}

type VacancyJsonId struct {
	Id string `json:"id"`
}
type VacancyParameterID struct{
	Id string `json:"id"`
}
type VacancyParameterName struct{
	Name string `json:"name"`
}

type VacancyJson struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Area AreaJson `json:"area"`
	Salary struct {
		From int64 `json:"from"`
		To int64 `json:"to"`
		Currency string `json:"currency"`
	} `json:"salary"`
	Experince VacancyParameterID `json:"experience"`
	Schedule VacancyParameterID `json:"schedule"`
	Employment VacancyParameterID `json:"employment"`
	Description string `json:"description"`
	Skills []VacancyParameterName `json:"key_skills"`
	Employer EmployerJson `json:"employer"`
	DatePublish string `json:"published_at"`
	Url string `json:"alternate_url"`

}