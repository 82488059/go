package filecontrol

const (
	BfxPathFileControl = "/usr/local/zshield/etc/BfxPathFileControl.txt"

	Allow  = 0
	Reject = 1

	Aram   = 0x01
	Read   = 1 << 1
	Write  = 1 << 2
	Delete = 1 << 3
	Rename = 1 << 4
	Access = 1 << 5

	NotifyPolicy = 1 //策略更新
)

//FilePolicy 文件管控策略
type FilePolicy struct {
	username uint32
	mode     uint32
	prochash [16]byte
	procpath [256]byte
	filepath [256]byte
}

//KerMsg 内核发送过来消息
type KerMsg struct {
	mode        uint32
	result      uint32
	createdTime [64]byte
	procpath    [256]byte
	filepath    [256]byte
}
