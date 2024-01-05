package models

import (
	"time"

	"gorm.io/gorm"
)

type Chg struct {
	gorm.Model
	Nome           string     `json:"nome"`
	Rdm            string     `json:"rdm"`
	DataChg        time.Time  `json:"dataChg"`
	Firewalls      []Firewall `gorm:"foreignKey:ChgID"`
	NumeroDoTicket string     `json:"numeroDoTicket"`
}
