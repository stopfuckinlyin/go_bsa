package go_bsa

const BSAFileSize int32 = 0x10

type BSAFileRecord struct {
	hash   uint64
	size   uint32
	offset uint32
}
