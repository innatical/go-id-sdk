package main

import (
	"github.com/innatical/id-sdk"
	"net/http"
	"os"
)

func main() {
	idSDK.New(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "http://localhost:8080/callback")

	idSDK.Client.SetIDURL("http://localhost:4020")
	idSDK.Client.SetIDServerURL("http://localhost:6969")

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		token, err := idSDK.GetToken(code)

		if err != nil {
			println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(token.AccessToken))
		return
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		url := idSDK.CreateURL("identity team", "state")

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})

	println("Visit http://localhost:8080/authorize to start the flow")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
