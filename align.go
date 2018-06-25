package gonwr

import "fmt"

func idx(i, j, blen int) int {
	return (i * blen) + j
}

type cell struct {
	next, score int
}

// Align takes two Rune slices as well as three integers for the Needleman-Wunsch scoring function.
// It also takes a Rune that sets the filler character, e.g. rune('#')
// It returns the final score, as well as two aligned Rune slices.
func Align(a, b []rune, filler rune, match, mismatch, gap int) (runeSl1, runeSl2 []rune, score int) {
	tbmap := make(map[int]cell)
	// adding first element
	a = append([]rune{rune(' ')}, a...)
	b = append([]rune{rune(' ')}, b...)
	alen := len(a)
	blen := len(b)
	// initialise matrix, optimise here should 2*blen
	f := make([]int, 2*blen)
	// fill matrix
	f[0] = 0
	rowcount := 1
	// optimise here len(f) should be alen*blen
	for i := 1; i < len(f); i++ {
		if i < blen {
			tbmap[i] = cell{next: i - 1, score: gap * i}
			f[i] = gap * i
			continue
		}
		// optimise here, that is a row change so the row to fill is f[blen] if we keep it to 2*blen, copy over values
		if i%blen == 0 {
			for j := range f[0:blen] {
				f[j] = f[j+blen]
			}
			f[blen] = gap * rowcount
			tbmap[i] = cell{next: i - blen, score: gap * rowcount}
			rowcount++
			continue
		}
		indexB := i % blen
		indexA := i / blen
		score := match
		if a[indexA] != b[indexB] {
			score = mismatch
		}
		// remains
		left := score + tbmap[i-1].score
		up := score + tbmap[i-blen].score
		diag := score + tbmap[i-(blen+1)].score
		result := diag
		nextCell := i - (blen + 1)
		if diag < left {
			result = left
			nextCell = i - 1
			if left < up {
				result = up
				nextCell = i - blen
			}
		}
		if diag < up && diag >= left {
			result = up
			nextCell = i - blen
		}
		tbmap[i] = cell{next: nextCell, score: result}
	}
	fmt.Println()
	for k, v := range tbmap {
		fmt.Printf("key[%d] value[%d]\n", k, v)
	}
	path := []int{}
	// optimise here len(f) should be alen*blen, try to get rid of f
	start := (alen * blen) - 1
	fmt.Println("here", start)
	for start != 0 {
		score = score + tbmap[start].score
		path = append(path, start)
		start = tbmap[start].next
		fmt.Println("new", start)
	}
	for i := len(path) - 1; i >= 0; i-- {
		indexA := path[i] / blen
		if indexA >= 0 && indexA < len(a) {
			runeSl1 = append(runeSl1, a[indexA])
			a[indexA] = filler
		}
		indexB := (path[i] % blen)
		if indexB >= 0 && indexB < len(b) {
			runeSl2 = append(runeSl2, b[indexB])
			b[indexB] = filler
		}
	}
	return runeSl1, runeSl2, score
}
