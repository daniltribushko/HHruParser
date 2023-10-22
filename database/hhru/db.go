package hhru

import (
	"HHRUPARSER/models"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// Название субд
var dbType string = "postgres"

// Параметры подключения к базе данных
var connParam string = "user=postgres password=123456654321 dbname=hh_ru_go_db sslmode=disable"

/*
Получение региона из бд по id региона с hh.ru
*/
func GetAreaByAreaId(areaId string) models.AreaDb {
	stringError := "GetAreaByAreaId: "
	db, err := sql.Open(dbType, connParam)
	//Регион, который будет возвращаться при ошибке
	areaNil := models.AreaDb{Id: 0, AreaId: "", AreaName: ""}
	//Если при подключении к бд получили ошибки, то возвращаем пустой регион и ошибку
	if err != nil {
		log.Println(stringError, err)
		return areaNil
	}
	defer db.Close()
	//Запрос на получение региона
	row := db.QueryRow("select * from areas where areaid=$1", areaId)
	area := models.AreaDb{}
	err = row.Scan(&area.Id, &area.AreaId, &area.AreaName)
	if err != nil {
		log.Println(stringError, err)
	}

	return area
}

/*
Получение всех регионов
*/
func GetAllAreas() []models.AreaDb {
	stringError := "GetAllAreas: "
	db, err := sql.Open(dbType, connParam)
	if err != nil {
		log.Println(stringError, err)
		return nil
	}
	defer db.Close()
	result := []models.AreaDb{}
	//Отправляем запрос на получение регионов
	rows, err := db.Query("select * from areas")
	if err != nil {
		log.Println(stringError, err)
		return nil
	}
	defer rows.Close()
	//Добавляем регионы в срез пока есть данные, при ошибки пропускаем строку
	for rows.Next() {
		var area models.AreaDb
		err := rows.Scan(&area.Id, &area.AreaId, &area.AreaName)
		if err != nil {
			log.Println(stringError, err)
			continue
		}
		result = append(result, area)
	}
	return result
}

/*
Сохранение регионов в бд
*/
func InsertArea(value models.AreaJson) {
	stringError := "InsertArea: "
	db, err := sql.Open(dbType, connParam)
	if err != nil {
		log.Println(stringError, err)
		return
	}
	defer db.Close()
	//Сохраняем в бд
	_, err = db.Exec("insert into areas(areaid, areaname) values ($1, $2)", value.Id, value.Name)
	if err == nil {
		log.Println(value.Name, ": регион сохранен")
	} else {
		log.Println(stringError, err)
	}
}

/*
Получение работодателя из бд по id работодателю с hh.ru
*/
func GetEmployerByEmployerId(employerId string) models.EmployerDb {
	stringError := "GetEmployerByEmployerId: "
	db, err := sql.Open(dbType, connParam)
	//Работодатель, который будет возвращаться при ошибке
	employerNul := models.EmployerDb{Id: 0, EmployerId: "", Name: "", Url: "", LogoUrl: ""}
	//Если произошла ошибка при подключении к бд, то возвращается пустой работодатель и ошибка
	if err != nil {
		log.Println(stringError, err)
		return employerNul
	}
	defer db.Close()
	//Запрос на получение работодателя из бд
	row := db.QueryRow("select * from employers where employerid=$1", employerId)
	employer := models.EmployerDb{}
	err = row.Scan(&employer.Id, &employer.EmployerId, &employer.Name, &employer.Url, &employer.LogoUrl)
	if err != nil {
		log.Println(stringError, err)
	}
	return employer
}

/*
Получение всех работодателей из бд
*/
func GetAllEmployers() []models.EmployerDb {
	stringError := "GetAllEmployers: "
	db, err := sql.Open(dbType, connParam)
	if err != nil {
		log.Println(stringError, err)
		return nil
	}
	result := []models.EmployerDb{}
	defer db.Close()
	//Запрос на получение всех работодателей с бд
	rows, err := db.Query("select * from employers")
	if err != nil {
		log.Println(stringError, err)
		return nil
	}
	defer rows.Close()
	//Проходим по строкам данных
	for rows.Next() {
		//Если нет ошибки добавляем в срез работодателя, иначе идем дальше
		employer := models.EmployerDb{}
		err := rows.Scan(&employer.Id, &employer.EmployerId, &employer.Name, &employer.Url, &employer.LogoUrl)
		if err != nil {
			log.Println(stringError, err)
			continue
		}
		result = append(result, employer)
	}
	return result
}

/*
Сохранение работодателей в бд
*/
func InsertEmployer(value models.EmployerJson) {
	stringError := "InserEmployer: "
	db, err := sql.Open(dbType, connParam)
	if err != nil {
		log.Println(stringError, err)
		return
	}
	defer db.Close()
	//Запрос на добавление данных
	_, err = db.Exec("insert into employers (employerid, name, url, logo_url) values($1, $2, $3, $4)",
		value.Id, value.Name, value.Url, value.Logo.Url)
	if err == nil {
		log.Println(value.Name, ": работодатель сохранен")
	} else {
		log.Println(stringError, err)
	}
}

/*
Получение вакансии из бд по id с hh.ru
*/
func GetVacancyByVacancyId(vacancyId string) models.VacancyDb {
	stringError := "GetVacancyByVacancyId: "
	db, err := sql.Open(dbType, connParam)
	//Вакансия, которая будет возвращаться при ошибке
	vacancyNul := models.VacancyDb{Id: 0, VacancyId: "", Name: "", ProgrammingLanguage: "",
		AreaId: 0, SalaryFrom: 0, SalaryTo: 0, Currency: "", Experience: "", Schedule: "", Employment: "",
		Description: "", EmployerId: 0, DatePublish: time.Now(), Url: ""}
	if err != nil {
		log.Println(stringError, err)
		return vacancyNul
	}
	defer db.Close()
	//Запрос на получение вакансии
	row := db.QueryRow("select * from vacancies where vacancyid=$1", vacancyId)
	var vacancy models.VacancyDb
	err = row.Scan(&vacancy.Id, &vacancy.VacancyId, &vacancy.Name, &vacancy.ProgrammingLanguage, &vacancy.AreaId,
		&vacancy.SalaryFrom, &vacancy.SalaryTo, &vacancy.Currency, &vacancy.Experience, &vacancy.Schedule, &vacancy.Employment,
		&vacancy.Description, &vacancy.EmployerId, &vacancy.DatePublish, &vacancy.Url)
	if err != nil {
		log.Println(stringError, err)
	}
	return vacancy
}

/*
Получение всех вакансий с бд
*/
func GetAllVacancies() []models.VacancyDb {
	stringError := "GetAllVacancies: "
	db, err := sql.Open(dbType, connParam)
	if err != nil {
		log.Println(stringError, err)
		return nil
	}
	defer db.Close()
	result := []models.VacancyDb{}
	//Запрос на получение вакансий
	rows, err := db.Query("select * from vacancies")
	if err != nil {
		log.Println(stringError, err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		vacancy := models.VacancyDb{}
		err = rows.Scan(&vacancy.Id, &vacancy.VacancyId, &vacancy.Name, &vacancy.ProgrammingLanguage, &vacancy.AreaId,
			&vacancy.SalaryFrom, &vacancy.SalaryTo, &vacancy.Currency, &vacancy.Experience, &vacancy.Schedule,
			&vacancy.Employment, &vacancy.Description, &vacancy.EmployerId, &vacancy.DatePublish, &vacancy.Url)
		if err != nil {
			log.Println(stringError, err)
			continue
		}
		result = append(result, vacancy)
	}
	return result
}

/*
Сохранение навыхов в бд
*/
func insertSkills(vacancies []models.VacancyJson) {
	stringError := "insertSkills: "
	db, err := sql.Open(dbType, connParam)
	if err != nil {
		log.Println(stringError, err)
		return
	}
	defer db.Close()
	//Срез скилов
	var skills []models.VacancyParameterName
	query := strings.Builder{}
	//Формируем запрос
	query.WriteString("insert into skills (name, vacancyid)\nvalues\n")
	//Количество скилов
	var countSkills int
	//Количество вакансий
	countVacancies := len(vacancies)
	for indexV, value := range vacancies {
		//Получаем вакансию по id и скилы
		vacancy := GetVacancyByVacancyId(value.Id)
		skills = value.Skills
		if vacancy.Id == 0 {
			log.Println("Вакансия не найдена")
			continue
		}
		//Количество скиллов
		countSkills = len(value.Skills)
		for index, value := range skills {
			//Добавляем значения в запрос
			query.WriteString(fmt.Sprintf("('%s', %d)", value.Name, vacancy.Id))
			if index + 1 != countSkills || indexV + 1 != countVacancies {
				query.WriteString(",\n")
			}
		}
	}
	//Отправляем запрос
	_, err = db.Exec(query.String())
	if err != nil {
		log.Println(stringError, err)
	}
}

/*
Добавление вакансий в бд
*/
func InsertVacancies(pl models.ProgrammingLanguage, vacancies []models.VacancyJson) {
	stringError := "InsertVacancies: "
	db, err := sql.Open(dbType, connParam)
	if err != nil {
		log.Println(stringError, err)
		return
	}
	defer db.Close()
	//Формируем запрос
	query := strings.Builder{}
	query.WriteString("insert into vacancies (vacancyid, name, programminglanguage, area, salaryfrom, salaryto, currency, " +
		"experience, schedule, employment, description, employer, datepublish, url)\nvalues\n")
	count := len(vacancies)
	var area models.AreaJson
	var employer models.EmployerJson
	var areaDb models.AreaDb
	var employerDb models.EmployerDb
	//Проверка есть ли новые вакансии
	canAdd := false
	//Проходим по вакансиям
	for index, value := range vacancies {
		//Если вакансии нету в бд, то сохраняем
		if GetVacancyByVacancyId(value.Id).Id == 0 {
			canAdd = true
			area = value.Area
			employer = value.Employer
			areaDb = GetAreaByAreaId(area.Id)
			employerDb = GetEmployerByEmployerId(employer.Id)
			if areaDb.Id == 0 {
				InsertArea(area)
				areaDb = GetAreaByAreaId(area.Id)
			}
			if employerDb.Id == 0 {
				InsertEmployer(employer)
				employerDb = GetEmployerByEmployerId(employer.Id)
			}
			//Добавляем значения в бд
			query.WriteString(fmt.Sprintf("('%s','%s','%s','%d',%d,%d,'%s','%s','%s','%s','%s', %d, '%s', '%s')",
				value.Id, value.Name, pl, areaDb.Id, value.Salary.From, value.Salary.To, value.Salary.Currency,
				models.GetExperienceFromString(value.Experience.Id), models.GetScheduleFromString(value.Schedule.Id),
				models.GetEmploymentFromString(value.Employment.Id), value.Description,
				employerDb.Id, value.DatePublish, value.Url))
			if index+1 != count {
				query.WriteString(",")
			}
		}
	}
	if canAdd {
		_, err = db.Exec(query.String())
		if err != nil {
			log.Println(stringError, err)
		} else {
			insertSkills(vacancies)
		}
	} else {
		log.Println("Нет новых вакансий")
	}
}
