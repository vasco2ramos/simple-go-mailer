package main

import (
    "bytes"
    "log"
    "net/http"
    "html/template"
    "gopkg.in/gomail.v2"
    "fmt"
    "encoding/json"
    "io/ioutil"
)

type Credentials struct {
    Host string `json:"host"`
    User string `json:"user"`
    Pass string `json:"pass"`
    Port int    `json:"port"`
}

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

  // This should be in it's own function
  file, e := ioutil.ReadFile("./credentials.json")
  check(e)
  var c Credentials
  json.Unmarshal(file, &c)
  // -----------------------------------

  d := gomail.NewPlainDialer(c.Host, c.Port, c.User, c.Pass)

  if err := d.DialAndSend(m); err != nil {
      panic(err)
  }

  fmt.Println("Mail Sent!")
}


func getReportTemplate(parameters interface{}) string {
  // Reading Html Template
  t := template.New("reportEmail.html") //create a new template
  t, err := t.ParseFiles("tmpl/reportEmail.html") //open and parse a template text file
  check(err)
  // Creating Email Document from Template
  var doc bytes.Buffer
  fmt.Println(parameters)
  err = t.Execute(&doc, &parameters)
  check(err)

  s := doc.String()

  return s
}


func getPostRequest(rw http.ResponseWriter, req *http.Request) {
    // Parsing Post Data
    err := req.ParseForm()
    check(err)
    // Build Parameters from Post
    parameters := struct {
        Client string
        Report string
    }{
        req.Form.Get("clientName"),
        req.Form.Get("report"),
    }
    // Build Report
    s := getReportTemplate(parameters)
    sendEmail("vascoasramos@gmail.com", "vasco@tyba.com", "Testing Emails", s)
}

func main() {
    http.HandleFunc("/email", getPostRequest)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
