package linkredirector

type LinkConsumerHandler interface {
	Create()
	ChangeOriginalLink()
	Delete()
}
