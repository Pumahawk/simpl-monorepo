{{- define "replaceUnderscoreWithDash" -}}
  {{- regexReplaceAll "_" . "-" -}}
{{- end }}