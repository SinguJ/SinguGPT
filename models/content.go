package models

type Content interface {
    ToString() string
}

type Message struct {
    Msg string
}

func (m *Message) ToString() string {
    return m.Msg
}
