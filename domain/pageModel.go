package domain

type PageModel struct {
	Model	interface{}
	Menu	Menu
}

type Menu struct {
	CurrentLocation	string
	Sections 				interface{}
}