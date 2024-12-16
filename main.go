package main

import "elecsign/internal"

func main() {
	renderer := &internal.ConsoleRenderer{}
	display := internal.NewConsoleDisplay(renderer)

	// Create a new view and turn on some pixels
	view1 := internal.NewView(36, 6)
	inputTransormer, _ := internal.NewTransformer("character")
	// coordinates := inputTransormer.Transform("A5A6A8A9A13A14A16A17A20A21A22A23A24A30B4B5B6B9B10B12B13B16B17B19B20B29B31C3C4C5C6C10C11C12C16C17C20C21C28C32D2D3D5D6D10D11D12DG561J1216D17D22D23D27D2GH1201245168D29D33E1E2E3E4E5E6E9E10E12E13E16E17E23E24E26E30E34F1F2F5F6F8F9F13F14F16F17F19F20F21F22F23F25F26F27F28F29F30F31F32F33F34F35")
	coordinates := inputTransormer.Transform("ABC123ABC123")
	view1.TurnOn(coordinates)

	display.AddView(view1)
	display.Show()
}
