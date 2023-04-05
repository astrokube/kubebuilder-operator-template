# See the OWNERS docs: https://git.k8s.io/community/contributors/guide/owners.md

approvers:
{{range $i, $team := .github.teams.approvers}}
- {{.team}}
{{end}}



reviewers:
{{range $i, $team := .github.teams.reviewers}}
- {{.team}}
{{end}}