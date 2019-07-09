package main

import (
		"github.com/valyala/fasttemplate"
		"log"
		"fmt"
		"io"
		"os/exec"
		)
		
func main(){
	template := "<h1>Hello, [user]!</h1> <b>You won [prize]!!!</b> [foobar]"
	t, err := fasttemplate.NewTemplate(template, "[", "]")
	if err != nil {
		log.Fatalf("unexpected error when parsing template: %s", err)
	}
	message := t.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
		switch tag {
		case "user":
			return w.Write([]byte("John"))
		case "prize":
			return w.Write([]byte("$100500"))
		default:
			return w.Write([]byte(fmt.Sprintf("[unknown tag %q]", tag)))
		}
	})
	cmd:= exec.Command("sendmail", "villasenor.ariel@gmail.com")
	stdin, err :=cmd.StdinPipe()
	if err!=nil{
		log.Fatal(err)
	}
	go func(){
		defer stdin.Close()
		email := "From: arielv@msi-GT70\nTo: villasenor.ariel@gmail.com\nSubject: Test 573\nMime-Version: 1.0\nContent-Type: text/html\n" + message +"\n."
		io.WriteString(stdin, email)
	}()
    out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
	fmt.Printf("%s", message)
}