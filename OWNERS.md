# See the OWNERS docs: https://git.k8s.io/community/contributors/guide/owners.md

approvers:
{{- range .github.teams.approvers}}
- {{ .name }}
{{- end}}
reviewers:
{{- range .github.teams.reviewers}}
- {{ .name }}
{{- end}}