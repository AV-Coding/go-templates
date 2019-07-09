package main

import (
	"fmt"
	"os/exec"
	"io"
	"html/template"
	"log"
	"bytes"
)

type Inventory struct{
	Material string
	Count int
	Verb string
}

func main(){
	var output bytes.Buffer
	sweater := Inventory{"wool", 17, "burn"}
	script, err := template.New("testing").Parse("{{ .Count }} items are made of {{ .Material }}. How many would you like to {{ .Verb }}?")
	if err != nil{
		log.Fatal(err)
	}
	script.Execute(&output, sweater)
	message := output.String()
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
}