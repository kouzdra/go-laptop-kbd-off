package main

import "log"
//import "fmt"
import "os/exec"
import "regexp"

func main () {
	out, err := exec.Command("xinput", "list").Output()
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile("AT Translated Set 2 keyboard[ \t]+id=([0-9]+)")
	sm := re.FindSubmatch(out)
	//fmt.Printf("The submatch is %d\n----\n%q\n---\n", len (sm), sm)
	if len (sm) == 2 {
		err1 := exec.Command("xinput", "float", string(sm[1])).Run ()
		if err1 != nil {
			log.Fatal(err1)
		}
	}
}
