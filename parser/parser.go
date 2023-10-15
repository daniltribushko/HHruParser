package parser

import (
	"HHRUPARSER/models"
	urlservice "HHRUPARSER/urlService"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var maxPerPage int = 100
var maxCountPage int = 20
var russiaId int = 113

/*
*
Получение id вакансий по России
*/
func getAllVacanciesIdFromRussia() map[models.ProgrammingLanguage][]models.VacancyJsonId {
	//Возвращаемый словарь
	var result map[models.ProgrammingLanguage][]models.VacancyJsonId = map[models.ProgrammingLanguage][]models.VacancyJsonId{}
	//Словарь параметров
	parameters := map[string]string{"per_page": strconv.Itoa(maxPerPage), "area": strconv.Itoa(russiaId)}
	//Сервис для работы с url адресом https://api.hh.ru/vacancies
	var url urlservice.UrlService = urlservice.UrlService{Url: "https://api.hh.ru/vacancies"}
	//Указатель на сервис
	var urlUk *urlservice.UrlService = &url
	//Список языков программирования
	var programmingLanguagesList []models.ProgrammingLanguage = models.ProgrammingLanguagesLit
	//Добавляем в url словарь параметров
	urlUk.AddParameters(parameters)
	//Проходим по языкам программирования
	for _, value := range programmingLanguagesList {
		//Список вакансий
		var listIds []models.VacancyJsonId
		//Добавляем в адрес язык программирования, как text
		urlUk.AddParameter("text", value.GetUrlName())
		//Проходим по страницам
		for i := 0; i < maxCountPage; i++ {
			//Добавляем в адрес номер страницы в качестве параметра
			urlUk.AddParameter("page", strconv.Itoa(i))
			//Делаем ограничения на 1 секунду
			time.Sleep(time.Second)
			body := sendGetRequest(url.Url)
			//Если ответ на запрос пустой то выходим из цикла с номерами страниц
			if body == nil {
				break
			}
			//Читаем ответ на запрос
			var response models.VacancyJsonIdResponse
			err := json.Unmarshal(body, &response)

			if err != nil {
				fmt.Println(err)
				break
			}
			//Получаем список id вакансий, если список пустой, то выходим из цикла страниц, иначе добавляем в список с id
			var ids []models.VacancyJsonId = response.Ids
			if len(ids) == 0 {
				break
			} else {
				listIds = append(listIds, ids...)
			}
		}
		//Добавляем в основной словарь значение язык программирования=список id вакансий
		result[value] = listIds
	}

	return result
}

/*
*
Получение вакансий по языкам программирования по всей России по id
*/
func GetAllVacanciesFromRussia() map[models.ProgrammingLanguage][]models.VacancyJson {
	//Финальный словарь
	var result map[models.ProgrammingLanguage][]models.VacancyJson = map[models.ProgrammingLanguage][]models.VacancyJson{}
	ids := getAllVacanciesIdFromRussia()
	var service urlservice.UrlService = urlservice.UrlService{Url: "https://api.hh.ru/vacancies"}
	var su *urlservice.UrlService = &service
	//Проходим по словарю языков программирования с id
	for pl, id := range ids {
		var vacancies []models.VacancyJson = []models.VacancyJson{}
		var length int = len(id)
		for index, value := range id {
			//Задержка 500 секунд
			time.Sleep(time.Second / 2)
			//Добавляем в url адрес id вакансии
			su.AddPathParameter(value.Id)
			//Получаем ответ
			body := sendGetRequest(service.Url)
			//Если ответ пустой, то удаляем id и выходим из цикла
			if body == nil {
				su.DeleteLastPathPart()
				break
			}
			//Удаляем id из адреса
			su.DeleteLastPathPart()
			//Преобразуем ответ на получение вакансии
			var vacanciesResponse models.VacancyJsonResponse
			err := json.Unmarshal(body, &vacanciesResponse)

			if err != nil {
				fmt.Println(err)
				return nil
			}
			//Добавляем вакансию в список вакансий
			vacancy := vacanciesResponse.Vacancies
			vacancies = append(vacancies, vacancy)

			log.Printf("%s - Вакансия получена %d/%d", string(pl), index + 1, length)
		}

		result[pl] = vacancies
	}

	return result
}

/*
*
Функция для отправки get запросов
*/
func sendGetRequest(url string) []byte {
	re, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	log.Println(url)
	log.Println(re.Request.Response.StatusCode)

	body, err := io.ReadAll(re.Body)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return body
}
