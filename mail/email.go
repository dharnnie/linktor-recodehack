package mail 


import (
	"log"
	"net/smtp"
	// "bytes"
	// "html/template"
)

func SendVerificationMail(URL,nick string, to string){
	// msg := "Hi " + nick + ","+ "\n" +
	// 		"Thanks for signing up. Please confirm your account by \n" +
	// 		"pasting this link in a new tab of your browser: \n " + URL +
	// 		"\n\n" +
	// 		"Cheers!"

	msg := "Hi " + nick + ","+ "\n" +
	 		"Daniel just signed you up to his hackathon hack!. If you get this email, it means his app can now send emails and he has reached his first Milestone \n" +
	 		"Also, users can paste this link in a new tab of their browser to verify: \n " + URL +
	 		"\n\n" +
	 		"Cheers!"

	send(msg, to)


}


func send(body string, to string){
	from := "dcodes.daniel@gmail.com"
	secret := "	danielcodes"

	msg := "From: " + from + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject:Hey Sir! First Milestone reached!" + "\r\n\r\n" +
			body + "\r\n"

			err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, secret, "smtp.gmail.com"), from, []string{to}, []byte(msg))
			if err != nil{
				log.Printf("Error: %s ", err)
				return
			}		
			log.Println("Email sent")	  
}


















































// type TemplateData struct{
// 	Nick string
// 	URL string
// }

// // reciever of the email
// type Recipient struct{
// 	Addr []string
// }
// // sender of the email from App - Could be admin, Samuel, Daniel or any other employee
// type Sender struct{
// 	Addr string
// 	Secret string
// }

// type Request struct {
// 	from    string
// 	to      []string
// 	subject string
// 	body    string
// }
// var auth smtp.Auth

// func SendVerificationMail(uniqueURL, nick string,resAddr string){
// 	msg := "Glad to have you!"

// 	admin := Sender{"dcodes.daniel@gmail.com", "danielcodes"}
// 	rAddr := make([]string, 2)
// 	rAddr[0] = resAddr // get the address from the request in handlers and 
// 	recipient := Recipient{rAddr}

// 	auth = smtp.PlainAuth("", admin.Addr, admin.Secret, "smtp.gmail.com")
// 	tData := TemplateData{nick,uniqueURL}

// 	r := NewRequest(recipient.Addr, "We are glad to have you!", msg )
// 	templateName := "mail.html"
// 	err := r.ParseTemplate(templateName, tData)
// 	fmt.Println()
// 	if err != nil{
// 		ok, _ := r.SendEmail()
// 		fmt.Println(ok)
// 	}else{
// 		fmt.Println("Looks like email was not sent", err)
// 	}
	
// }

// func NewRequest(to []string, subject, body string) *Request {
// 	return &Request{to:to,subject: subject,body:body,}
// }

// func (r *Request) SendEmail()(bool, error){

// 	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
// 	subject := "Subject: " + r.subject + "!\n"
// 	msg := []byte(subject + mime + "\n" + r.body)
// 	addr := "smtp.gmail.com:587"
// 	eee := make([]string, 3)
// 	eee[0] = "daniel.osineye@gmail.com"

// 	if err := smtp.SendMail(addr, auth, "dcodes.daniel@gmail.com", eee ,msg); err != nil{
// 		fmt.Println(err)
// 		return false, err
// 	}
// 	return true, nil
// }

// func (r *Request) ParseTemplate(templateFilename string, data interface{})error{
// 	t, err := template.ParseFiles(templateFilename)
// 	if err != nil{
// 		return err
// 	}
// 	buf := new(bytes.Buffer)
// 	if err = t.Execute(buf, data); err != nil {
// 		return err
// 	}
// 	r.body = buf.String()
// 	return nil
// }


	