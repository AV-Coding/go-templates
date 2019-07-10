package main

import (
	"flag"
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
	First string
	Last string
	User string
	Status string
	Test string
}

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

func getflags()(string, string, string, string, string, string, string, bool){
	firstptr := flag.String("first", "first", "firstname")
	lastptr := flag.String("last", "last", "lastname")
	userptr := flag.String("user", "user", "username")
	statusptr := flag.String("status", "status", "status of instance")
	fromptr := flag.String("from", "arielv@msi-GT70", "from address")
	toptr := flag.String("to", "villasenor.ariel@gmail.com", "to address")
	subjectptr := flag.String("subject", "Cyverse", "Subject")
	boolptr := flag.Bool("html", false, "a bool")
	flag.Parse()
	return *firstptr, *lastptr, *userptr, *statusptr, *fromptr, *toptr, *subjectptr, *boolptr
}
func main(){
	var output bytes.Buffer
	first, last, user, status, from, to, subject, html := getflags()
	sweater := Inventory{First: first, Last: last, User: user, Status: status, Test: "Test"}
	// "Test" variable doesn't exist in the template. Does not cause any issues
	script, err := template.ParseFiles("template-html.tmpl")
	if err != nil{
		log.Fatal(err)
	}
	//Here we execute our template with our "sweater" struct.
	//"output" will be a slice of bytes, it will contain the template with the updated data.
	script.Execute(&output, sweater)
	fmt.Println(output.String())
	//Convert slice of bytes to a string.
	message := output.String()
	email := constructemail(message, from, to, subject, html)
	sendmail(email, to)
}