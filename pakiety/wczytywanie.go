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
	"math/rand"
	"time"
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

func Rysuj (wysokosc int, szerokosc int, jednostki [][]int, swiat [][]int, window *sdl.Window) (error) {
	if swiat[wysokosc][szerokosc] >= 20 {
		
	typ_swiat := swiat[wysokosc][szerokosc]
	szerokosc_swiat := szerokosc*20+(2*(szerokosc+1))
	wysokosc_swiat := wysokosc*20+(2*(wysokosc+1))
	
	rect := sdl.Rect{int32(szerokosc_swiat), int32(wysokosc_swiat), 20, 20}
	
	var color sdl.Color
	
	switch typ_swiat { 
	case 20: //WODA 
		color = sdl.Color{R: 0, G: 183, B: 229, A: 255} 
	case 21://TRAWA 
		color = sdl.Color{R: 0, G: 204, B: 0, A: 255} 
	case 22: //WZGORZE 
		color = sdl.Color{R: 204, G: 102, B: 0, A: 255}
	}
	
	var surface *sdl.Surface
	var err error
	surface, err = window.GetSurface() 
	if err != nil { 
		return err 
	} 

	pixel := sdl.MapRGBA(surface.Format, color.R, color.G, color.B, color.A)
	surface.FillRect(&rect, pixel) 
	
	window.UpdateSurface() 
	
	typ_jednostki := jednostki[wysokosc][szerokosc]
	szerokosc_jednostki := szerokosc*10+(12*(szerokosc+1))
	wysokosc_jednostki := wysokosc*10+(12*(wysokosc+1))
	
	rect = sdl.Rect{int32(szerokosc_jednostki), int32(wysokosc_jednostki), 10, 10}
	
	switch typ_jednostki { 
	case 0: //BRAK JEDNOSTKI
		return nil
	case 1://JEDNOSTKA(MAIN)
		color = sdl.Color{R: 88, G: 65, B: 18, A: 220}
	case 2://JEDNOSTKA(WROG_BYLE_POLE)
		color = sdl.Color{R: 51, G: 0, B: 0, A: 220}
	case 3://JEDNOSTKA(WROG)
		color = sdl.Color{R: 236, G: 0, B: 0, A: 220} 
	}
	

	surface, err = window.GetSurface() 
	if err != nil { 
		return err 
	} 

	pixel = sdl.MapRGBA(surface.Format, color.R, color.G, color.B, color.A)
	surface.FillRect(&rect, pixel) 
	
	window.UpdateSurface() 
	
	}
	return nil

}

