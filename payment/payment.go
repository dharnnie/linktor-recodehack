package payment


import(
	"fmt"
	"net/http"
	"github.com/dharnnie/linktor/sess"
	"github.com/dharnnie/linktor/db"
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

type Wallet struct{
	Nick string
	Email string
	Gender string
	Addr1 string
	Addr2 string
	DOB string
	State string
	Rel string
	BranchCode string
}
func PaymentPage(w http.ResponseWriter, r *http.Request){
	var n string
	if r.Method == "GET"{
		if sess.SessionExists(w,r){
			sess.InitSessionValues(w,r)
			n = sess.GetSessionNick(w,r)
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
}

func NewWallet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sess.InitSessionValues(w,r)
	n := sess.GetSessionNick(w,r)
	wallet := Wallet{
		Nick : n,
		Email : r.FormValue("email"),
		Gender : r.FormValue("gender"),
		Addr1 : r.FormValue("address1"),
		Addr2 : r.FormValue("address1"),
		DOB : r.FormValue("dob"),
		State : r.FormValue("state"),
		Rel : r.FormValue("religion"),
		BranchCode : r.FormValue("bcode"),
	}
	db.CreateWallet(n,wallet.Email, wallet.Gender,wallet.Addr1, wallet.Addr2, wallet.DOB, wallet.State,wallet.Rel, wallet.BranchCode)
	gtd := Done{"You now have a wallet"}
	Person := ThisUser{n,gtd}
	Person.servePaymentPage(w)
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

func Pay(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		
	}
}