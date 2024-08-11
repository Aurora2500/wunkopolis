package appstate

type Appstate interface {
	UpdateState() AppstateChange
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
	change := appstate.stack[len(appstate.stack)-1].UpdateState()
	switch change.kind {
	case noop:
	case push:
		appstate.stack = append(appstate.stack, change.new)
	case pop:
		appstate.stack = appstate.stack[:len(appstate.stack)-1]
	case swap:
		appstate.stack[len(appstate.stack)-1] = change.new
	}
}

func (appstate *AppstateStack) IsEmpty() bool {
	return len(appstate.stack) == 0
}

type ChangeType int

const (
	noop ChangeType = iota
	push
	pop
	swap
)

type AppstateChange struct {
	kind ChangeType
	new  Appstate
}
