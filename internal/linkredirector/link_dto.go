package linkredirector

type LinkDTO struct {
	ID           uint   `json:"id" validate:"required,gt=0"`
	OriginalLink string `json:"originalLink" validate:"required,url"`
}
