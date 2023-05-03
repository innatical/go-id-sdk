package main

import (
	idSDK "github.com/innatical/id-sdk"
	"net/http"
	"os"
)

func main() {
	client := idSDK.New(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "http://localhost:8080/callback")

	//user, err := client.GetCurrentUser("")
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//
	//println(user.Username)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		token, err := client.GetToken(code)

		if err != nil {
			println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(token.AccessToken))
		return
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		url := client.CreateURL("identity team", "state")

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})

	println("Visit http://localhost:8080/authorize to start the flow")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
