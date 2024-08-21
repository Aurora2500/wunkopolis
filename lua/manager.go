package lua

import (
	"errors"
	"wunkopolis/assets"
	"wunkopolis/ui"
	"wunkopolis/variables"

	rl "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
)

var L *lua.LState

func init() {
	L = lua.NewState()
	defer L.Close()
	L.DoFile("lua/main.lua")

	L.SetGlobal("SetVariable", L.NewFunction(setVaribleLua))

}

func setVaribleLua(L *lua.LState) int {
	variables.Variables[L.ToString(1)] = L.ToInt(2)
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

	width := float32(10)
	if num, ok := table.RawGetString("width").(lua.LNumber); ok {
		width = float32(num)
	} else {
		return window, errors.New("Could not find width for " + name)
	}
	height := float32(10)
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
	case "Text":
		{
			content := ""
			if str, ok := table.RawGetString("text").(lua.LString); ok {
				content = string(str)
			}
			size := 16
			if s, ok := table.RawGetString("size").(lua.LNumber); ok {
				size = int(s)
			}
			return &ui.Text{Content: content, Size: float32(size)}, nil
		}
	case "Flexbox":
		{
			border := 10
			if bor, ok := table.RawGetString("border").(lua.LNumber); ok {
				border = int(bor)
			}
			direction := 0
			if dir, ok := table.RawGetString("direction").(lua.LNumber); ok {
				direction = int(dir)
			}
			background := ui.NPatchBox{}
			if text, ok := table.RawGetString("background").(lua.LString); ok {
				texture := assets.Manager.GetTexture(string(text))
				background = ui.NPatchBox{
					Texture: texture, NPatchInfo: rl.NPatchInfo{
						Source: ui.Area{
							X:      0,
							Y:      0,
							Width:  float32(texture.Width),
							Height: float32(texture.Height),
						},
						Left:   int32(border),
						Right:  int32(border),
						Top:    int32(border),
						Bottom: int32(border),
					},
				}
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
			return &ui.Flexbox{Background: background, Direction: ui.FlexDirection(direction), Anchor: ui.FlexAnchor(anchor), Padding: padding, Elements: elements, Border: float32(border)}, nil
		}
	case "Scrollbox":
		{
			border := 10
			if bor, ok := table.RawGetString("border").(lua.LNumber); ok {
				border = int(bor)
			}
			background := ui.NPatchBox{}
			if text, ok := table.RawGetString("background").(lua.LString); ok {
				texture := assets.Manager.GetTexture(string(text))
				background = ui.NPatchBox{
					Texture: texture, NPatchInfo: rl.NPatchInfo{
						Source: ui.Area{
							X:      0,
							Y:      0,
							Width:  float32(texture.Width),
							Height: float32(texture.Height),
						},
						Left:   int32(border),
						Right:  int32(border),
						Top:    int32(border),
						Bottom: int32(border),
					},
				}
			}
			direction := 0
			if dir, ok := table.RawGetString("direction").(lua.LNumber); ok {
				direction = int(dir)
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
			return &ui.Scrollbox{Background: background, Direction: ui.FlexDirection(direction), Padding: padding, Elements: elements, Border: float32(border)}, nil
		}
	case "Button":
		{
			fontSize := float32(20)
			if fsize, ok := table.RawGetString("fontSize").(lua.LNumber); ok {
				fontSize = float32(fsize)
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
				FontSize: fontSize,
				Base:     assets.Manager.GetTexture(buttonType + "Button"),
				Hover:    assets.Manager.GetTexture(buttonType + "ButtonHover"),
				Pressed:  assets.Manager.GetTexture(buttonType + "ButtonPressed"),
				Icon:     assets.Manager.GetTexture(buttonIcon),
				Text:     buttonText,
			}, nil
		}
	case "Tabs":
		{
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
			tabNames := []string{}
			if elems, ok := table.RawGetString("names").(*lua.LTable); ok {
				for i := 1; i <= elems.Len(); i++ {
					if elem, ok := elems.RawGetInt(i).(lua.LString); ok {
						tabNames = append(tabNames, string(elem))
					}
				}
			}
			for i := len(tabNames) - 1; i < len(elements); i++ {
				tabNames = append(tabNames, "")
			}
			fontSize := 16
			if fsize, ok := table.RawGetString("fontSize").(lua.LNumber); ok {
				fontSize = int(fsize)
			}
			tabs := ui.Tabs{
				Tabs:       elements,
				TabButtons: ui.Flexbox{Anchor: ui.Center, Padding: 8},
				TabNames:   tabNames,
				FontSize:   float32(fontSize),
			}
			tabs.Setup()
			return &tabs, nil
		}
	case "NPatchRect":
		{
			textureName := ""
			if tex, ok := table.RawGetString("texture").(lua.LString); ok {
				textureName = string(tex)
			}
			border := 10
			if bor, ok := table.RawGetString("borderSize").(lua.LNumber); ok {
				border = int(bor)
			}
			texture := assets.Manager.GetTexture(textureName)
			return &ui.NPatchBox{Texture: texture, NPatchInfo: rl.NPatchInfo{
				Source: ui.Area{
					X:      0,
					Y:      0,
					Width:  float32(texture.Width),
					Height: float32(texture.Height),
				},
				Left:   int32(border),
				Right:  int32(border),
				Top:    int32(border),
				Bottom: int32(border),
			}}, nil
		}
	case "Image":
		{
			textureName := ""
			if tex, ok := table.RawGetString("texture").(lua.LString); ok {
				textureName = string(tex)
			}
			texture := assets.Manager.GetTexture(textureName)
			return &ui.Image{Image: texture}, nil
		}
	}

	return &ui.NPatchBox{}, errors.New("ui elemnt not found")
}

func GetWindow(name string) (ui.Window, error) {
	if tbl, ok := L.GetGlobal(name).(*lua.LTable); ok {
		if window, err := toWindow(tbl, name); err == nil {
			return window, nil
		}
	}
	return ui.Window{}, errors.New("Did not find window named " + name)
}
