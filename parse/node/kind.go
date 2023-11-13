package node

//go:generate go run "golang.org/x/tools/cmd/stringer" -type Kind kind.go
type Kind int

const (
	KindUnspecified Kind = iota

	KindIdentifier
	KindKeyword
	KindPath
	KindStructTypeField
	KindType
	KindTypeSize
)
