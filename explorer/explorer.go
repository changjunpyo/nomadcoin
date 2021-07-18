package explorer

import (
	"fmt"
	"github.com/changjunpyo/nomadcoin/blockchain"
	"log"
	"net/http"
	"text/template"
)

type HomeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
)

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	page := HomeData{"Home", nil}
	templates.ExecuteTemplate(w, "home", page)
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.BlockChain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(port int){
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}