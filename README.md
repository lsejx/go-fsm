# go-fsm
* Finite-state-machine.
* Generics.
* A map type TransitionMap can be used as a simple transition table.
* It allows custom transition table types which implement State interface.
<br><br>

# Import
	import "github.com/lsejx/go-fsm"

# Types
|name|kind|description|
|:---|:---|:----------|
|FSM|struct|contains initial state id, current state id, and StateMap.|
|StateMap|map|its key is state id, value is state data which implements State.|
|State|interface|wraps Next method, which receives a input as a parameter then returns next state id based on the input.|
|TransitionMap|map|implements State interface, its key is input, value is next state id.|

# Example
## Accepts "abc" using TransitionMap
	// stateIds
	stateA := 'a'
	stateB := 'b'
	stateC := 'c'
	initialState := stateA

	fsm := NewFSM(initialState, StateMap[rune, rune]{
		stateA: TransitionMap[rune, rune]{
			'a': stateA,
			'b': stateB,
			'c': stateA,
		},
		stateB: TransitionMap[rune, rune]{
			'a': stateA,
			'b': stateA,
			'c': stateC,
		},
		stateC: TransitionMap[rune, rune]{
			'a': stateC,
			'b': stateC,
			'c': stateC,
		},
	})

	fsm.CurrentStateId() // 'a'

	fsm.Input('a') // nil error
	fsm.CurrentStateId() // 'a'

	fsm.Input('b') // nil error
	fsm.CurrentStateId() // 'b'

	fsm.Input('z') // ErrUndefinedInput
	fsm.CurrentStateId() // 'b'

	fsm.Input('b') // nil error
	fsm.CurrentStateId() // 'a'

	fsm.Input('b') // nil error
	fsm.Input('c') // nil error
	fsm.CurrentStateId() // 'c'

	fsm.Reset()
	fsm.CurrentStateId() // 'a'



