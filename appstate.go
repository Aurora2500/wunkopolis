package main

type Appstate interface {
	UpdateState()
}

type AppstateStack struct {
	stack []Appstate
}

func NewAppState() AppstateStack {
	return AppstateStack{
		stack: make([]Appstate, 0, 4),
	}
}

func (appstate *AppstateStack) Update() {
	appstate.stack[len(appstate.stack)-1].UpdateState()
}

func (appstate *AppstateStack) IsEmpty() bool {
	return len(appstate.stack) == 0
}
