package main

import "text/template"

var markdownTemplate = template.Must(template.New("markdownTemplate").Parse(
	`
# Ebucks Dealz
## {{ .Name }}
[Product Page]({{ .URL }})

Price: {{ .Price }}

Image: {{ .Image }}

Savings: {{ .Savings }}

{{ if ne .Percentage "" -}}
Percentage off: {{ .Percentage }}
{{- end }}
	`,
))
