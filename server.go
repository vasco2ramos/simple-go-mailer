package main

import (
    "bytes"
    "log"
    "net/http"
    "html/template"
    "gopkg.in/gomail.v2"
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
  d := gomail.NewPlainDialer("smtp.gmail.com", 587, "**@gmail.com", "**")

  // Send the email.
  if err := d.DialAndSend(m); err != nil {
      panic(err)
  }
}


func getEmailPost(rw http.ResponseWriter, req *http.Request) {

    // Parsing Post Data
    err := req.ParseForm()
    check(err)

    // Reading Html Template
    t := template.New("reportEmail.html") //create a new template
    t, err = t.ParseFiles("tmpl/reportEmail.html") //open and parse a template text file
    check(err)

    // Build Parameters from Post
    parameters := struct {
        Title string
        Client string
        Report string
    }{
        "New Candidate",
        "Lorem Ipsum",
        //req.PostFormValue("clientName"),
        "Lorem Ipsum Dolor Sit Amet MotherFuckers",
    }

    // Creating Email Document from Template
    var doc bytes.Buffer
    err = t.Execute(&doc, &parameters)
    check(err)
    s := doc.String()

    sendEmail("vascoasramos@gmail.com", "vasco@tyba.com", "New Candidate", s)


}

func main() {
    http.HandleFunc("/email", getEmailPost)
    log.Fatal(http.ListenAndServe(":8082", nil))
}
