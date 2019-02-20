package exregexp

import "fmt"
import "regexp"

// TestRegexp is
func TestRegexp() {

	str := `
	Index(0), Member(12345), IP(122.111.1.2), ActionNumber(1157)
	Index(0), Member(12345), IP(122.111.1.2), ActionNumber(1157)
	Index(0), Member(12345), IP(122.111.1.2), ActionNumber(1157)
	Index(0), Member(32345), IP(122.111.1.2), ActionNumber(1157)
	`

	// expStr := `ActionNumber\(\d+\)`
	expStr := `(Member\(\d+\))|ActionNumber\(\d+\)`
	reg := regexp.MustCompile(expStr)

	ss := reg.FindAllString(str, -1)

	for i := range ss {
		fmt.Printf("%d-%s\r\n", i, ss[i])
	}

}

func testFunction() {
	fmt.Println("testFunction")

	expStr := `PACKET_ENUM`
	regexp.MustCompile(expStr)

}
