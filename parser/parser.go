package parser

import (
	"HHRUPARSER/models"
	urlservice "HHRUPARSER/urlService"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

var maxPerPage int = 100
var maxCountPage int = 20
var russiaId int = 113

/**
	Получение id вакансий по России
*/
func getAllVacanciesIdFromRussia() []models.VacancyJsonId {
	var result []models.VacancyJsonId 
	parameters := map[string]string{"per_page": strconv.Itoa(maxPerPage), "area": strconv.Itoa(russiaId)}
	var url urlservice.UrlService = urlservice.UrlService{Url: "https://api.hh.ru/vacancies"}
	var urlUk *urlservice.UrlService = &url
	var programmingLanguagesList []models.ProgrammingLanguage  = models.ProgrammingLanguagesLit

	urlUk.AddParameters(parameters)
	for _, value := range programmingLanguagesList{
		urlUk.AddParameter("text", value.GetUrlName())
		for  i := 0; i < maxCountPage; i++ {
			urlUk.AddParameter("page", strconv.Itoa(i))
			time.Sleep(1000)
			body := sendGetRequest(url.Url)

			if body == nil {
				break
			}

			var response models.VacancyJsonIdResponse
			err := json.Unmarshal(body, &response)

			if err != nil {
				fmt.Println(err)
				break
			}
			
			var ids []models.VacancyJsonId = response.Ids

			if len(ids) == 0 {
				fmt.Println(ids)
				break
			} else {
				result = append(result, ids...)
			}
		}
	}

	return result
}

/**
	Получение вакансий по языкам программирования по всей России по id
*/
func GetAllVacanciesFromRussia() []models.VacancyJson {
	var result []models.VacancyJson = []models.VacancyJson{}
	ids := getAllVacanciesIdFromRussia()

	var service urlservice.UrlService = urlservice.UrlService{Url: "https://api.hh.ru/vacancies"}
	var su * urlservice.UrlService = &service
	var length int = len(ids)
	for index, id := range ids {
		su.AddPathParameter(id.Id)
		body := sendGetRequest(service.Url)

		if (body == nil){
			su.DeleteLastPathPart()
			break
		}

		su.DeleteLastPathPart()
		var vacanciesResponse models.VacancyJsonResponse
		err := json.Unmarshal(body, &vacanciesResponse)

		if (err != nil){
			fmt.Println(err)
			return nil
		}

		vacancy := vacanciesResponse.Vacancies
		result = append(result, vacancy)

		fmt.Printf("%d/%d Вакансия получена\n", index, length)
		time.Sleep(500)
	}

	return result
}

/**
	Функция для отправки get запросов
*/
func sendGetRequest(url string) []byte {
	re, err := http.Get(url)

	fmt.Println(url)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	body, err := io.ReadAll(re.Body)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return body
}