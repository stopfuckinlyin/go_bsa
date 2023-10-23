package go_bsa

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
