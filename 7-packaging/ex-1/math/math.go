package math

type Math struct {
	A int
	B int
}
func (m Math) Add() int {
	m.A = m.A + 1

	return m.A + m.B
}