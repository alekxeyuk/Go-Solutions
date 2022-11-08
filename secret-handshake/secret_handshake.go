package secret

var ops = [4]string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

func Handshake(code uint) []string {
	out := []string{}

	reversed := code&16 == 16

	for i := 0; i < 4; i++ {
		if code&1 == 1 {
			if reversed {
				out = append([]string{ops[i]}, out...)
			} else {
				out = append(out, ops[i])
			}
		}
		code >>= 1
	}

	return out
}
