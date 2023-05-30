package rysowanie 

import ( 
	"github.com/fstanis/screenresolution"
	"github.com/veandco/go-sdl2/sdl" 
)

func RysujOkno () (error) { 
	resolution := screenresolution.GetPrimary() 
	width := int32(resolution.Width)-300 //ZMIANA SZEROKOSCI
	height := int32 (resolution.Height)-600 //ZMIANA WYSOKOSCI 
	if err := sdl. Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	} 
	defer sdl.Quit() 
	
	window, err := sdl.CreateWindow("Gra", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil { 
		return err
	}		
	defer window.Destroy()
	
	surface, err := window.GetSurface() 
	if err != nil { 
		return err 
	} 

	surface.FillRect(nil, 0) 

	running := true 

	for running { 
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() { 
			switch event. (type) { 
			case sdl.QuitEvent: 
				running = false
				break 
			} 
		} 
	} 
	return nil 
} 

// func RysujPole (wysokosc int, szerokosc int, typ int) (error) { 
	// rect := sdl.Rect{szerokosc, wysokosc, 100, 100}
	
	// switch typ { 
	// case 0: //WODA 
		// colour := sdl.Color{R: 0, G: 183, B: 229, A: 255} 
	// case 1://TRAWA 
		// colour := sdl.Color{R: 0, G: 204, B: 0, A: 255} 
	// case 2 //WZGORZE 
		// colour := sdl.Color{R: 204, G: 102, B: 0, A: 255}
	// }
	
	// pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	// surface.FillRect(&rect, pixel) 
	
	// window.UpdateSurface() 
// }
