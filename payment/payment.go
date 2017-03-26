package payment


import(
	"fmt"
	"net/http"
	"github.com/dharnnie/linktor/sess"
	"github.com/dharnnie/linktor/db"
	"log"
	"html/template"
	"net/url"
	 "bytes"
	 "strconv"
	 "io"
	 "os"
	// "encoding/json"
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

func Auth0(w http.ResponseWriter, r *http.Request) {
	apiUrl := "https://pwcstaging.herokuapp.com"
    resource := "/oauth/token"
    data := url.Values{}
    data.Set("client_id", "58d6ab494ec0a21000a915f3")
    data.Add("client_secret", "1wB8VcAP5dKkmqzipJumlP6ym9wTCeMPUWeF4MvM7rBt5MBKgsYEgsqEEfLsWbpx27i9hXmXA6LPYjg0jPZmhRLUgOddLaeSjlwW")
    data.Set("grant_type", "client_credentials")

    u, _ := url.ParseRequestURI(apiUrl)
    u.Path = resource
    urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/"

    client := &http.Client{}
    res, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
    res.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
    res.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    res.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    resp, _ := client.Do(res)
    fmt.Println(resp.Status)
    io.Copy(os.Stdout, resp.Body)


   // fmt.Println(resp.Body)
}

func RequestAccount(w http.ResponseWriter, r *http.Request) {
	apiUrl := "https://pwcstaging.herokuapp.com"
    resource := "/account/validation"
    data := url.Values{}
    data.Set("bankcode", "58d6ab494ec0a21000a915f3")
    data.Add("accountnumber", "1wB8VcAP5dKkmqzipJumlP6ym9wTCeMPUWeF4MvM7rBt5MBKgsYEgsqEEfLsWbpx27i9hXmXA6LPYjg0jPZmhRLUgOddLaeSjlwW")

    u, _ := url.ParseRequestURI(apiUrl)
    u.Path = resource
    urlStr := fmt.Sprintf("%v", u)

    client := &http.Client{}
    res, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
    
    res.Header.Add("Authorization", "auth_token= bearer")
    res.Header.Add("Content-Type", "application/json")
    res.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    resp, _ := client.Do(res)
    fmt.Println(resp.Status)
    io.Copy(os.Stdout, resp.Body)
}

// func CreateAccount(w http.ResponseWriter, r *http.Request){

// }


// 58d6ab494ec0a21000a915f3
// 1wB8VcAP5dKkmqzipJumlP6ym9wTCeMPUWeF4MvM7rBt5MBKgsYEgsqEEfLsWbpx27i9hXmXA6LPYjg0jPZmhRLUgOddLaeSjlwW