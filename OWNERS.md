# See the OWNERS docs: https://git.k8s.io/community/contributors/guide/owners.md
approvers:
{{- range .owners}}
{{- if .coll.Has .roles ."approver"}}
- {{ .alias }}
{{- end}}
{{- end}}
reviewers:
{{- range .owners}}
  {{- if .coll.Has .roles "reviewer"}}
- {{ .alias }}
{{- end}}
{{- end}}