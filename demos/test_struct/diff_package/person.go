package diff_package

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) GetFirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(firstName string) {
	p.firstName = firstName
}
