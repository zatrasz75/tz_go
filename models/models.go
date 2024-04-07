package models

type Car struct {
	ID      int    `json:"ID"`
	RegNum  string `json:"regNum"`
	Mark    string `json:"mark"`
	Model   string `json:"model"`
	Year    int    `json:"year"`
	OwnerID int    `json:"ownerId"`
	Owner   People
}

type People struct {
	ID         int    `json:"ID"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
