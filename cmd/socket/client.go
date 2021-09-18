package socket

type client struct {
	connectionType string
	host           string
	port           string
}

type Client interface {
	Run()
}
