package domain

type EncoreTab struct {
	Id  uint64 `json:"-" db:"id"`
	Url string `json:"url"`
}
