simple-go-mailer
================

Dependencies
------------

"gopkg.in/gomail.v2"


Usage
-----

You need to create your own templates and add them to app/tmpl. There is already
a sample you can use. See: templating in go

- Run simple-mailer.go
- Change your credentials as per app/credentials.json
- Change the email endpoints on getPostRequest function
- Submit post requests to localhost:8080/email

Next Steps
----------

- Refactor some things (e.g. Email Endpoints)
- Testing
- Allow interaction with 3rd Party services.
