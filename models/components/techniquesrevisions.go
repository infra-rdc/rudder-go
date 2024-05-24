package components

type TechniquesRevisions struct {
	// Revision ID
	Revision string `json:"revision"`
	// Commit's date
	Date string `json:"date"`
	// Commit's author
	Author string `json:"author"`
	// Commit's message
	Message string `json:"message"`
}

func (o *TechniquesRevisions) GetRevision() string {
	if o == nil {
		return ""
	}
	return o.Revision
}

func (o *TechniquesRevisions) GetDate() string {
	if o == nil {
		return ""
	}
	return o.Date
}

func (o *TechniquesRevisions) GetAuthor() string {
	if o == nil {
		return ""
	}
	return o.Author
}

func (o *TechniquesRevisions) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}
