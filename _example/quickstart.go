package main

import "github.com/vxcontrol/golua/lua"

func adder(L *lua.State) int {
	a := L.ToInteger(1)
	b := L.ToInteger(2)
	L.PushInteger(int64(a + b))
	return 1
}

func main() {
	L := lua.NewState()
	defer L.Close()
	L.OpenLibs()

	L.GetGlobal("print")
	L.PushString("Hello World!")
	L.CheckType(2, lua.LUA_TSTRING)
	L.Call(1, 0)

	L.Register("adder", adder)
	L.DoString("print(adder(2, 2))")
}
