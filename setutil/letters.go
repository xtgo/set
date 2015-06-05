package setutil

import "fmt"

type Letters []byte

func (s Letters) Len() int           { return len(s) }
func (s Letters) Less(i, j int) bool { return s[i] < s[j] }
func (s Letters) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Letters) String() string     { return string(s) }

func (s Letters) Mark(pivot, i, j int) string {
	t := append([]byte(nil), s...)
	switch {
	case i != j:
		t[j] -= 'a' - 'A'
		fallthrough
	case i >= 0:
		t[i] -= 'a' - 'A'
	}
	return fmt.Sprintf("[%s|%s]", t[:pivot], t[pivot:])
}