func czypuste (jednostki [][]int, i int, j int, kto int, swiat [][]int) (bool) { //i, j kordynaty
	if (jednostki[i][j]==0 || (jednostki[i][j]==3 && kto == 1) || jednostki[i][j]==2) && (swiat[i][j] != 0 && swiat[i][j] != 20) {
		return true
	} else {
		return false
	   }
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

func PrzemiescGracza (znak rune, jedn *[][]int, window *sdl.Window, swiat [][]int) {
    jednostki := *jedn
    i, j := znajdzgracza(jednostki)

    switch znak {
    case 's'://S
        if czypuste(jednostki, i+1, j, 1, swiat) {
            jednostki[i][j]=0
            jednostki[i+1][j]=1
			Rysuj(i, j, jednostki, swiat, window)
			Rysuj(i+1, j, jednostki, swiat, window)
        }
    case 'w': //W
        if czypuste(jednostki, i-1, j, 1, swiat) {
            jednostki[i][j]=0
            jednostki[i-1][j]=1
			Rysuj(i, j, jednostki, swiat, window)
			Rysuj(i-1, j, jednostki, swiat, window)
        }
    case 'a': //A
        if czypuste(jednostki, i, j-1, 1, swiat) {
            jednostki[i][j]=0
            jednostki[i][j-1]=1
			Rysuj(i, j, jednostki, swiat, window)
			Rysuj(i, j-1, jednostki, swiat, window)
         }
    case 'd': //D
        if czypuste(jednostki, i, j+1, 1, swiat) {
            jednostki[i][j]=0
            jednostki[i][j+1]=1
			Rysuj(i, j, jednostki, swiat, window)
			Rysuj(i, j+1, jednostki, swiat, window)
        }
    }
}

func wzrok (i int, j int, jednostki [][]int, swi *[][]int, window *sdl.Window) () {
	swiat := *swi
	if swiat[i][j] < 20 {
		swiat[i][j]=swiat[i][j]+20
		Rysuj(i, j, jednostki, swiat, window)
	}
}

func Widocznosc (jednostki[][]int, swi *[][]int, window *sdl.Window) ()  {
	i, j := znajdzgracza(jednostki)
	// swiat := *swi
	// wzrok(i, j, jednostki, swiat, window)
	// wzrok(i+1, j, jednostki, swiat, window)
	// wzrok(i, j+1, jednostki, swiat, window)
	// wzrok(i-1, j, jednostki, swiat, window)
	// wzrok(i, j-1, jednostki, swiat, window)
	wzrok3(i, j, jednostki, swi, window, 2)
}

func wzrok3 (i int, j int, jednostki [][]int, swiat *[][]int, window *sdl.Window, zasieg int) () {
	swi := *swiat
	wzrok(i, j, jednostki, swiat, window)
	
	if zasieg >= 0 && i>0 && j>0 && i<39 && j<39 && (zasieg==2 || (swi[i][j]!=2 && swi[i][j]!=22)) {
		wzrok3(i-1, j, jednostki, swiat, window, zasieg-1)
		wzrok3(i+1, j, jednostki, swiat, window, zasieg-1)
		wzrok3(i, j-1, jednostki, swiat, window, zasieg-1)
		wzrok3(i, j+1, jednostki, swiat, window, zasieg-1)
	}
}

func SzukajWroga (i int, j int, jednostki [][]int) (int, int) {
	for a:=i; a<len(jednostki); a++ {
		for b:=j; b<len(jednostki[a]); b++ {
			if jednostki[a][b]==3 {
				return a, b
			}
		}
	}
	return -1, -1
}

func PrzemiescWroga (jedn *[][]int, swiat [][]int, window *sdl.Window) () {
	jednostki := *jedn
	// i, j := 0, 0
	
	for i, _ := range jednostki {
		for j, _ := range jednostki {
				if jednostki[i][j]==3 {
					defer RuchWroga(i, j, jedn, swiat, window)	
				}
		}
	}
	
			// i, j = SzukajWroga(i, j, jednostki)
}

func RuchWroga (i int, j int, jedn *[][]int, swiat [][]int, window *sdl.Window) () {
	jednostki := *jedn
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	random := r1.Intn(4) + 1
	
	switch random {
	case 1:
		if czypuste(jednostki, i-1, j, 3, swiat) {
			jednostki[i][j] = 2
			err := Rysuj(i, j, jednostki, swiat, window)
			if err != nil {
				panic(err)
			}
			jednostki[i-1][j]=3
			err = Rysuj(i-1, j, jednostki, swiat, window)
			if err != nil {
				panic(err)
			}
		}
	case 2:
		if czypuste(jednostki, i+1, j, 3, swiat) {
			jednostki[i][j] = 2
			err := Rysuj(i, j, jednostki, swiat, window)
			if err != nil {
				panic(err)
			}
			jednostki[i+1][j]=3
			err = Rysuj(i+1, j, jednostki, swiat, window)
			if err != nil {
				panic(err)
			}
		}
	case 3:
		if czypuste(jednostki, i, j-1, 3, swiat) {
			jednostki[i][j] = 2
			err := Rysuj(i, j, jednostki, swiat, window)
			if err != nil {
				panic(err)
			}
			jednostki[i][j-1]=3
			err = Rysuj(i, j-1, jednostki, swiat, window)
			if err != nil {
				panic(err)
			}
		}
	case 4:
		if czypuste(jednostki, i, j+1, 3, swiat) {
			jednostki[i][j] = 2
			err := Rysuj(i, j, jednostki, swiat, window)
			if err != nil {
				panic(err)
			}
			jednostki[i][j+1]=3
			err = Rysuj(i, j+1, jednostki, swiat, window)
			if err != nil {
				panic(err)
			}
		}
	}
}

func UsunSlady (jedn *[][]int, swiat [][]int, window *sdl.Window) () {
	jednostki := *jedn
	
	for i, _ := range jednostki {
		for j, _ := range jednostki {
			if jednostki[i][j]==2 {
				jednostki[i][j]=0
				Rysuj(i, j, jednostki, swiat, window)
			}
		}
	}
}