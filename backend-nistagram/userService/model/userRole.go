package model

type UserRole int

const(
	RegularRole UserRole = iota
	AdministratorRole
	AgentRole
)