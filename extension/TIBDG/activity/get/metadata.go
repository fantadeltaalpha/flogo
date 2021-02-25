package get

type Settings struct {
	ConnectionRef   string `md:"dgConnection,required"`
	TableName	string `md:"tableName,required"`
}

/*type MessageData struct {
	Expiration    int
	DeliveryDelay int
	Destination   string
	Content       string
}

type Input struct {
	Message MessageData `md:"message"`
}

func (o *Input) ToMap() map[string]interface{} {
	msg, _ := coerce.ToObject(o.Message)
	return map[string]interface{}{
		"message": msg,
	}
}

func (o *Input) FromMap(values map[string]interface{}) error {

	msg, err := coerce.ToObject(values["message"])
	if err != nil {
		return err
	}

	o.Message = MessageData{}

	o.Message.Expiration, _ = coerce.ToInt(msg["expiration"])
	o.Message.DeliveryDelay, _ = coerce.ToInt(msg["deliveryDelay"])
	o.Message.Content, _ = coerce.ToString(msg["content"])
	o.Message.Destination, _ = coerce.ToString(msg["destination"])

	return nil
}*/
