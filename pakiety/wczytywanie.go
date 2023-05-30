// Plik wczytywanie.go
package wczytywanie
import (
	"os"
	"bufio"
	"strings"
	"strconv"
	// "github.com/fstanis/screenresolution"
	"github.com/veandco/go-sdl2/sdl" 
	"log"
)

// Funkcja wczytująca dane z pliku do tablicy dwuwymiarowej
func WczytajMapę (nazwaPliku string) ([] [] int, error) {
	file, err := os.Open (nazwaPliku)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	//tablica 10x10
	swiat := make([] [] int, 10)
	for i := range swiat {
		swiat [i] = make([]int, 10)
	}


	scanner := bufio.NewScanner(file)
	wiersz := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		for kolumna, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			swiat [wiersz][kolumna] = num
		}
		wiersz++
	}


	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	return swiat, nil
}


func RysujOkno (dawajokno chan *sdl.Window) () { 
	// resolution := screenresolution.GetPrimary() 
	// width := int32(resolution.Width)-300 //ZMIANA SZEROKOSCI
	// height := int32 (resolution.Height)-600 //ZMIANA WYSOKOSCI 
	width := int32(850)
	height := int32(850)
	if err := sdl. Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	} 
	defer sdl.Quit() 
	
	window, err := sdl.CreateWindow("Gra", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil { 
		log.Fatal(err)
	}		
	defer window.Destroy()
	
	surface, err := window.GetSurface() 
	if err != nil { 
		log.Fatal(err)
	} 

	surface.FillRect(nil, 0) 
	dawajokno <- window
	
	running := true 
	for running { 
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() { 
			switch event. (type) { 
			case *sdl.QuitEvent: 
				running = false
				break 
			} 
		} 
	} 
} 

func RysujPole (wysokosc int, szerokosc int, typ int, window *sdl.Window) (error) { 
	szerokosc = szerokosc*50+(10*(szerokosc+1))
	wysokosc = wysokosc*50+(10*(wysokosc+1))
	
	rect := sdl.Rect{int32(szerokosc), int32(wysokosc), 50, 50}
	
	var color sdl.Color
	
	switch typ { 
	case 0: //WODA 
		color = sdl.Color{R: 0, G: 183, B: 229, A: 255} 
	case 1://TRAWA 
		color = sdl.Color{R: 0, G: 204, B: 0, A: 255} 
	case 2: //WZGORZE 
		color = sdl.Color{R: 204, G: 102, B: 0, A: 255}
	}
	

	surface, err := window.GetSurface() 
	if err != nil { 
		return err 
	} 

	pixel := sdl.MapRGBA(surface.Format, color.R, color.G, color.B, color.A)
	surface.FillRect(&rect, pixel) 
	
	window.UpdateSurface() 
	

	return nil
}
