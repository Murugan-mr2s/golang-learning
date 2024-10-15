package model

type User struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Users *[]User

/*
func (u *Users) ToJSON(rw io.Writer) error {
	e := json.NewEncoder(rw)
	return e.Encode(u)
}

func (u *User) ToJSON(rw io.Writer) error {
	e := json.NewEncoder(rw)
	return e.Encode(u)
}

func (u *User) fromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}
*/
