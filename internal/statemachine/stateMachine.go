package statemachine

import "github.com/Nikik0/dataCollectorBot/internal/model"

type StateMachine struct {
	currentState State
	transitions  map[State]map[Event]State
}

const (
	AcceptTermsNeeded = Event(rune(iota))
	NameNeeded
	SurnameNeeded
	BirthdateNeeded
	EmailNeeded
	ConfirmationNeeded
	FinishNeeded
	Default
)

func (sm *StateMachine) GetCurrentState() State {
	return sm.currentState
}

func (sm *StateMachine) SetCurrentState(s State) {
	sm.currentState = s
}

func (sm *StateMachine) hasNext() bool {
	switch sm.currentState.(type) {
	case *FlowFinishedState:
		return false
	default:
		return true
	}
}

func NewStateMachine() *StateMachine {
	sm := &StateMachine{
		currentState: nil,
		transitions:  make(map[State]map[Event]State),
	}
	setStatesForSM(sm)
	return sm
}

func (sm *StateMachine) SendEvent(event Event, u *model.User, msg *model.Message) {
	st := sm.transitions[sm.currentState][event]
	err := st.PerformStateAction(u, msg)
	if err != nil {
		return
	}
}

func setStatesForSM(sm *StateMachine) {
	sm.transitions[&InitiateWorkflowState{}] = map[Event]State{
		Default: &PersonalDataConfirmationState{},
	}

	sm.transitions[&PersonalDataConfirmationState{}] = map[Event]State{
		Default:            &NameRequestState{},
		ConfirmationNeeded: &ConfirmationState{},
	}

	sm.transitions[&NameRequestState{}] = map[Event]State{
		Default:            &SurnameRequestState{},
		ConfirmationNeeded: &ConfirmationState{},
	}

	sm.transitions[&SurnameRequestState{}] = map[Event]State{
		Default:            &BirthDateRequestState{},
		ConfirmationNeeded: &ConfirmationState{},
	}

	sm.transitions[&BirthDateRequestState{}] = map[Event]State{
		Default:            &EmailRequestState{},
		ConfirmationNeeded: &ConfirmationState{},
	}

	sm.transitions[&EmailRequestState{}] = map[Event]State{
		Default:            &ConfirmationState{},
		ConfirmationNeeded: &ConfirmationState{},
	}

	sm.transitions[&ConfirmationState{}] = map[Event]State{
		NameNeeded:      &NameRequestState{},
		SurnameNeeded:   &SurnameRequestState{},
		BirthdateNeeded: &BirthDateRequestState{},
		EmailNeeded:     &EmailRequestState{},
		FinishNeeded:    &FlowFinishedState{},
	}
}
