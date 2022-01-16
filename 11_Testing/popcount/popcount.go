package popcount

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// another init way
var pc2 = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return pc
}()

// PopCount returns the population count (number of set bits) of x
func PopCount(x uint64) int {
	/*
		byte(x) wil truncate x, remains low 8 bits
		for any 8 bits int, pc has saved the result of count of bit 1
	*/
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	/*
		byte(x) wil truncate x, remains low 8 bits
		for any 8 bits int, pc has saved the result of count of bit 1
	*/
	r := byte(0)
	for i := 0; i < 8; i++ {
		x >>= i * 8
		r += pc2[byte(x)]
	}
	return int(r)
}

func PopCount3(x uint64) int {
	r := 0
	for x != 0 {
		if x&1 == 1 {
			r++
		}
		x >>= 1
	}
	return r
}

func PopCount4(x uint64) int {
	r := 0
	for x != 0 {
		r++
		x = x & (x - 1)
	}
	return r
}
