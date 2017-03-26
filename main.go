package main

import (
	"github.com/dharnnie/linktor/handlers"
	"github.com/dharnnie/linktor/profile"
	"github.com/dharnnie/linktor/payment"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	serve()
}

func serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	http.HandleFunc("/bootstrap-3.2.0-dist/", handlers.ServeResource)
	http.HandleFunc("/css/", handlers.ServeResource)
	http.HandleFunc("/dist/", handlers.ServeResource)
	http.HandleFunc("/font-awesome-4.5.0/", handlers.ServeResource)
	http.HandleFunc("/images/", handlers.ServeResource)
	http.HandleFunc("/js/", handlers.ServeResource)
	http.HandleFunc("/less/", handlers.ServeResource)
	http.HandleFunc("/pages/", handlers.ServeResource)
	http.HandleFunc("/vendor/", handlers.ServeResource)
	http.HandleFunc("/start/", handlers.ServeResource)
	http.HandleFunc("/imgs/", handlers.ServeImages)

	myMux := mux.NewRouter()

	myMux.HandleFunc("/", handlers.Index)
	myMux.HandleFunc("/sign-up", handlers.SignUpServlet)
	myMux.HandleFunc("/confirm", handlers.ConfirmUser)
	myMux.HandleFunc("/login", handlers.LoginServlet)
	myMux.HandleFunc("/logout", handlers.LogoutServlet)
	myMux.HandleFunc("/profile/view", profile.ViewProfileServlet)
	myMux.HandleFunc("/profile/edit", profile.EditProfileServlet)
	myMux.HandleFunc("/profile/update", profile.UpdateServlet)
	myMux.HandleFunc("/profile/pic/update", profile.UpdatePic)
	myMux.HandleFunc("/tutor/request", handlers.RequestTutorServlet)
	myMux.HandleFunc("/earn", handlers.BecomeATutorServlet)
	myMux.HandleFunc("/me/payment", handlers.PaymentServlet)
	myMux.HandleFunc("/me", profile.ViewProfileServlet)
	myMux.HandleFunc("/me/createwallet", payment.NewWallet)
	//myMux.HandleFunc("/paymentfiles", handlers.PaymentFiles)
	myMux.HandleFunc("/me", profile.ViewProfileServlet)
	myMux.HandleFunc("/requests", handlers.NotificationServlet)	

	http.Handle("/", myMux)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}
