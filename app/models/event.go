package models

//go:generate easyjson -all
type Event struct {
	ID        uint          `json:"-" gorm:"primary_key"`
	EventType string        `json:"eventType" gorm:"not null"`
	CreateAt  int64         `json:"createAt" gorm:"not null"`
	IsRead    bool          `json:"isRead"`
	Uid       uint          `json:"uid" gorm:"not null"` // кому придет уведомление
	MakeUid   uint          `json:"-" gorm:"not null"`   // кто сделал действие
	MakeUsr   *User         `json:"makeUser" gorm:"-"`
	MetaData  EventMetaData `json:"metaData" gorm:"foreignkey:eid"`
}

//easyjson:json
type Events []Event

func (event *Event) TableName() string {
	return "events"
}

type EventMetaData struct {
	ID         uint   `json:"-" gorm:"primary_key"`
	Eid        uint   `json:"-"`
	Uid        uint   `json:"-"` // над кем/чем совершено действие
	Usr        *User  `json:"user,omitempty" gorm:"-"`
	Bid        uint   `json:"bid,omitempty"`
	Cid        uint   `json:"cid,omitempty"`
	Tid        uint   `json:"tid,omitempty"`
	EntityData string `json:"entityData,omitempty"`
	About      string `json:"about,omitempty"`
}

func (metaData *EventMetaData) TableName() string {
	return "event_meta_data"
}
