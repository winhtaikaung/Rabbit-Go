package main

// https://golang.org/pkg/fmt/

import (
	"fmt"
	"rabbit"
	"regexp"
)

func main() {
	// x := "World"
	// y := 70
	// var s string
	// var i int
	// fmt.Println("Hello World\n")
	// fmt.Print("Hello World\n")
	// fmt.Printf("Hello World\n")
	// fmt.Printf("Hello %s\n", "World")
	// fmt.Printf("Hello %s\n", x)
	// fmt.Printf("My Age %d\n", y)
	// fmt.Sscanf(" 1234567 ", "%5s%d", &s, &i
	// fmt.Printf("var s=%s\n", s)
	// fmt.Printf("var i=%d\n", i)
	// fmt.Printf("%6.2f", 12.0285)
	// rabbit.Uni2zg("asdf")
	fmt.Println(rabbit.Zg2uni("သီဟိုဠ္မွ ဉာဏ္ႀကီးရွင္သည္ အာယုဝၯနေဆးၫႊန္းစာကို ဇလြန္ေဈးေဘး ဗာဒံပင္ထက္ အဓိ႒ာန္လ်က္ ဂဃနဏဖတ္ခဲ့သည္။"))

	// fmt.Println(rabbit.Uni2zg("သီဟိုဠ်မှ ဉာဏ်ကြီးရှင်သည် အာယုဝဍ္ဎနဆေးညွှန်းစာကို ဇလွန်ဈေးဘေး ဗာဒံပင်ထက် အဓိဋ္ဌာန်လျက် ဂဃနဏဖတ်ခဲ့သည်။"))
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
}
