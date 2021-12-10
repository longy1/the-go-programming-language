package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return len(s.words) > word && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for len(s.words) <= word {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			for j := 0; j < 64; j++ {
				x := 64*i + j
				if t.Has(x) {
					s.Remove(x)
				}
			}
		} else {
			break
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	var symm []int
	for i := range s.words {
		if i < len(t.words) {
			for j := 0; j < 64; j++ {
				x := 64*i + j
				if s.Has(x) && t.Has(x) {
					symm = append(symm, x)
				}
			}
		} else {
			break
		}
	}
	s.UnionWith(t)
	for i := range symm {
		s.Remove(symm[i])
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for bit := 0; bit < 64; bit++ {
			if word&(1<<uint(bit)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				_, err := fmt.Fprintf(&buf, "%d", 64*i+bit)
				if err != nil {
				}
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len can be implemented by saving len var into struct, or saving hash map
func (s *IntSet) Len() int {
	var l int
	for i := range s.words {
		for j := 0; j < 64; j++ {
			if s.words[i]&(1<<j) != 0 {
				l++
			}
		}
	}
	return l
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var c IntSet
	c.words = append(c.words, s.words...)
	return &c
}

func (s *IntSet) AddAll(vals ...int) {
	for _, x := range vals {
		s.Add(x)
	}
}

func (s *IntSet) Elems() []int {
	var elems []int
	for i := range s.words {
		for j := 0; j < 64; j++ {
			x := i*64 + j
			if s.Has(x) {
				elems = append(elems, x)
			}
		}
	}
	return elems
}

func main() {
	var x, y IntSet
	x.AddAll(1, 2, 3)
	y.AddAll(2, 3, 4, 5, 6)
	x.SymmetricDifference(&y)
	for k, v := range y.Elems() {
		fmt.Println(k, v)
	}
	fmt.Println(&x)
}
