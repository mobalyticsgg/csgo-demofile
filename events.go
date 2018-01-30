package demofile

type Event interface {
	Execute(data *Data)
}

type FireEvent struct {
}

func (f *FireEvent) Execute(data *Data) {

}

type MovementEvent struct {
}

func (m *MovementEvent) Execute(data *Data) {

}
