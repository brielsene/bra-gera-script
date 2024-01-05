package models

import "gorm.io/gorm"

type Firewall struct {
	gorm.Model
	Nome      string `json:"nome"`
	ManagerIp string `json:"managerIp"`
	ChgID     uint
}
