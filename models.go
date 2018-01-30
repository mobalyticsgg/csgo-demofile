package demofile

import "github.com/golang/geo/r3"

type command int

const (
	signOn command = iota + 1
	packet
	synctick
	console
	user
	dataTables
	stop
	customData
	stringTables
)

// Header represents header of demofile
type Header struct {
	Filestamp       string
	Protocol        int32
	NetworkProtocol int32
	ServerName      string
	ClientName      string
	MapName         string
	GameDirectory   string
	PlaybackTime    float64
	PlaybackTicks   int32
	PlaybackFrames  int32
	SignOnLenght    int32
}

// QAngle represents angle in demofile
type QAngle struct {
	Pitch float64
	Yaw   float64
	Roll  float64
}

// OriginViewAngles represents container with view and angle
type OriginViewAngles struct {
	ViewOrigin      r3.Vector
	ViewAngles      QAngle
	LocalViewAngles QAngle
}

// SplitCmdInfo represents part of cmd info packet
type SplitCmdInfo struct {
	Flags     int32
	Original  OriginViewAngles
	Resampled OriginViewAngles
}

// Data represents common info on moment ingame tick
type Data struct {
}
