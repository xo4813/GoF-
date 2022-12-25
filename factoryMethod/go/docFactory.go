package factoryMethod

func MakeDoc(docType int) iDocument {

	switch docType {
	case 1:
		return newLeaveDoc()

	case 2:
		return newExpenseDoc()

	default:
		return nil
	}
}
