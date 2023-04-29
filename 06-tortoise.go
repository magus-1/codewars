package kata

import "fmt"

// 6 Kyu - Tortoise Racing

func Race(v1, v2, g int) [3]int {
	// Object 1:
	// 	Initial position: g (f)
	// 	Speed: v1 (f/s)
	// Object 2:
	//	Initial position: 0 (f)
	// 	Speed: v2 (f/s)
	// Return the time when both positions match, as the array [3]int{hour minute second}
	// Time equation: Xf = X0 + (v2-v1) * t
	// t = Xf / (v2-v1)
	tEnd := int(float64(g) / float64((v2 - v1)) * 3600)
	if tEnd < 0 {
		return [3]int{-1, -1, -1}
	}
	h := tEnd / 3600
	m := (tEnd % 3600) / 60
	s := ((tEnd % 3600) % 60) / 1
	fmt.Printf("tEnd: %v, h: %v, m: %v, s: %v \n", tEnd, h, m, s)
	return [3]int{h, m, s}

}
