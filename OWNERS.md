# See the OWNERS docs: https://git.k8s.io/community/contributors/guide/owners.md
approvers:
{{range .owners}}
{{if .role.Has "approver"}}
- {{ .alias }}
{{end}}
{{end}}

reviewers:
{{range .owners}}
{{if .role.Has "reviewer"}}
- {{ .alias }}
{{end}}
{{end}}