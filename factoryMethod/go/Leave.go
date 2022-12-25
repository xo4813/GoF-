package factoryMethod

type leaveDocument struct {
	document
}

func newLeaveDoc() iDocument {
	return &leaveDocument{
		document: document{}}
}
