package repositories

type EmailTemplatesRepository interface {
	GetAllTemplates()
	GetTemplate(templateID string)
	AddTemplate()
}
