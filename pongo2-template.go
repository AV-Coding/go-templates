package main

import (
	"fmt"
	"github.com/flosch/pongo2"
	"io"
	"os/exec"
	"log"
)

func constructemail(message string, from string, to string, subject string, html bool) string{
	var email string
	if html{
		email = "From: " + from + "\nTo: " + to + "\nSubject: " + subject + "\nMime-Version: 1.0\nContent-Type: text/html\n" + message +"\n."
	}else{
		email = "From: " + from + "\nTo: " + to + "\nSubject: " + subject + "\n" + message +"\n."
	}
	return email
}
func sendmail(email string, to string){
	cmd:= exec.Command("sendmail", to)
	// Executes returns command struct to execute the program.
	stdin, err :=cmd.StdinPipe()
	//Creates a pipe to the command line
	if err!=nil{
		log.Fatal(err)
	}
	go func(){
		defer stdin.Close()
		io.WriteString(stdin, email)
	}()
	//Runs the command
    out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}

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
fmt.Println(output)
}