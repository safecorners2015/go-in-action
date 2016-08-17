package entities

type user struct {
	Name  string
	Email string
}

type admin struct {
	user
	Rights int
}
