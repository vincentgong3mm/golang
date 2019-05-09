package dominion

type DrawObject interface {
	DrawRect() error
}

func ConvertTermString(width int, str string) string {
	ln := make([]byte, 0, 256)

	lenstr := len(str)
	lenspace := (width - lenstr) / 2

	ln = append(ln, '|')

	space := make([]byte, lenspace)
	for i := range space {
		space[i] = ' '
	}

	space2 := make([]byte, lenspace+1, lenspace+1)
	for i := range space2 {
		space2[i] = ' '
	}
	ln = append(ln, space...)

	tmpName := []byte(str)
	ln = append(ln, tmpName...)

	if lenstr%2 == 0 {
		ln = append(ln, space...)
	} else {
		ln = append(ln, space2...)
	}

	ln = append(ln, '|')

	return string(ln)
}
