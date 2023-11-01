package models

type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	License bool   `json:"license"`
}

func (p *Person) ValidationPerson() bool {
	if p.Age >= 18 && p.Age <= 40 {
		return true
	}
	if p.Age > 40 && p.Age <= 60 && p.License == true {
		return true
	}
	return false
}
