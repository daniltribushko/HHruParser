create table Vacancies(
	Id SERIAL PRIMARY KEY,
	VacancyId VARCHAR(255) NOT NULL,
	Name VARCHAR(255) NOT NULL,
	ProgrammingLanguage programming_language NOT NULL,
	Area INTEGER REFERENCES areas(Id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
	SalaryFrom INTEGER,
	SalaryTo INTEGER,
	Currency VARCHAR(10),
	Experience experience NOT NULL,
	Schedule schedule NOT NULL,
	Employment employment NOT NULL,
	Description VARCHAR(10000) NOT NULL,
	Employer INTEGER REFERENCES employers(Id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
	DatePublish TIMESTAMP NOT NULL,
	URL VARCHAR(500) NOT NULL
);