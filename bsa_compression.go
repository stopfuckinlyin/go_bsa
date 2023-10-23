package go_bsa

type CompressionStrategy uint32

const (
	Safe       uint32 = 0         // Original compression state is preserved
	Unsafe            = 1 << iota // Uncompressed files may be compressed
	Speed                         // Compression strategy favors fast compression
	Size                          // Compression strategy favors high compression
	Aggressive                    // Compression strategy that tries many different strategies until a good ratio (or no compression) is found. Very slow!
)

type CompressionOptions struct {
	strategy                  CompressionStrategy
	extensionCompressionLevel map[string]int
}
