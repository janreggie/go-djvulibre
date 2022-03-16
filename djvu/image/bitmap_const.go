package image

const (
	_MAXRUNSIZE       = 0x3fff
	_RUNOVERFLOWVALUE = 0xc0
	_RUNMSBMASK       = 0x3f
	_RUNLSBMASK       = 0xff
)

const zerosize = 4096

var _zerobuffer [zerosize]byte
