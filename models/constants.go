package models

type ProgrammingLanguage string

var ProgrammingLanguagesLit []ProgrammingLanguage = []ProgrammingLanguage{Java, Scala, Go, Kotlin, CSharp, 
	CPlus, PHP, Python, Lua, 
	Rust, JavaScript, TypeScript}

const (
	Java ProgrammingLanguage = "Java"
	Scala ProgrammingLanguage = "Scala"
	Go ProgrammingLanguage = "Go"
	Kotlin ProgrammingLanguage = "Kotlin"
	CSharp ProgrammingLanguage = "C#"
	CPlus ProgrammingLanguage = "C++"
	PHP ProgrammingLanguage = "PHP"
	Python ProgrammingLanguage = "Python"
	Lua ProgrammingLanguage = "Lua"
	Rust ProgrammingLanguage = "Rust"
	JavaScript ProgrammingLanguage = "JavaScript"
	TypeScript ProgrammingLanguage = "TypeScript"
)

func (p ProgrammingLanguage) GetUrlName() string {
	var result string
	switch (p){
		case Go : result = "Golang"
		case CSharp : result = "C%23"
		case CPlus : result = "C%2B%2B"
		default : result = string(p)
	}

	return result
}