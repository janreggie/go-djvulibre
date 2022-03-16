package image

type zerobuffer struct {
	buffer []byte
}

func newZerobuffer(zerosize uint32) zerobuffer {
	return zerobuffer{
		buffer: make([]byte, zerosize),
	}
}

func zeroes(required uint32) *zerobuffer {
	panic("unimplemented")
}
