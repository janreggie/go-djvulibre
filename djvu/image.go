package djvu

/** Main DjVu Image data structure.  This class defines the internal
  representation of a DjVu image.  This representation consists of a few
  pointers referencing the various components of the DjVu image.  These
  components are created and populated by the decoding function.  The
  rendering functions then can use the available components to compute a
  pixel representation of the desired segment of the DjVu image. */
type Image struct{}
