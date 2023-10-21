package flags

import "flag"

var (
	Path string
	Task string
)

func GetFlag() {
	taskPtr := flag.String("task", "server", "server")
	path := flag.String("config", ".\\config\\config.yaml", "define config file path")
	flag.Parse()

	Path = *path
	Task = *taskPtr
}
