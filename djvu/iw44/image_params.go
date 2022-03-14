package iw44

// EncoderParams describe encoding parameters for an IW44 image.
//
// This data structure gathers the quality specification parameters
// needed for encoding each chunk of an IW44 file.
// Chunk data is generated until meeting
// either the slice target, the size target or the decibel target.
type EncoderParams struct {
	// Slice target.
	// Data generation for the current chunk stops
	// if the total number of slices
	// (in this chunk and all the previous chunks)
	// reaches this value.
	//
	// The default value `0` has a special meaning:
	// data will be generated regardless of the number of slices in the file.
	slices uint64

	// Size target.
	// Data generation for the current chunk stops
	// if the total data size
	// (in this chunk and all the previous chunks),
	// expressed in bytes,
	// reaches this value.
	//
	// The default value `0` has a special meaning:
	// data will be generated regardless of the file size.
	bytes uint64

	// Decibel target.
	// Data generation for the current chunk stops
	// if the estimated luminance error, expressed in decibels,
	// reaches this value.
	//
	// The default value `0` has a special meaning:
	// data will be generated regardless of the estimated luminance error.
	// In fact, setting the value to `0`
	// shortcuts the computation of the estimated luminance error
	// and sensibly speeds up the encoding process.
	decibels float64
}

// CRCBMode represents the chrominance processing selector.
// Determines how the chrominance information should be processed.
type CRCBMode uint8

const (
	// The wavelet transform will discard the chrominance information
	// and only keep the luminance.
	// The image will show in shades of gray.
	CRCB_NONE CRCBMode = iota

	// The wavelet transform will process the chrominance
	// at only half the image resolution.
	// This option creates smaller files
	// but may create artifacts in highly colored images.
	CRCB_HALF

	// The wavelet transform will process the chrominance at full resolution.
	// This is the default.
	CRCB_NORMAL

	// The wavelet transform will process the chrominance at full resolution.
	// This option also disables the chrominance encoding delay (see `SetCrcbDelay`)
	// which usually reduces the bitrate associated with the chrominance information.
	CRCB_FULL
)
