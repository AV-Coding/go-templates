package main

import (
	"fmt"
	"os/exec"
	"io"
	"html/template"
	"log"
	"bytes"
)
//struct that holds different variables
//(This allows us to pass multiple values to our template.)
type Inventory struct{
	Material string
	Count int
	Verb string
}

func main(){
	var output bytes.Buffer
	sweater := Inventory{"wool", 17, "burn"}
	//script is the new template object.
	script, err := template.New("testing").Parse("<html><h1>{{ .Count }} items are made of {{ .Material }}.</h1> <b>How many would you like to {{ .Verb }}?</b></html>")
	if err != nil{
		log.Fatal(err)
	}
	//Here we execute our template with our "sweater" struct.
	//"output" will be a slice of bytes, it will contain the template with the updated data.
	script.Execute(&output, sweater)
	//Convert slice of bytes to a string.
	message := output.String()
	cmd:= exec.Command("sendmail", "villasenor.ariel@gmail.com")
	// Executes returns command struct to execute the program.
	stdin, err :=cmd.StdinPipe()
	//Creates a pipe to the command line
	if err!=nil{
		log.Fatal(err)
	}
	go func(){
		defer stdin.Close()
		email := "From: arielv@msi-GT70\nTo: villasenor.ariel@gmail.com\nSubject: Test 570\nMime-Version: 1.0\nContent-Type: text/html\n" + message +"\n."
		io.WriteString(stdin, email)
	}()
	//Runs the command
    out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}