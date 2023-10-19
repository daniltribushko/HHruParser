--Таблица работодателей
create table Employers(
	Id SERIAL PRIMARY KEY,
	--Id работодателя с hh.ru
	EmployerId VARCHAR(255) NOT NULL,
	--Наименование работодателя
	Name VARCHAR(255) NOT NULL,
	--Url адрес работодателя
	Url VARCHAR(255),
	--Url логотипа работодателя
	Logo_Url VARCHAR(255)
);