# See the OWNERS docs: https://git.k8s.io/community/contributors/guide/owners.md
approvers:
{{- range .owners}}
{{- if ."approver" in .roles}}
- {{ .alias }}
{{- end}}
{{- end}}
reviewers:
{{- range .owners}}
{{- if "reviewer" in .roles }}
- {{ .alias }}
{{- end}}
{{- end}}