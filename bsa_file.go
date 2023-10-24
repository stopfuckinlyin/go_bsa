package go_bsa

type CompressionStrategy uint32

const (
	Safe       CompressionStrategy = 0         // Original compression state is preserved
	Unsafe                         = 1 << iota // Uncompressed files may be compressed
	Speed                                      // Compression strategy favors fast compression
	Size                                       // Compression strategy favors high compression
	Aggressive                                 // Compression strategy that tries many different strategies until a good ratio (or no compression) is found. Very slow!
)

type CompressionOptions struct {
	strategy                  CompressionStrategy
	extensionCompressionLevel map[string]int
}

type ArchiveSettings struct {
	defaultCompressed bool
	bStringPrefixed   bool
	options           CompressionOptions
}

const BSAFileSize int32 = 0x10

type BSAFileRecord struct {
	hash   uint64
	size   uint32
	offset uint32
}

const (
	FlagCompress      uint32 = 1 << 30
	DeflateLevelSize  int32  = 9
	DeflateLevelSpeed int32  = 1
	DeflateLevelMixed int32  = 5
)

type BSAFile struct {
	name                    string
	fileName                string
	path                    string
	isCompressed            bool
	size                    uint32
	originalSize            uint32
	fileRecord              BSAFileRecord
	hash                    uint64
	data                    []byte
	settings                ArchiveSettings
	forceCompressionChecked bool
	optDeflateLevel         uint32
}
