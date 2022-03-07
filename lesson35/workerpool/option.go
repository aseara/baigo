package workerpool

// Option pool 的可调选项
type Option func(*Pool)

// WithBlock 设置 Pool 的 block
func WithBlock(block bool) Option {
	return func(p *Pool) {
		p.block = block
	}
}

// WithPreAllocWorkers 设置 Pool 的 preAlloc
func WithPreAllocWorkers(preAlloc bool) Option {
	return func(p *Pool) {
		p.preAlloc = preAlloc
	}
}
