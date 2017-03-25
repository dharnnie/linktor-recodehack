package handlers

import(
	"net/http"
	"html/template"
	"fmt"
	"github.com/dharnnie/linktor/db"
	"github.com/dharnnie/linktor/sess"
	"github.com/dharnnie/linktor/enc"
	"github.com/dharnnie/linktor/profile"
	"github.com/dharnnie/linktor/mail"
	"github.com/dharnnie/linktor/payment"
	"encoding/json"
)


type UserLogin struct{
	Nick string
	password string
	LoginFail string
}

type SignUpDets struct{
	Nick string
	Fname string
	Lname string
	Email string
	Bio string
	Password string
	SignUpError string
	Login UserLogin
	ImagePath string
}



var key = "123456789012345678901234"



type Payload struct{
	Email string
	Num string // either bvn or phone number
	Gender string
	Addr1 string
	Addr2 string
	DOB string 	
	State string
	Rel string
	BranchCode string

	// other files
}



type Blank struct{
	Nick string
	ImagePath string
}

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
func ServeResource(w http.ResponseWriter, r  *http.Request) {
	path := "templates/" + r.URL.Path
	http.ServeFile(w, r, path)	
}

func ServeImages(w http.ResponseWriter, r *http.Request) {
	path := "assets/" + r.URL.Path
	http.ServeFile(w,r,path)
}

func Index(w http.ResponseWriter, r *http.Request) {
	if sess.SessionExists(w,r){
		sess.InitSessionValues(w,r)
		n := sess.GetSessionNick(w,r)
		data := Blank{n, ""}
		fmt.Println("SessionExists with: ", n) // here....
		data.Dashboard(w)
	}else{
		t, err := template.ParseFiles("templates/homepage.html")
		smplErr(err, "Error at Index Servlet")
		t.Execute(w, nil)	
		fmt.Println("Session Linktor does not exist")	
	}
}

func SignUpServlet(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		t, err := template.ParseFiles("templates/start/signup.html")
		smplErr(err, "Error at SignUp Servlet")
		t.Execute(w, nil)
	}else{
		logindata := UserLogin{"", "", ""} // uses the same struct type for 
		ThisUser := SignUpDets{r.FormValue("nick"),r.FormValue("form-first-name"), r.FormValue("form-last-name"), r.FormValue("form-email"), r.FormValue("bio"), r.FormValue("password"), "", logindata, ""}
		fmt.Println(ThisUser)

		res := db.SignUpAuth(ThisUser.Nick)
		if res == "ok"{
			hashedPass := enc.Encrypt(key, ThisUser.Password)
			
			fmt.Println("Hashed pass is : ", hashedPass)
			db.SignUp(ThisUser.Nick, ThisUser.Fname, ThisUser.Lname, ThisUser.Email, ThisUser.Bio, hashedPass)
			db.InitSignUp(ThisUser.Nick)
			sess.SaveSession(w,r,ThisUser.Nick)
			sess.InitSessionValues(w,r)
			fmt.Println(sess.GetSessionNick(w,r))
			ThisUser.Dashboard(w,r)
			// hash nick and send email
			hashedNick := enc.Encrypt(key, ThisUser.Fname)
			// save hashed nick to verification table
			db.InitializeVerification(ThisUser.Nick, hashedNick, ThisUser.Email)
			verURL := "localhost:3000/confirm?v="+ hashedNick
			// send verification email
			mail.SendVerificationMail(verURL, ThisUser.Nick, ThisUser.Email)
		}else{
			// redirect to signup
			ThisUser.SignUpError = "Nick Exists - Choose Another"
			t, err := template.ParseFiles("template/start/signup.html")
			smplErr(err, "Could not parse signup at SignUp else")
			t.Execute(w, ThisUser)
		}	
	}
}

func ConfirmUser(w http.ResponseWriter, r *http.Request){
	v := r.FormValue("v")
	// unhash person, check if any email matches nick, update verification table or return error page
	Nick := enc.Decrypt(key, v)
	db.ConfirmNewUser(Nick)
}

