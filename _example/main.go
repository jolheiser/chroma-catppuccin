package main

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/quick"
	"github.com/alecthomas/chroma/v2/styles"

	_ "go.jolheiser.com/chroma-catppuccin" // Import for Register side-effect
)

//go:embed main.go
var quine string

const style = "catppuccin"

var (
	tmplFormat = template.Must(template.New("").Parse(`{{range .}}<a href="../?format={{.}}">{{.}}</a> | {{end}}`))
	tmplStyle  = template.Must(template.New("").Parse(`{{range .}}<a href="../?style={{.}}">{{.}}</a> | {{end}}`))
)

func main() {
	http.HandleFunc("/format", func(w http.ResponseWriter, _ *http.Request) {
		tmplFormat.Execute(w, formatters.Names())
	})
	http.HandleFunc("/style", func(w http.ResponseWriter, _ *http.Request) {
		tmplStyle.Execute(w, styles.Names())
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		format := "html"
		if r.FormValue("format") != "" {
			format = r.FormValue("format")
		}
		style := style
		if r.FormValue("style") != "" {
			style = r.FormValue("style")
		}
		if err := quick.Highlight(w, quine, "go", format, style); err != nil {
			log.Println(err)
		}
	})

	go func() {
		log.Println("Listening on http://localhost:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Println(err)
		}
	}()

	var _ FooStringer = &Foo{
		Bar: `bar`,
		Baz: []int{
			1,
			1_000,
			0xFFFFFF,
		},
		Ban: false,
		Bab: map[rune]interface{}{
			'a':    (*Foo)(nil),
			'\x99': nil,
		},
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill, os.Interrupt)
	<-ch
}

// Foo is a fooer
type Foo struct {
	Bar string               `json:"bar"`
	Baz []int                `json:"baz"`
	Ban bool                 `json:"ban"`
	Bab map[rune]interface{} `json:"bab"`
}

// String implements fmt.Stringer (and also FooStringer)
func (f *Foo) String() string {
	return "lorem ipsum"
}

// FooStringer is a fmt.Stringer copy
type FooStringer interface {
	String() string
}
