package event

func (Event) RegistryComponent() string {
	return "event.event"
}

func (Declare) RegistryComponent() string {
	return "event.declare"
}

func (Filter) RegistryComponent() string {
	return "event.filter"
}

func (Handler) RegistryComponent() string {
	return "event.handler"
}
