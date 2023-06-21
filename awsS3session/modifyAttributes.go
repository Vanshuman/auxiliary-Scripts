package awsS3session

func modifyAttributes(attr *Attributes, deleteAttr string) {
	delete(attr.Documents, deleteAttr)
}

// ModifyAttributesJson It writes modified Attributes object to a designation JSON file
func ModifyAttributesJson(attributes *Attributes, deleteAttr string) {
	modifyAttributes(attributes, deleteAttr)
}
