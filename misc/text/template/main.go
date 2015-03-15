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

func main() {
	const letter =`Dear {{.Honorific}} {{.Name}}
{{if .Donated}}
Thank you for donating!
{{else}}
It makes us sad that you didn't donate, something something puppies and kittens. :(
{{end}}
`

	recipients := []Recipient{
		{"Mr", "Burns", false},
		{"Mr", "Smithers", true},
		{"Mrs", "Simpson", true},
	}

	t := template.Must(template.New("letter").Parse(letter))

	for _,r := range recipients{
		err := t.Execute(os.Stdout, r)
		if err != nil{
			log.Println("executing template:", err)
		}
	}
}
