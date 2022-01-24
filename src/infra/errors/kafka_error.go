package errors

type KafkaError struct {
	CommonError
}

func (err KafkaError) Error() string {
	return err.ClientMessage
}

func (err CommonError) ToKafkaError() KafkaError {
	return KafkaError{
		CommonError: err,
	}
}
