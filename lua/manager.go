package lua

import lua "github.com/yuin/gopher-lua"

var Variables map[string]int

func init() {
	L := lua.NewState()
	defer L.Close()

	Variables = make(map[string]int)

	L.SetGlobal("SetVariable", L.NewFunction(setVarible))
}

func setVarible(L *lua.LState) int {
	Variables[L.ToString(1)] = L.ToInt(2)
	return 0
}
