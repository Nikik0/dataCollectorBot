package model

type User struct {
	tgId                  int64
	acceptedTerms         bool
	name                  string
	surname               string
	birthdate             string
	email                 string
	finished              bool
	mistakeCorrectionNeed bool
}

func NewUser(tgId int64) *User {
	return &User{
		tgId:                  tgId,
		acceptedTerms:         false,
		finished:              false,
		mistakeCorrectionNeed: false,
	}
}

func (u *User) TgId() int64 {
	return u.tgId
}

func (u *User) SetAcceptedTerms(v bool) {
	u.acceptedTerms = v
}

func (u *User) SetName(n string) {
	u.name = n
}

func (u *User) SetSurname(sn string) {
	u.surname = sn
}

func (u *User) SetBirthdate(bd string) {
	u.birthdate = bd
}

func (u *User) SetEmail(e string) {
	u.email = e
}

func (u *User) SetFinished(f bool) {
	u.finished = f
}

func (u *User) IsMistakeCorrectionNeeded() bool {
	return u.mistakeCorrectionNeed
}

func (u *User) SetMistakeCorrectionNeeded(n bool) {
	u.mistakeCorrectionNeed = n
}
