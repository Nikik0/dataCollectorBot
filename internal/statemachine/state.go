package statemachine

import "github.com/Nikik0/dataCollectorBot/internal/model"

type State interface {
	performStateAction(u *model.User)
}

type PersonalDataConfirmationState struct {
}

func (state *PersonalDataConfirmationState) performStateAction(u *model.User) {

}

type NameRequestState struct {
}

func (state *NameRequestState) performStateAction(u *model.User) {

}

type SurnameRequestState struct {
}

func (state *SurnameRequestState) performStateAction(u *model.User) {

}

type BirthDateRequestState struct {
}

func (state *BirthDateRequestState) performStateAction(u *model.User) {

}

type EmailRequestState struct {
}

func (state *EmailRequestState) performStateAction(u *model.User) {

}

type ConfirmationState struct {
}

func (state *ConfirmationState) performStateAction(u *model.User) {

}

type FlowFinishedState struct {
}

func (state *FlowFinishedState) performStateAction(u *model.User) {

}
