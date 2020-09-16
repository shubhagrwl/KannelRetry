package v1

import (
	"fmt"
	"net/http"
	"strings"
)

func ReportRetry(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	// http://localhost:9090/Report?id=12345&route_id=%i&status=%d&phone=%q&sender=%Q&message=%A&message_id=%F&ts=%T&meta=%D
	id := r.FormValue("id")
	routeID := r.FormValue("route_id")
	status := r.FormValue("status")
	phone := r.FormValue("phone")
	sender := r.FormValue("sender")
	message := r.FormValue("message")
	ts := r.FormValue("ts")
	meta := r.FormValue("meta")

	if strings.Compare(status, "delivrd") != 0 {
		//call kannel API
		//cgi-bin/sendsms?username=kannel&password=kannel&to=9953940590&text=Tests&smsc=SampleSMPP&from=mobme&dlr-mask=31&pid=64&validity=1440&dlr-url=%s
	} else {
		// store it in DB
	}

	fmt.Println(id, routeID, status, phone, sender, message, ts, meta)

}
