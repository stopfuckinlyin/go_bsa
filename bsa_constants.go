package go_bsa

type ArchiveFlags int

const (
	NamedDirectories ArchiveFlags = 1 << iota
	NamedFiles
	Compressed
	unk1
	unk2
	unk3
	unk4
	unk5
	BStringPrefixed
	unk6
	unk7
)

type FileFlags int

const (
	Nif FileFlags = 1 << iota //meshes
	Dds                       //textures
	Xml                       //menus
	Wav                       //sounds
	Mp3                       //voices
	Txt                       //shaders
	Spt                       //trees
	Tex                       //fonts
	Ctl                       //misc?
	Lip                       //bsaopt\io\bsa.C LIP
	Fuz                       //bsaopt\io\bsa.C FUZ
	Bik                       //bsaopt\io\bsa.C BIK
	Jpg                       //bsaopt\io\bsa.C JPG
	Ogg                       //bsaopt\io\bsa.C OGG
	Pex                       //bsaopt\io\bsa.C GID / PEX
	unk16
	unk17
	unk18
	unk19
	unk20
	unk21
	unk22
	unk23
	unk24
	unk25
	unk26
	unk27
	unk28
	unk29
	unk30
	unk31
)
