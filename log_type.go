package log_watcher

type LogConfig struct {
	Stdout           bool
	File             bool
	Path             string
	MaximumLogSize   int
	MaximumLogAge    int
	MaximumLogBackup int
}
