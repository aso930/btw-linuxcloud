package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

var defaultTemplate = `<!DOCTYPE html>
<html>
<head>
<title>Linux and the cloud - Web App</title>
<style>
  body{
background: #151515
}
.title{
font-size: 30px;
font-weight: normal;
font-family: Monaco, "Bitstream Vera Sans Mono", "Lucida Console", Terminal, monospace;
color: #b5e853;
letter-spacing: -0.03em;
text-shadow: 0 1px 1px rgba(0,0,0,0.1),0 0 5px rgba(181,232,83,0.1),0 0 10px rgba(181,232,83,0.1);
}
.subtitle
{
    color: #eaeaea;
    font-size: 28px;
    line-height: 1.5;
    font-family: Monaco, "Bitstream Vera Sans Mono", "Lucida Console", Terminal, monospace;
}
</style>
</head>
<body>
<p class="title">
Salut{{.Name}},
</p>
<p class="subtitle">
Aceasta aplicatie ruleaza pe {{.PrivateIP}}
</p>
</body>
</html>`

type pageData struct {
	Name      string
	PrivateIP string
}

func main() {
	var listenPort = flag.String("port", "18080", "This is the port the application will bind to.")
	flag.Parse()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", helloAnonymous).Methods("GET")
	router.HandleFunc("/{token}", hello).Methods("GET")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *listenPort), router))
}
func getIP() string {
	ifaces, err := net.Interfaces()
	var privateIPByte net.IP
	addrs, err := ifaces[2].Addrs()
	for _, addr := range addrs {
		switch v := addr.(type) {
		case *net.IPNet:
			privateIPByte = v.IP
		case *net.IPAddr:
			privateIPByte = v.IP
		}
		if privateIPByte == nil || privateIPByte.IsLoopback() {
			continue
		}
		privateIPByte = privateIPByte.To4()
		if privateIPByte == nil {
			continue
		}
		return privateIPByte.String()
	}
	if err != nil {
		fmt.Println(err)
	}
	return ""
}
func logRequest(remoteIP string, method string, path string) {
	fmt.Printf("%s call on endpoint %s received from: %s\n", method, path, remoteIP)
}
func helloAnonymous(w http.ResponseWriter, r *http.Request) {
	logRequest(r.RemoteAddr, r.Method, r.URL.RequestURI())
	var data pageData
	data.Name = ""
	data.PrivateIP = getIP()
	t := template.Must(template.New("").Parse(defaultTemplate))
	t.ExecuteTemplate(w, "", data)
}
func hello(w http.ResponseWriter, r *http.Request) {
	logRequest(r.RemoteAddr, r.Method, r.URL.RequestURI())
	vars := mux.Vars(r)
	var data pageData
	data.Name = " " + vars["token"]
	data.PrivateIP = getIP()
	t := template.Must(template.New("").Parse(defaultTemplate))
	t.ExecuteTemplate(w, "", data)

}
