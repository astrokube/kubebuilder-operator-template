# See the OWNERS docs: https://git.k8s.io/community/contributors/guide/owners.md
approvers:
{{- range .owners}}
{{- if .coll.Has .roles "approver" }}
- {{ .alias }}
{{- end}}
{{- end}}
reviewers:
{{- range .owners}}
{{- if .roles.Has in "reviewer" }}
- {{ .alias }}
{{- end}}
{{- end}}