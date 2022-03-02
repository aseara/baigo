package instrumenter

// Instrumenter 代码增强器接口
type Instrumenter interface {
	Instrument(string) ([]byte, error)
}
