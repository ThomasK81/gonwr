package gonwr

func idx(i, j, blen int) int {
	return (i * blen) + j
}

// Align takes two Rune slices as well as three integers for the Needleman-Wunsch scoring function.
// It also takes a Rune that sets the filler character, e.g. rune('#')
// It returns the final score, as well as two aligned Rune slices.
func Align(a, b []rune, filler rune, match, mismatch, gap int) (runeSl1, runeSl2 []rune, score int) {
	tbmap := make(map[int]int)
	// adding first element
	a = append([]rune{rune(' ')}, a...)
	b = append([]rune{rune(' ')}, b...)
	alen := len(a)
	blen := len(b)
	// initialise matrix
	f := make([]int, alen*blen)
	// fill matrix
	f[0] = 0
	rowcount := 1
	for i := 1; i < len(f); i++ {
		if i < blen {
			f[i] = gap * i
			tbmap[i] = i - 1
			continue
		}
		if i%blen == 0 {
			f[i] = gap * rowcount
			tbmap[i] = i - blen
			rowcount++
			continue
		}
		indexB := i % blen
		indexA := i / blen
		score := match
		if a[indexA] != b[indexB] {
			score = mismatch
		}
		left := score + f[i-1]
		up := score + f[i-blen]
		diag := score + f[i-(blen+1)]
		result := diag
		tbmap[i] = i - (blen + 1)
		if diag < left {
			result = left
			tbmap[i] = i - 1
			if left < up {
				result = up
				tbmap[i] = i - blen
			}
		}
		if diag < up && diag >= left {
			result = up
			tbmap[i] = i - blen
		}
		f[i] = result
	}
	path := []int{}
	start := len(f) - 1
	for start != 0 {
		score = score + f[start]
		path = append(path, start)
		start = tbmap[start]
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
