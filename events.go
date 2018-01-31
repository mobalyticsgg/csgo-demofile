package demofile

// Event represents interface to execute events like generics :3
type Event interface {
	Execute(data *Data)
}

type TickDoneEvent struct {
}

func (e *TickDoneEvent) Execute(data *Data) {

}

type MatchStartedEvent struct {
}

func (e *MatchStartedEvent) Execute(data *Data) {

}

type RoundAnnounceMatchStartedEvent struct {
}

func (e *RoundAnnounceMatchStartedEvent) Execute(data *Data) {

}

type RoundEndedEvent struct {
}

func (e *RoundEndedEvent) Execute(data *Data) {

}

type RoundOfficialyEndedEvent struct {
}

func (e *RoundOfficialyEndedEvent) Execute(data *Data) {

}

type RoundMVPEvent struct {
}

func (e *RoundMVPEvent) Execute(data *Data) {

}

type RoundStartedEvent struct {
}

func (e *RoundStartedEvent) Execute(data *Data) {

}

type WinPanelMatchEvent struct {
}

func (e *WinPanelMatchEvent) Execute(data *Data) {

}

type FinalRoundEvent struct {
}

func (e *FinalRoundEvent) Execute(data *Data) {

}

type LastRoundHalfEvent struct {
}

func (e *LastRoundHalfEvent) Execute(data *Data) {

}

type FreezetimeEndedEvent struct {
}

func (e *FreezetimeEndedEvent) Execute(data *Data) {

}

type PlayerFootstepEvent struct {
}

func (e *PlayerFootstepEvent) Execute(data *Data) {

}

type PlayerTeamChangeEvent struct {
}

func (e *PlayerTeamChangeEvent) Execute(data *Data) {

}

type PlayerJumpEvent struct {
}

func (e *PlayerJumpEvent) Execute(data *Data) {

}

type PlayerKilledEvent struct {
}

func (e *PlayerKilledEvent) Execute(data *Data) {

}

type BotTakenOverEvent struct {
}

func (e *BotTakenOverEvent) Execute(data *Data) {

}

type WeaponFiredEvent struct {
}

func (e *WeaponFiredEvent) Execute(data *Data) {

}

type HeExplodedEvent struct {
}

func (e *HeExplodedEvent) Execute(data *Data) {

}

type FlashExplodedEvent struct {
}

func (e *FlashExplodedEvent) Execute(data *Data) {

}

type DecoyStartEvent struct {
}

func (e *DecoyStartEvent) Execute(data *Data) {

}

type DecoyEndEvent struct {
}

func (e *DecoyEndEvent) Execute(data *Data) {

}

type SmokeStartEvent struct {
}

func (e *SmokeStartEvent) Execute(data *Data) {

}

type SmokeEndEvent struct {
}

func (e *SmokeEndEvent) Execute(data *Data) {

}

type FireNadeStartEvent struct {
}

func (e *FireNadeStartEvent) Execute(data *Data) {

}

type FireNadeEndEvent struct {
}

func (e *FireNadeEndEvent) Execute(data *Data) {

}

type PlayerFlashedEvent struct {
}

func (e *PlayerFlashedEvent) Execute(data *Data) {

}

type BombBeginPlant struct {
}

func (e *BombBeginPlant) Execute(data *Data) {

}

type BombAbortPlant struct {
}

func (e *BombAbortPlant) Execute(data *Data) {

}

type BombPlantedEvent struct {
}

func (e *BombPlantedEvent) Execute(data *Data) {

}

type BombDefusedEvent struct {
}

func (e *BombDefusedEvent) Execute(data *Data) {

}

type BombExplodedEvent struct {
}

func (e *BombExplodedEvent) Execute(data *Data) {

}

type BombBeginDefuseEvent struct {
}

func (e *BombBeginDefuseEvent) Execute(data *Data) {

}

type BombAbortDefuseEvent struct {
}

func (e *BombAbortDefuseEvent) Execute(data *Data) {

}

type PlayerHurtEvent struct {
}

func (e *PlayerHurtEvent) Execute(data *Data) {

}

type PlayerBindEvent struct {
}

func (e *PlayerBindEvent) Execute(data *Data) {

}

type PlayerDisconnectEvent struct {
}

func (e *PlayerDisconnectEvent) Execute(data *Data) {

}

type SayTextEvent struct {
}

func (e *SayTextEvent) Execute(data *Data) {

}

type SayText2Event struct {
}

func (e *SayText2Event) Execute(data *Data) {

}

type RankUpdateEvent struct {
}

func (e *RankUpdateEvent) Execute(data *Data) {

}

type ItemEquipEvent struct {
}

func (e *ItemEquipEvent) Execute(data *Data) {

}

type ItemPickupEvent struct {
}

func (e *ItemPickupEvent) Execute(data *Data) {

}

type ItemDropEvent struct {
}

func (e *ItemDropEvent) Execute(data *Data) {

}
