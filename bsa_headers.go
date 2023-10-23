package go_bsa

const BSAHeaderSize int32 = 0x24

type BSAHeaders struct {
	field                 uint32
	version               uint32
	offset                uint32
	archiveFlags          ArchiveFlags
	folderCount           uint32
	fileCount             uint32
	totalFolderNameLength uint32
	totalFileNameLength   uint32
	fileFlags             FileFlags
}
