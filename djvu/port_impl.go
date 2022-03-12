package djvu

// Concrete implementation of a Port
type port struct {
	// Insert elements here...
}

// Copy implements Port
func (*port) Copy() Port {
	panic("unimplemented")
}

// IdToFile implements Port
func (*port) IdToFile(source Port, id string) *File {
	panic("unimplemented")
}

// IdToUrl implements Port
func (*port) IdToUrl(source Port, id string) Url {
	panic("unimplemented")
}

// Inherits implements Port
func (*port) Inherits(className string) bool {
	panic("unimplemented")
}

// NotifyChunkDone implements Port
func (*port) NotifyChunkDone(source Port, name string) {
	panic("unimplemented")
}

// NotifyDecodeProgress implements Port
func (*port) NotifyDecodeProgress(source Port, done float64) {
	panic("unimplemented")
}

// NotifyDocFlagsChanged implements Port
func (*port) NotifyDocFlagsChanged(source *Document, setMask uint64, clearMask uint64) {
	panic("unimplemented")
}

// NotifyError implements Port
func (*port) NotifyError(source Port, msg string) bool {
	panic("unimplemented")
}

// NotifyFileFlagsChanged implements Port
func (*port) NotifyFileFlagsChanged(source *File, setMask uint64, clearMask uint64) {
	panic("unimplemented")
}

// NotifyRedisplay implements Port
func (*port) NotifyRedisplay(source *Image) {
	panic("unimplemented")
}

// NotifyRelayout implements Port
func (*port) NotifyRelayout(source *Image) {
	panic("unimplemented")
}

// NotifyStatus implements Port
func (*port) NotifyStatus(source Port, msg string) bool {
	panic("unimplemented")
}

// RequestData implements Port
func (*port) RequestData(source Port, url *Url) *DataPool {
	panic("unimplemented")
}

// Makes sure port implements Port in the compiler level
var _ Port = &port{}
