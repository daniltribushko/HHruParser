package urlservice

import (
	"fmt"
	"regexp"
	"strings"
)

type UrlService struct {
	Url string
}

/**
	Добавление параметра в url адрес
*/
func (u *UrlService) AddParameter(key string, value string){
	var newUrl strings.Builder
	oldUrl := u.Url
	newUrl.WriteString(oldUrl)

	if !strings.Contains(oldUrl, "?") {
		newUrl.WriteString("?")
	}
	
	//Если ключ параметра уже указан, то меняем значение
	if strings.Contains(oldUrl, key) {
		if strings.EqualFold(key, "page") {
			if strings.Count(oldUrl, "page") == 2{
				oldUrl = newUrl.String()
				newUrl.Reset()
				newUrl.WriteString(replaceValueUrl(key, value, oldUrl))
			} else {
				newUrl.WriteString("&")
				newUrl.WriteString(key)
				newUrl.WriteString("=")
				newUrl.WriteString(value)
			}
		} else {
			oldUrl = newUrl.String()
			newUrl.Reset()
			newUrl.WriteString(replaceValueUrl(key, value, oldUrl))
		}
	}else{
		//Проверка указаны ли параметры в url
		if strings.Contains(oldUrl, "?") == true && strings.Contains(oldUrl, "=") {
			newUrl.WriteString("&")
		}
		newUrl.WriteString(key)
		newUrl.WriteString("=")
		newUrl.WriteString(value)
	}

	(*u).Url = newUrl.String()
	
}

/**
	Добавление словаря параметров
*/
func (u *UrlService) AddParameters(parameters map [string]string){
	var newUrl strings.Builder
	newUrl.WriteString(u.Url)

	if !strings.Contains(u.Url, "?") {
		newUrl.WriteString("?")
	}else if strings.Contains(u.Url, "=") {
		newUrl.WriteString("&")
	}

	index := 0
	for key, value := range parameters {
		if !strings.Contains(newUrl.String(), key){
			newUrl.WriteString(key)
			newUrl.WriteString("=")
			newUrl.WriteString(value)

			if len(parameters) != index + 1 {
				newUrl.WriteString("&")
			}
		} else {
			oldUrl := newUrl.String()
			newUrl.Reset()
			newUrl.WriteString(replaceValueUrl(key, value, oldUrl))
		}
		index++

		(*u).Url = newUrl.String()
	}
}

/**
	Добавление параметра в путь url
*/
func (u *UrlService) AddPathParameter(value string){
	oldUrl := u.Url
	var newUrl strings.Builder
	newUrl.WriteString(oldUrl)
	if !strings.EqualFold(u.Url[:len(u.Url)-1], "/") {
		newUrl.WriteString("/")
	}

	newUrl.WriteString(value)
	(*u).Url = newUrl.String()
}

/**
	Удаление последней части путья url
*/
func (u *UrlService) DeleteLastPathPart() {
	var mas []string = strings.Split(u.Url, "/")
	var newUrl strings.Builder = strings.Builder{}
	for index, value := range mas[:len(mas)-1] {
		newUrl.WriteString(value)
		if (index != len(mas)-2){
			newUrl.WriteString("/")
		}
	}
	(*u).Url = newUrl.String()
}

/**
	Замена значения параметра в url адресе
*/
func replaceValueUrl(key string, value string, url string) string {
	var result string

	re, err := regexp.Compile(fmt.Sprintf(`&%s=.*&|\?%s=.*&|&%s=.*|\?%s=.*`, key, key, key, key))
	if err != nil {
		panic(err)
	}
	matched := re.MatchString(url)

	if matched {
		findString := re.FindString(url)
		oldVaue := strings.Split(findString, "=")[1]
		if strings.Contains(oldVaue, "&") {
			oldVaue = oldVaue[:len(oldVaue)-1]
		}
		newFindString := strings.Replace(findString, oldVaue, value, -1)
		result = strings.Replace(url, findString, newFindString, 1)
	}
	return result
}