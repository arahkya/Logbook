# Logbook
Application to save log of anything you want.

## SSL
### Gen Key
`openssl genrsa -out logbook.key 4096
### Gen PEM file
`openssl req -x509 -new -days 365 -key .\logbook.key -outform PEM -out logbook.pem -config .\logbook.cnf -extensions v3_req`
