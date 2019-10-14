package license

import (
	"fmt"
	"os"
	"testing"
	"text/template"
	"time"

	"github.com/gjrtimmer/tcltm/pkg/resource"
)

type Author struct {
	Name  string
	Email string
}

func TestMIT(t *testing.T) {
	license, _ := resource.Get("/license/BSD2-CLAUSE")
	lic := string(license)

	tmpl := template.Must(template.New("MIT").Parse(lic))

	obj := struct {
		Authors []Author
		Year    int
	}{
		Authors: []Author{
			{
				Name:  "G.J.R. Timmer",
				Email: "gjr.timmer@gmail.com",
			},
			{
				Name:  "Second User",
				Email: "second@example.com",
			},
		},
		Year: time.Now().Year(),
	}

	if err := tmpl.Execute(os.Stdout, obj); err != nil {
		fmt.Println(err)
	}
}
