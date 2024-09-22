package statemachine

type StateMachine struct {
	currentState State
}

func (sm *StateMachine) getCurrentState() State {
	return sm.currentState
}

func (sm *StateMachine) hasNext() bool {
	switch sm.currentState.(type) {
	case *FlowFinishedState:
		return false
	default:
		return true
	}
}
