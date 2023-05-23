package main

import . "swigtests/li_cdata_cpp"

func main() {
	s := []byte("ABC\x00abc")
	m := Malloc(256)
	Memmove(m, s)
	ss := Cdata(m, 7)
	if string(ss) != "ABC\x00abc" {
		panic("failed")
	}
}
