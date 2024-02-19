package fsm

import "errors"

var (
	// ErrUndefinedInput is returned from TransitionMap when the input is not defined in the TransitionMap.
	ErrUndefinedInput error = errors.New("undefined input")
)

// State is the interface that wraps only Next method, which is used to get next state after transition with the input.
type State[Input any, StateId comparable] interface {
	Next(Input) (StateId, error)
}

// TransitionMap is a map type which implements State interface.
// Key of the map is a input.
// Value of the map is a state id after the transition.
// This can be used as a simple transition table which only uses map.
type TransitionMap[Input, StateId comparable] map[Input]StateId

// Next returns state id associated with the input.
// ErrUndefinedInput is returned when the input is not defined in the TransitionMap.
func (tm TransitionMap[Input, StateId]) Next(i Input) (stateId StateId, err error) {
	id, ok := tm[i]
	if !ok {
		var tmp StateId
		return tmp, ErrUndefinedInput
	}
	return id, nil
}

// StateMap is a map type which associates state id and state.
type StateMap[Input any, StateId comparable] map[StateId]State[Input, StateId]

// FSM is initialized by function NewFSM.
type FSM[Input any, StateId comparable] struct {
	initial StateId
	current StateId
	states  StateMap[Input, StateId]
}

// Input inputs a value into FSM.
// Non nil error is returned when State.Next returned non nil error.
func (fsm *FSM[Input, StateId]) Input(input Input) error {
	id, err := fsm.states[fsm.current].Next(input)
	if err != nil {
		return err
	}
	fsm.current = id
	return nil
}

// CurrentStateId returns state id of current state.
func (fsm *FSM[Input, StateId]) CurrentStateId() (stateId StateId) {
	return fsm.current
}

// Reset resets fsm to initial state.
func (fsm *FSM[Input, StateId]) Reset() {
	fsm.current = fsm.initial
}

// NewFSM returns new struct FSM with initial state, and all states.
// Data of All states are is constructed as map type StateMap.
func NewFSM[Input any, StateId comparable](initialStateId StateId, states StateMap[Input, StateId]) FSM[Input, StateId] {
	return FSM[Input, StateId]{
		initial: initialStateId,
		current: initialStateId,
		states:  states,
	}
}
