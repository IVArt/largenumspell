package largenumspell

import "testing"

var get_res []string

func BenchmarkSpellout(b *testing.B) {
	var res []string
	for i := 0; i < b.N; i++ {
		res, _ = Spellout("-11000555555546733987333.022173455555892222e+1098755555789")
	}
	get_res = res
}
