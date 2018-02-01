package demofile

import "github.com/golang/geo/r3"

const (
	maxFilestampSize = 8
	maxOsPath        = 260
	maxSplitScreens  = 2
	maxCustomFiles   = 4
	maxPlayerName    = 128
	signedGUIDLen    = 33
)

type command int

const (
	cmdSignOn command = iota + 1
	cmdPacket
	cmdSynctick
	cmdConsole
	cmdUser
	cmdDataTables
	cmdStop
	cmdCustomData
	cmdStringTables
)

// Header represents header of demofile
type Header struct {
	Filestamp       string
	Protocol        uint32
	NetworkProtocol uint32
	ServerName      string
	ClientName      string
	MapName         string
	GameDirectory   string
	PlaybackTime    float32
	PlaybackTicks   uint32
	PlaybackFrames  uint32
	SignOnLenght    uint32
}

// OriginViewAngles represents container with view and angle
type OriginViewAngles struct {
	ViewOrigin      r3.Vector
	ViewAngles      r3.Vector
	LocalViewAngles r3.Vector
}

// SplitCmdInfo represents part of cmd info packet
type SplitCmdInfo struct {
	Flags     int32
	Original  OriginViewAngles
	Resampled OriginViewAngles
}

// CmdInfo represents cmd info packet
type CmdInfo struct {
	Parts [maxSplitScreens]SplitCmdInfo
	Data  []byte
}

// PlayerInfo represents player info
type PlayerInfo struct {
	Version         int64
	XUID            int64
	Name            [maxPlayerName]byte
	UserID          int
	GUID            [signedGUIDLen]byte
	FriendsID       uint32
	FriendsName     [maxPlayerName]byte
	FakePlayer      bool
	IsHLTV          bool
	CustomFiles     [maxCustomFiles]int32
	FilesDownloaded int8
}

type Packet struct {
	Cmd        uint8
	Tick       int32
	PlayerSlot uint8
}

type Chunk struct {
	Lenght uint32
	Data   []byte
}

// Data represents common info on moment ingame tick
type Data struct {
}
