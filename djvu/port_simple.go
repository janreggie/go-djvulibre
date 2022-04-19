package djvu

// SimplePort provides basic functionality for the Port interface.
// A SimplePort is automatically created
// when you create a File or a Document without specifying a Port.
// This simple port can retrieve data for local urls
// (i.e. urls referring to local files)
// and display error messages on STDERR.
// All other notifications are ignored.
type SimplePort struct {
	// TODO: Fields
}

// Copy implements Port
func (*SimplePort) Copy() Port {
	panic("unimplemented")
}

// IdToFile implements Port
func (*SimplePort) IdToFile(source Port, id string) *File {
	panic("unimplemented")
}

// IdToUrl implements Port
func (*SimplePort) IdToUrl(source Port, id string) Url {
	panic("unimplemented")
}

// NotifyChunkDone implements Port
func (*SimplePort) NotifyChunkDone(source Port, name string) {
	panic("unimplemented")
}

// NotifyDecodeProgress implements Port
func (*SimplePort) NotifyDecodeProgress(source Port, done float64) {
	panic("unimplemented")
}

// NotifyDocFlagsChanged implements Port
func (*SimplePort) NotifyDocFlagsChanged(source *Document, setMask uint64, clearMask uint64) {
	panic("unimplemented")
}

// NotifyFileFlagsChanged implements Port
func (*SimplePort) NotifyFileFlagsChanged(source *File, setMask uint64, clearMask uint64) {
	panic("unimplemented")
}

// NotifyRedisplay implements Port
func (*SimplePort) NotifyRedisplay(source *Image) {
	panic("unimplemented")
}

// NotifyRelayout implements Port
func (*SimplePort) NotifyRelayout(source *Image) {
	panic("unimplemented")
}

/// Returns 1 if #class_name# is #"DjVuPort"# or #"DjVuSimplePort"#.
func (p *SimplePort) Inherits(className string) bool {
	panic("unimplemented")
}

/** If #url# is local, it created a \Ref{DataPool}, connects it to the
  file with the given name and returns.  Otherwise returns #0#. */
func (p *SimplePort) RequestData(source Port, url *Url) *DataPool {
	panic("unimplemented")
}

/// Displays error on #stderr#. Always returns 1.
func (p *SimplePort) NotifyError(source Port, msg string) bool {
	panic("unimplemented")
}

/// Displays status on #stderr#. Always returns 1.
func (p *SimplePort) NotifyStatus(source Port, msg string) bool {
	panic("unimplemented")
}

// Enforces interface in the compiler level
var _ Port = &SimplePort{}
