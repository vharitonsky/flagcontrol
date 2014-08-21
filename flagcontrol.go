package flagcontrol

import (
	"flag"
	"log"
	"net/http"
	"text/template"
)

type flagTemplateData struct {
	Flags []*flag.Flag
}

var (
	running              = false
	port                 = flag.String("flagcontrolport", "1024", "Port to run the server on")
	displayFlagsTemplate = template.Must(template.New("displayFlagsTemplate").Parse(`
		<html>
			<body>
			<ul>
				
				{{range .Flags}}
				<li>
					{{.Name}}/{{.Value}}/{{.DefValue}}/{{.Usage}}
				</li>
				{{end}}
				
			</ul>
			</body>
		</html>
	`))
)

func Server(w http.ResponseWriter, r *http.Request) {
	var flags []*flag.Flag
	flag.VisitAll(func(f *flag.Flag) {
		flags = append(flags, f)
	})
	displayFlagsTemplate.Execute(w, flagTemplateData{Flags: flags})
}
