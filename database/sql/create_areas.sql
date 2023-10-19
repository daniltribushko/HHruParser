--Таблица регионов
create table Areas(
	Id SERIAL PRIMARY KEY,
	--Id региона в hh.ru
	AreaId VARCHAR(255) NOT NULL,
	--Наименование регина в hh.ru
	AreaName VARCHAR(255) NOT NULL
)