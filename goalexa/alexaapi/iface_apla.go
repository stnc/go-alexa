package alexaapi

//
//
// Interface: Alexa.Presentation.APLA

const (
	DirectiveTypeAplaRenderDocument DirectiveType = "Alexa.Presentation.APLA.RenderDocument"
)

//
//
// Directive: Alexa.Presentation.APLA.RenderDocument

type DirectiveAlexaPresentationAplaRenderDocument struct {
	Type        DirectiveType  `json:"type"`
	Token       string         `json:"token"`
	Document    map[string]any `json:"document"`
	Datasources map[string]any `json:"datasources,omitempty"`
}

func CreateDirectiveAplaRenderDocumentLink(
	token string,
	url string,
	datasources map[string]any,
) *DirectiveAplRenderDocument {
	return &DirectiveAplRenderDocument{
		Type:  DirectiveTypeAplaRenderDocument,
		Token: token,
		Document: map[string]any{
			"type": "Link",
			"src":  url,
		},
		Datasources: datasources,
	}
}

func CreateDirectiveAplaRenderDocument(
	token string,
	document map[string]any,
	datasources map[string]any,
) *DirectiveAplRenderDocument {
	return &DirectiveAplRenderDocument{
		Type:        DirectiveTypeAplaRenderDocument,
		Token:       token,
		Document:    document,
		Datasources: datasources,
	}
}
