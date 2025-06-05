package node

//go:generate go run "golang.org/x/tools/cmd/stringer" -type Order order.go
type Order int

const (
	OrderASC Order = iota
	OrderDESC
)
