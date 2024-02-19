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
	stateIni := '0'
	stateA := 'a'
	stateB := 'b'
	stateC := 'c'

	fsm := NewFSM(stateIni, StateMap[rune, rune]{
		stateIni: TransitionMap[rune, rune]{
			'a': stateA,
			'b': stateIni,
			'c': stateIni,
		},
		stateA: TransitionMap[rune, rune]{
			'a': stateIni,
			'b': stateB,
			'c': stateIni,
		},
		stateB: TransitionMap[rune, rune]{
			'a': stateIni,
			'b': stateIni,
			'c': stateC,
		},
		stateC: TransitionMap[rune, rune]{
			'a': stateA,
			'b': stateIni,
			'c': stateIni,
		},
	})


	fsm.CurrentStateId() // '0'

	fsm.Input('b') // nil error
	fsm.CurrentStateId() // '0'

	fsm.Input('a') // nil error
	fsm.CurrentStateId() // 'a'

	fsm.Input('b') // nil error
	fsm.CurrentStateId() // 'b'

	fsm.Input('z') // ErrUndefinedInput
	fsm.CurrentStateId() // 'b'

	fsm.Input('b') // nil error
	fsm.CurrentStateId() // '0'

	fsm.Input('a') // nil error
	fsm.Input('b') // nil error
	fsm.Input('c') // nil error
	fsm.CurrentStateId() // 'c'

	fsm.Reset()
	fsm.CurrentStateId() // '0'



