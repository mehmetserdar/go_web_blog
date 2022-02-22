package helpers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func SetAlert(w http.ResponseWriter, r *http.Request, message string) error {
	session, err := store.Get(r, "go-alert")
	if err != nil {
		fmt.Println(err)
		return err
	}

	session.AddFlash(message)

	return session.Save(r, w)

}

func GetAlert(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	session, err := store.Get(r, "go-alert")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	data := make(map[string]interface{})
	flashes := session.Flashes()

	if len(flashes) > 0 {
		data["is_alert"] = true
		data["message"] = flashes[0]
	} else {
		data["is_alert"] = false
		data["message"] = nil
	}

	session.Save(r, w)

	return data

}
