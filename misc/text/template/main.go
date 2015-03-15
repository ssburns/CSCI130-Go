package main

import(
	"log"
	"os"
	"text/template"
)

type Recipient struct{
	Honorific, Name string
	Donated bool
}

type UpcomingEvents struct{
	Events []string
}

func main() {
	const letter =`Dear {{.Person.Honorific}} {{.Person.Name}}
{{if .Person.Donated}}
Thank you for donating!
{{else}}
It makes us sad that you didn't donate, something something puppies and kittens. :(
{{end}}
`
	const letter2 = "{{.Extra}}"

	recipients := []Recipient{
		{"Mr", "Burns", false},
		{"Mr", "Smithers", true},
		{"Mrs", "Simpson", true},
		{"Mrs", "Krabappel", false},
	}

	upcomingEvents := UpcomingEvents{Events:[]string{
		"event1",
		"event2",
		"event3"}}


	t := template.Must(template.New("letter").Parse(letter))

	for _,r := range recipients{
		data := struct{
			Person *Recipient
			MoreInfo UpcomingEvents
			}{
			&r,
			upcomingEvents,
			}
		err := t.Execute(os.Stdout, data)
		if err != nil{
			log.Println("executing template:", err)
		}
	}
}
