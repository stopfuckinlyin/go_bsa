package go_bsa

const BSAFolderSize int32 = 0x10

type BSAFolderRecord struct {
	hash   uint64
	count  uint32
	offset uint32
}
