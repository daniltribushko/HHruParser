package models

import (
	"fmt"
	"time"
)

type VacancyJsonIdResponse struct {
	Ids []VacancyJsonId `json:"items"`
}
type VacancyJsonId struct {
	Id string `json:"id"`
}
type VacancyParameterID struct {
	Id string `json:"id"`
}
type VacancyParameterName struct {
	Name string `json:"name"`
}

type VacancyJson struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Area   AreaJson `json:"area"`
	Salary struct {
		From     int64  `json:"from"`
		To       int64  `json:"to"`
		Currency string `json:"currency"`
	} `json:"salary"`
	Experience  VacancyParameterID     `json:"experience"`
	Schedule    VacancyParameterID     `json:"schedule"`
	Employment  VacancyParameterID     `json:"employment"`
	Description string                 `json:"description"`
	Skills      []VacancyParameterName `json:"key_skills"`
	Employer    EmployerJson           `json:"employer"`
	DatePublish string                 `json:"published_at"`
	Url         string                 `json:"alternate_url"`
}

type VacancyDb struct {
	Id                  int
	VacancyId           string
	Name                string
	ProgrammingLanguage ProgrammingLanguage
	AreaId              int
	SalaryFrom          int
	SalaryTo            int
	Currency            string
	Experience          string
	Schedule            string
	Employment          string
	Description         string
	EmployerId          int
	DatePublish         time.Time
	Url                 string
}

func (v VacancyJson) ToString() string {
	area := v.Area
	salary := v.Salary
	employer := v.Employer

	return fmt.Sprintf("Vacancy{Id=%s, Name=%s, Area{Id=%s, Name=%s}, Salary{From=%d, To=%d, Currency=%s}, "+
		"Experience{Id=%s}, Schedule{Id=%s}, Employment{Id=%s}, Description=%s, Skills=%s, Employer{Id=%s, Name=%s, Url=%s, "+
		"Logo{Id=%s}}, DatePublish=%s, Url=%s}", v.Id, v.Name, area.Id, area.Name, salary.From, salary.To, salary.Currency,
		v.Experience.Id, v.Schedule.Id, v.Employer.Id, v.Description, v.Skills, employer.Id, employer.Name, employer.Url,
		employer.Logo.Url, v.DatePublish, v.Url)
}
