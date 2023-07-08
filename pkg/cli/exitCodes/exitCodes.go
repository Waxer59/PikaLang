package exitCodes

type ExitCode int

const (
	Success ExitCode = iota
	FileNameError
	FileReadError
	GetWDError
	FileExtensionError
)
