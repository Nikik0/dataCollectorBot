package statemachine

import (
	"github.com/Nikik0/dataCollectorBot/internal/model"
	"github.com/Nikik0/dataCollectorBot/internal/repository"
	"github.com/Nikik0/dataCollectorBot/internal/utils"
)

type State interface {
	PerformStateAction(u *model.User, msg *model.Message) error
}

type InitiateWorkflowState struct {
}

func (state *InitiateWorkflowState) PerformStateAction(u *model.User, msg *model.Message) error {
	return nil
}

type PersonalDataConfirmationState struct {
}

func (state *PersonalDataConfirmationState) PerformStateAction(u *model.User, msg *model.Message) error {
	accepted, err := utils.ValidateAcceptedTerms(msg.Text)
	if err != nil {
		return err
	}
	u.SetAcceptedTerms(accepted)
	return nil
}

type NameRequestState struct {
}

func (state *NameRequestState) PerformStateAction(u *model.User, msg *model.Message) error {
	name, err := utils.ValidateName(msg.Text)
	if err != nil {
		return err
	}
	u.SetName(name)
	return nil
}

type SurnameRequestState struct {
}

func (state *SurnameRequestState) PerformStateAction(u *model.User, msg *model.Message) error {
	surname, err := utils.ValidateSurname(msg.Text)
	if err != nil {
		return err
	}
	u.SetSurname(surname)
	return nil
}

type BirthDateRequestState struct {
}

func (state *BirthDateRequestState) PerformStateAction(u *model.User, msg *model.Message) error {
	birthdate, err := utils.ValidateSurname(msg.Text)
	if err != nil {
		return err
	}
	u.SetBirthdate(birthdate)
	return nil
}

type EmailRequestState struct {
}

func (state *EmailRequestState) PerformStateAction(u *model.User, msg *model.Message) error {
	email, err := utils.ValidateSurname(msg.Text)
	if err != nil {
		return err
	}
	u.SetEmail(email)
	return nil
}

type ConfirmationState struct {
}

func (state *ConfirmationState) PerformStateAction(u *model.User, msg *model.Message) error {
	//todo logic
	return nil
}

type FlowFinishedState struct {
}

func (state *FlowFinishedState) PerformStateAction(u *model.User, msg *model.Message) error {
	u.SetFinished(true)
	repository.SaveUser(u)
	return nil
}
