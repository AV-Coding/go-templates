package main

import (
	"fmt"
	"github.com/flosch/pongo2"
	"io"
	"os/exec"
	"log"
)

func main(){
	tmpl, err := pongo2.FromString("Hello {{ name|capfirst }}! How was the {{ type }} restaurant yesterday?\nDid they have {{ num1 }} or {{ num2 }} courses!?")
	if err != nil {
		panic(err)
	}
	output, err := tmpl.Execute(pongo2.Context{"name": "florian", "type": "mexican", "num1":1, "num2":2 })
	if err != nil {
		panic(err)
	}
	message := output
	cmd:= exec.Command("sendmail", "villasenor.ariel@gmail.com")
	stdin, err :=cmd.StdinPipe()
	if err!=nil{
		log.Fatal(err)
	}
	go func(){
		defer stdin.Close()
		email := "From: arielv@msi-GT70\nTo: villasenor.ariel@gmail.com\nSubject: Test 568\n" + message +"\n."
		io.WriteString(stdin, email)
	}()
    out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
fmt.Println(output)
}