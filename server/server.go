package main

import (
    "bytes"
    "log"
    "net/http"
    "html/template"
    "gopkg.in/gomail.v2"
    "fmt"
  //  "encoding/json"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func sendEmail(from string, to string, subject string, email string){
  m := gomail.NewMessage()
  m.SetHeader("From", from)
  m.SetHeader("To", to)
  m.SetHeader("Subject", subject)
  m.SetBody("text/html", email)

  d := gomail.NewPlainDialer("smtp.gmail.com", 587, "***@gmail.com", "***")

  if err := d.DialAndSend(m); err != nil {
      panic(err)
  }

  fmt.Println("Mail Sent!")
}


func getTemplate(templateName string) (*Template){
  // Reading Html Template
  t := template.New("reportEmail.html") //create a new template
  t, err := t.ParseFiles("tmpl/reportEmail.html") //open and parse a template text file
  check(err)

  return t
}


func getPostRequest(rw http.ResponseWriter, req *http.Request) {

    // Parsing Post Data
    err := req.ParseForm()
    check(err)

    t := getTemplate("reportEmail.html")

    // Build Parameters from Post
    parameters := struct {
        Title string
        Client string
        Report string
    }{
        "New Candidate",
        req.Form.Get("clientName"),
        req.Form.Get("report"),
    }
    fmt.Println(req.Form.Get("clientName"))
    // Creating Email Document from Template
    var doc bytes.Buffer
    err = t.Execute(&doc, &parameters)
    check(err)
    s := doc.String()

    sendEmail("vascoasramos@gmail.com", "vasco@tyba.com", "Testing Emails", s)
}

func main() {
    http.HandleFunc("/email", getPostRequest)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
