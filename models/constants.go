package models

type ProgrammingLanguage string
type Experience string
type Employment string
type Schedule string

var ProgrammingLanguagesLit []ProgrammingLanguage = []ProgrammingLanguage{Java, Scala, Go, Kotlin, CSharp,
	CPlus, PHP, Python, Lua,
	Rust, JavaScript, TypeScript}

const (
	Java       ProgrammingLanguage = "Java"
	Scala      ProgrammingLanguage = "Scala"
	Go         ProgrammingLanguage = "Go"
	Kotlin     ProgrammingLanguage = "Kotlin"
	CSharp     ProgrammingLanguage = "C#"
	CPlus      ProgrammingLanguage = "C++"
	PHP        ProgrammingLanguage = "PHP"
	Python     ProgrammingLanguage = "Python"
	Lua        ProgrammingLanguage = "Lua"
	Rust       ProgrammingLanguage = "Rust"
	JavaScript ProgrammingLanguage = "JavaScript"
	TypeScript ProgrammingLanguage = "TypeScript"
)

const (
	NoExperience Experience = "NoExperience"
	Between1And3 Experience = "Between1And3"
	Between3And6 Experience = "Between3And6"
	MoreThan6    Experience = "MoreThan6"
)

const (
	//Полная занятость
	Full Employment = "Full"
	//Частичная занятость
	Part Employment = "Part"
	//Проектная работа
	Project Employment = "Project"
	//Волонтерство
	Volunteer Employment = "Volunteer"
	//Стажировка
	Probation Employment = "Probation"
)

const (
	//Полный день
	FullDay Schedule = "FullDay"
	//Сменный график
	Shift Schedule = "Shift"
	//Гибкий график
	Flexible Schedule = "Flexible"
	//Удаленная работа
	Remote Schedule = "Remote"
	//Вахтовый метод
	FlyInFlyOut Schedule = "FlyInFlyOut"
)

/**
	Получение названия языка программирования в url кодировке
*/
func (p ProgrammingLanguage) GetUrlName() string {
	var result string
	switch p {
	case Go:
		result = "Golang"
	case CSharp:
		result = "C%23"
	case CPlus:
		result = "C%2B%2B"
	default:
		result = string(p)
	}

	return result
}

/**
	Преобразование опыта работы из строки в тип Experience
*/
func (e Experience) GetExperienceFromString(experience string) Experience {
	var result Experience
	switch experience {
	case "NoExperience":
		result = NoExperience
	case "Between1And3":
		result = Between1And3
	case "Between3And6":
		result = Between3And6
	case "MoreThan6":
		result = MoreThan6
	}

	return result
}

/**
	Преобразование типы работы из строки в тип Employment
*/
func (e Employment) GetEmploymentFromString(employment string) Employment {
	var result Employment
	switch employment {
	case "Full":
		result = Full
	case "Part":
		result = Part
	case "Probation":
		result = Probation
	case "Project":
		result = Project
	case "Volunteer":
		result = Volunteer
	}

	return result
}

/**
	Преобразование расписание работы из строки в тип Schedule
*/
func (s Schedule) GetScheduleFromString(schedule string) Schedule {
	var result Schedule
	switch schedule {
	case "FullDay":
		result = FullDay
	case "Shift":
		result = Shift
	case "Flexible":
		result = Flexible
	case "Remote":
		result = Remote
	case "FlyInFlyOut":
		result = FlyInFlyOut
	}
	return result
}
