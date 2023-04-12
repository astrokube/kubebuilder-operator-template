# See the OWNERS docs: https://git.k8s.io/community/contributors/guide/owners.md
approvers:
{{- range .owners}}
{{- if has .roles approver}}
- {{ .alias }}
{{- end}}
{{- end}}

reviewers:
{{- range .owners}}
  {{- if has .roles reviewer}}
- {{ .alias }}
{{- end}}
{{- end}}