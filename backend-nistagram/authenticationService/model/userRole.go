package model

type UserRole int

const(
	Regular UserRole = iota
	Administrator
	Agent
)