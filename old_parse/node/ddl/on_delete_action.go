package ddl

//go:generate go run "golang.org/x/tools/cmd/stringer" -type OnDeleteAction on_delete_action.go
type OnDeleteAction int

const (
	OnDeleteNoAction OnDeleteAction = iota
	OnDeleteCascade
)
