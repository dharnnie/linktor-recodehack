package payment


import(
	"fmt"
	"net/http"
	"github.com/dharnnie/linktor/sess"
	"log"
	"html/template"
)

type ThisUser struct{
	Nick string
	Done
}
type Done struct{
	Success string
}

func PaymentPage(w http.ResponseWriter, r *http.Request){
	if sess.SessionExists(w,r){
		sess.InitSessionValues(w,r)
		n := sess.GetSessionNick(w,r)
		fmt.Println("SessionExists at ViewProfile ", n) // here....
		gtd := Done{""}
		Person := ThisUser{n,gtd}
		Person.servePaymentPage(w) // change this to get payment
		fmt.Println("serveGetTutor called!!")
	}else{
		t, err := template.ParseFiles("templates/homepage.html")
		smplErr(err, "Error at Index Servlet")
		t.Execute(w, nil)	
		fmt.Println("Session Linktor does not exist")
	}
}

func PaymentFilesPage(w http.ResponseWriter, r *http.Request) {
	if sess.SessionExists(w,r){
		sess.InitSessionValues(w,r)
		n := sess.GetSessionNick(w,r)
		fmt.Println("SessionExists at ViewProfile ", n) // here....
		gtd := Done{""}
		Person := ThisUser{n,gtd}
		Person.servePaymentFilesPage(w) // change this to get payment
		fmt.Println("serveGetTutor called!!")
	}else{
		t, err := template.ParseFiles("templates/homepage.html")
		smplErr(err, "Error at Index Servlet")
		t.Execute(w, nil)	
		fmt.Println("Session Linktor does not exist")
	}
}

func (info ThisUser) servePaymentFilesPage(w http.ResponseWriter) {
	t, err := template.ParseFiles("templates/p/upload_payment_files.html")
	smplErr(err, "Could not parse upload_payment_files.html")
	t.Execute(w, info)
	fmt.Println("servePaymentPage called")
}

func (info ThisUser) servePaymentPage(w http.ResponseWriter) {
	t, err := template.ParseFiles("templates/p/payment.html")
	smplErr(err, "Could not parse payment.html")
	t.Execute(w, info)
	fmt.Println("servePaymentPage called")
}

func smplErr(e error, m string){
	if e != nil{
		log.Println(m, e)
	}
}