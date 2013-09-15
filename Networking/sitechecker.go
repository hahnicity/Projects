package main

import (
    "flag"
    "github.com/gcmurphy/getpass"
    "net/http"
    "net/smtp"
    "strconv"
    "time"
)

// XXX This whole array of constants can just be a struct
var (
    from          string = "grehm87@gmail.com"
    mailPort      string = "587"
    mailServer    string = "smtp.gmail.com"
    pass          string
    sleepDuration int
    to            string = "grehm87@gmail.com"
    url           string
)

type Body struct {
    to, subject, msg string    
}

func (b *Body) toByteArray() []byte {
    return []byte("To: " + b.to + "\r\nSubject: " + b.subject + "\r\n\r\n" + b.msg)
}

func PollURL() {
    for true {  // Run until the process is manually stopped
        resp, err := http.Get(url)
        if err != nil {
            panic(err)    
        }
        if resp.StatusCode != 200 {
            b := &Body{
                to, 
                "The url "+url+" is down", 
                url+" is responding with a "+strconv.Itoa(resp.StatusCode)+" error",
            }
            auth := smtp.PlainAuth("", from, pass, mailServer)
            smtp.SendMail(
                mailServer+":"+mailPort, 
                auth, 
                from,
                []string{to},
                b.toByteArray(),
            )
        }
        time.Sleep(time.Duration(sleepDuration))
    }
}

func main() {
    flag.StringVar(&url, "u", "http://gregoryrehm.com", "The url of the website to track")
    flag.IntVar(&sleepDuration, "d", 5, "The duration of time to sleep in between pinging the server")
    pass, _ = getpass.GetPassWithOptions("Password for "+from+": ", 0, getpass.DefaultMaxPass)
    flag.Parse()
    PollURL()
}
