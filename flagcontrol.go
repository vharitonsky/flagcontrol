package flagcontrol

import (
	"flag"
	"net/http"
	"text/template"
)

type flagTemplateData struct {
	Flags []*flag.Flag
}

var (
	displayFlagsTemplate = template.Must(template.New("displayFlagsTemplate").Parse(`
		<html>
			<body>
			<form method="POST">
				<ul>
				{{range .Flags}}
					<li><label>
						"{{.Name}}"
						<input type="text" placeholder="{{.DefValue}}" name="{{.Name}}" value="{{.Value}}">
						{{.Usage}}
						</label>
					</li>
				{{end}}
				</ul>
				<input type="submit" value="Save">
			</form>
			</body>
		</html>
	`))
)

func Server(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		flag.VisitAll(func(f *flag.Flag) {
			flag.Set(f.Name, r.FormValue(f.Name))
		})
	}
	var templateData flagTemplateData
	flag.VisitAll(func(f *flag.Flag) {
		templateData.Flags = append(templateData.Flags, f)
	})
	displayFlagsTemplate.Execute(w, templateData)
}
