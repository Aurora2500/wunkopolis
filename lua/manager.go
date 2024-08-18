package lua

import (
	"errors"
	"wunkopolis/assets"
	"wunkopolis/ui"

	lua "github.com/yuin/gopher-lua"
)

var Variables map[string]int

var L *lua.LState

func init() {
	L = lua.NewState()
	defer L.Close()

	Variables = make(map[string]int)

	L.SetGlobal("SetVariable", L.NewFunction(setVaribleLua))

	L.DoFile("lua/main.lua")

}

func setVaribleLua(L *lua.LState) int {
	Variables[L.ToString(1)] = L.ToInt(2)
	return 0
}

func toWindow(table *lua.LTable, name string) (ui.Window, error) {
	window := ui.Window{}

	title := ""
	if str, ok := table.RawGetString("title").(lua.LString); ok {
		title = string(str)

	} else {
		return window, errors.New("Could not find title for " + name)
	}

	width := float32(0)
	if num, ok := table.RawGetString("width").(lua.LNumber); ok {
		width = float32(num)
	} else {
		return window, errors.New("Could not find width for " + name)
	}
	height := float32(0)
	if num, ok := table.RawGetString("height").(lua.LNumber); ok {
		height = float32(num)
	} else {
		return window, errors.New("Could not find height for " + name)
	}

	iconName := ""
	if str, ok := table.RawGetString("icon").(lua.LString); ok {
		iconName = string(str)
	} else {
		return window, errors.New("Could not find icon for " + name)
	}

	var content ui.UIElem = &ui.Flexbox{}
	if tbl, ok := table.RawGetString("content").(*lua.LTable); ok {
		if uielem, err := toUIElem(tbl); err == nil {
			content = uielem
		}
	}

	return ui.Window{
		Title: title,
		Area: ui.Area{
			X:      0,
			Y:      0,
			Width:  width,
			Height: height,
		},
		Content: content,
		Icon:    assets.Manager.GetTexture(iconName),
	}, nil

}

func toUIElem(table *lua.LTable) (ui.UIElem, error) {
	elementType := ""
	if str, ok := table.RawGetString("type").(lua.LString); ok {
		elementType = string(str)
	}

	switch elementType {
	case "Flexbox":
		{
			direction := 0
			if dir, ok := table.RawGetString("direction").(lua.LNumber); ok {
				direction = int(dir)
			}
			anchor := 0
			if anc, ok := table.RawGetString("anchor").(lua.LNumber); ok {
				anchor = int(anc)
			}
			padding := float32(0)
			if pad, ok := table.RawGetString("padding").(lua.LNumber); ok {
				padding = float32(pad)
			}
			elements := []ui.UIElem{}
			if elems, ok := table.RawGetString("elements").(*lua.LTable); ok {
				for i := 1; i <= elems.Len(); i++ {
					if elem, ok := elems.RawGetInt(i).(*lua.LTable); ok {
						if uielem, err := toUIElem(elem); err == nil {
							elements = append(elements, uielem)
						}
					}
				}
			}
			return &ui.Flexbox{Dir: ui.FlexDirection(direction), Anchor: ui.FlexAnchor(anchor), Padding: padding, Elements: elements}, nil
		}
	case "Button":
		{
			fontSize := float32(20)
			if fsize, ok := table.RawGetString("fontSize").(lua.LNumber); ok {
				fontSize = float32(fsize)
			}
			textOffset := float32(12)
			if tffset, ok := table.RawGetString("textOffset").(lua.LNumber); ok {
				textOffset = float32(tffset)
			}
			buttonType := ""
			if btype, ok := table.RawGetString("buttonType").(lua.LString); ok {
				buttonType = string(btype)
			}
			buttonIcon := ""
			if icon, ok := table.RawGetString("icon").(lua.LString); ok {
				buttonIcon = string(icon)
			}
			buttonText := ""
			if text, ok := table.RawGetString("text").(lua.LString); ok {
				buttonText = string(text)
			}

			return &ui.Button{
				FontSize:   fontSize,
				TextOffset: textOffset,
				Base:       assets.Manager.GetTexture(buttonType + "Button"),
				Hover:      assets.Manager.GetTexture(buttonType + "ButtonHover"),
				Pressed:    assets.Manager.GetTexture(buttonType + "ButtonPressed"),
				Icon:       assets.Manager.GetTexture(buttonIcon),
				Text:       buttonText,
			}, nil
		}
	}

	return &ui.Box{}, errors.New("ui elemnt not found")
}

func GetWindow(name string) (ui.Window, error) {
	if tbl, ok := L.GetGlobal(name).(*lua.LTable); ok {
		if window, err := toWindow(tbl, name); err == nil {
			return window, nil
		}
	}
	return ui.Window{}, errors.New("Did not find window named " + name)
}
