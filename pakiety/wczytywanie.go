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
	
	//tablica 40x40
	swiat := make([] [] int, 40)
	for i := range swiat {
		swiat [i] = make([]int, 40)
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
	width := int32(882)
	height := int32(882)
	if err := sdl. Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	} 
	defer sdl.Quit() 
	
	window, err := sdl.CreateWindow("Podrobka cywilizacji", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
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
	szerokosc = szerokosc*20+(2*(szerokosc+1))
	wysokosc = wysokosc*20+(2*(wysokosc+1))
	
	rect := sdl.Rect{int32(szerokosc), int32(wysokosc), 20, 20}
	
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

func RysujJednostki (wysokosc int, szerokosc int, typ int, window *sdl.Window) (error) { 
	szerokosc = szerokosc*10+(12*(szerokosc+1))
	wysokosc = wysokosc*10+(12*(wysokosc+1))
	
	rect := sdl.Rect{int32(szerokosc), int32(wysokosc), 10, 10}
	
	var color sdl.Color
	
	switch typ { 
	case 0: //BRAK JEDNOSTKI
		return nil
	case 1://JEDNOSTKA(MAIN)
		color = sdl.Color{R: 88, G: 65, B: 18, A: 220}
	case 2://JEDNOSTKA(SOJUSZNIK)
		color = sdl.Color{R: 102, G: 255, B: 102, A: 220}
	case 3://JEDNOSTKA(WROG)
		color = sdl.Color{R: 236, G: 0, B: 0, A: 220} 
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

func czypuste (jednostki [][]int, i int, j int, kierunek rune) (bool) { //i, j kordynaty
	switch kierunek {
	case 'w':
		if jednostki[i-1][j]==0 {
			return true
		} else {
			return false
		  }
	case 's':
		if jednostki[i+1][j]==0 {
			return true
		} else {
			return false
		  }
	case 'd':
		if jednostki[i][j+1]==0 {
			return true
		} else {
			return false
		  }
	case 'a':
		if jednostki[i][j-1]==0 {
			return true
		} else {
			return false
		  }
	}
	return false

}

func znajdzgracza (jednostki [][]int) (int, int) {
	for i, _  := range jednostki {
        for j, _ := range jednostki {
			if jednostki[i][j]==1 {
				return i, j
			}
		}
	}
	return -1, -1
}

// func PrzemiescGracza (znak rune, jedn *[][]int, window *sdl.Window, swiat [][]int) () {
    // jednostki := *jedn
	// i, j := znajdzgracza(jednostki)
	
    // switch znak {
    // case 's'://S
        // if czypuste(jednostki, i, j, znak) {
			// jednostki[i][j]=0
            // jednostki[i+1][j]=1
			// RysujPole(i, j, swiat[i][j], window)
            // RysujJednostki(i+1, j, 1, window)
        // }
    // case 'w': //W
        // if czypuste(jednostki, i, j, znak) {
			// jednostki[i][j]=0
			// jednostki[i-1][j]=1
			// RysujPole(i, j, swiat[i][j], window)
			// RysujJednostki(i-1, j, 1, window)
        // }
    // case 'a': //A
        // if czypuste(jednostki, i, j, znak) {
            // jednostki[i][j]=0
            // jednostki[i][j-1]=1
            // RysujPole(i, j, swiat[i][j], window)
            // RysujJednostki(i, j-1, 1, window)
         // }
    // case 'd': //D
        // if czypuste(jednostki, i, j, znak) {
            // jednostki[i][j]=0
            // jednostki[i][j+1]=1
            // RysujPole(i, j, swiat[i][j], window)
            // RysujJednostki(i, j+1, 1, window)
        // }
    // }
// }

func PrzemiescGracza (znak rune, jedn *[][]int, window *sdl.Window, swiat [][]int) {
    jednostki := *jedn
    i, j := znajdzgracza(jednostki)

    switch znak {
    case 's'://S
        if czypuste(jednostki, i, j, znak) {
            jednostki[i][j]=0
            jednostki[i+1][j]=1
            RysujPole(i, j, swiat[i][j], window)
            RysujJednostki(i+1, j, 1, window)
        }
    case 'w': //W
        if czypuste(jednostki, i, j, znak) {
            jednostki[i][j]=0
            jednostki[i-1][j]=1
            RysujPole(i, j, swiat[i][j], window)
            RysujJednostki(i-1, j, 1, window)
        }
    case 'a': //A
        if czypuste(jednostki, i, j, znak) {
            jednostki[i][j]=0
            jednostki[i][j-1]=1
            RysujPole(i, j, swiat[i][j], window)
            RysujJednostki(i, j-1, 1, window)
         }
    case 'd': //D
        if czypuste(jednostki, i, j, znak) {
            jednostki[i][j]=0
            jednostki[i][j+1]=1
            RysujPole(i, j, swiat[i][j], window)
            RysujJednostki(i, j+1, 1, window)
        }
    }
}
