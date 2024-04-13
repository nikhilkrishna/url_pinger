package logger

type Logger interface {
    Log(message string)
    Close()
}
