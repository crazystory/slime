package context

type Tracer struct {
	Id string
}

func (t Tracer) String() string {
	return t.Id
}

func NewTracer(id string) Tracer {
	return Tracer{
		Id: id,
	}
}
