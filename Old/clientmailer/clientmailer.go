package clientmailer

import (
      //"encoding/json"
	    //"io/ioutil"
)

type client struct {
        Id      int
        Email   string
        Name    string
}

type candidate struct{
        Id      int
        Email   string
        Name    string
}

type email struct{
      Candidate     candidate
      Client        client
      Report        string
}



func NewEmail(cl client, ca candidate, rep string) *email {
  return &email{
      Client: cl,
      Candidate: ca,
      Report: rep,
  }
}


/*
func main() {
  m := gomail.NewMessage()
  m.SetHeader("From", "alex@example.com")
  m.SetHeader("To", "vasco@tyba.com")
  m.SetHeader("Subject", "Hello!")
  m.SetBody("text/html", "Hello <b>Me</b>!")
  d := gomail.NewPlainDialer("smtp.gmail.com", 587, "vascoasramos@gmail.com", "pai121natal")

  // Send the email to Bob, Cora and Dan.
  if err := d.DialAndSend(m); err != nil {
      panic(err)
  }

}
*/