func LoginServlet(w http.ResponseWriter, r *http.Request) {
	nLogin := UserLogin{r.FormValue("login-nick"), r.FormValue("login-password"),""}
	encedPass := enc.Encrypt(key, nLogin.password) 
	loginStat := db.LoginAuth(nLogin.Nick, encedPass)
	loginD := UserLogin{Nick: nLogin.Nick, password: "", LoginFail: ""}
	data := SignUpDets{"","","","","","","",loginD, ""}
	switch loginStat{
	case 11:
		fmt.Println("Login fail.. Username does not exist")
		BadNick(w, &data)
	case 21:
		fmt.Println("Details don't match!!")
		LoginMismatch(w, &data)
	case 20:
		fmt.Println("Details match...\n Setting session")
		sess.SaveSession(w,r,nLogin.Nick)
		sess.InitSessionValues(w,r)
		sNick := sess.GetSessionNick(w,r)
		data.Nick = sNick //nick found in session is passed to temp 
		data.DashboardHome(w,r)
	}
	// something went wrong, we could not process your
}

func LogoutServlet(w http.ResponseWriter, r *http.Request) {
	sess.DeleteSession(w,r)
	t, err := template.ParseFiles("templates/homepage.html")
	smplErr(err, "Could not parse homepage.html at LogoutServlet")
	t.Execute(w, nil)
}
func (b Blank) Dashboard(w http.ResponseWriter) {
	t, err := template.ParseFiles("templates/p/dashboard.html")
	smplErr(err, "Could not parse  dashboard.html at Blank method")
	img := profile.GetPicPath(b.Nick)
	b.ImagePath = "../imgs/" + img
	t.Execute(w, b)
}
// loads dashboard on succesful signup
func (v SignUpDets) DashboardHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/p/dashboard.html")
	// get details
	img := profile.GetPicPath(v.Nick)
	v.ImagePath = "../imgs/" + img
	smplErr(err, "Error parsing dashboard.html")
	//sess.SetLoginSession(v.Nick,w,r)
	t.Execute(w, v)
}

func (newU SignUpDets) Dashboard(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/p/dashboard.html")
	// get details
	img :=  profile.GetPicPath(newU.Nick)
	newU.ImagePath = "../imgs/" + img
	smplErr(err, "Error parsing dashboard.html")
	//sess.SetLoginSession(newU.Nick,w,r)
	t.Execute(w, newU)
}

func BadNick(w http.ResponseWriter, sud *SignUpDets) {
	t, err := template.ParseFiles("templates/start/signup.html")
	smplErr(err, "Error parsing signup.html")
	sud.Login.LoginFail = "Could not find your Nick" // modify
	badNickData := sud
	t.Execute(w, badNickData)
}

func LoginMismatch(w http.ResponseWriter, sud *SignUpDets) {
	t, err := template.ParseFiles("templates/start/signup.html")
	smplErr(err, "Could not parse signup.html")
	sud.Login.LoginFail = "Details dont match" // modify
	t.Execute(w, sud)
}

func PaymentServlet(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		payment.PaymentPage(w,r)
	}
}

func PaymentFiles(w http.ResponseWriter, r *http.Request) {
	var ac Payload
	if r.Method == "POST"{
		// save data in struct as a pointer
		r.ParseForm()
		payment.PaymentFilesPage(w,r)
		ac = Payload{
			Email : r.FormValue("email"),
			Num : r.FormValue("phone"),// either bvn or phone number
			Gender :r.FormValue("gender"),
			Addr1 : r.FormValue("address1"),
			Addr2 : r.FormValue("address2"),
			DOB : r.FormValue("dob"),
			State : r.FormValue("state"),
			Rel : r.FormValue("religion"),
			BranchCode : r.FormValue("bcode"),
		}
		fmt.Println(ac)
		SendPWCRequest(w, ac)

		// call upload files page

		// verify befoe saving
		// if success from paywithcapture, save to db

		


	}else{ // use a put method here
		// payment.PaymentSettingsDone()
	}
}

// from json  package
func SendPWCRequest(w http.ResponseWriter, x Payload){
	response, err := createPWCRequest(x) // Pay with capture
	if err != nil{
		fmt.Println("Err at SendPWCRequest")
		panic(err)
	}
	fmt.Println(string(response))
	//return string(response)
}

func createPWCRequest(x Payload) ([]byte, error	){
	p := x
	// p = Payload{
	// 	Email : x.Email,
	// 	Num  : x.Num,
	// 	Gender : x.Gender,
	// 	Addr1 : x.Addr1,
	// 	Addr2 : x.Addr2,
	// 	DOB : x.DOB,
	// 	State : x.State,
	// 	Rel : x.Rel,
	// 	BranchCode : x.BranchCode, 
	// }

	return json.MarshalIndent(p, "", "  ")
}
