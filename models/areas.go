package models

type AreaJson struct{
	Id string `json:"id"`
	Name string `json:"name"`
}

type AreaDb struct {
	Id int
	AreaId string
	AreaName string
}
