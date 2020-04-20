package main

type Inventory struct {
    Material string
    Count    uint
}

func demo1()  {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }
}
func main()  {
	demo1()
}