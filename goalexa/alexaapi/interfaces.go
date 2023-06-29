package alexaapi

type Interface string

const (
	InterfaceUnspecified           Interface = ""
	InterfaceAlexaPresentationApl  Interface = "Alexa.Presentation.APL"
	InterfaceAlexaPresentationApla Interface = "Alexa.Presentation.APLA"
	InterfaceAlexaPresentationAplt Interface = "Alexa.Presentation.APLT"
	InterfaceAlexaPresentationHtml Interface = "Alexa.Presentation.HTML"
	InterfaceAudioPlayer           Interface = "AudioPlayer"
	InterfaceConnections           Interface = "Connections"
	InterfaceDialog                Interface = "Dialog"
	InterfacePlaybackController    Interface = "PlaybackController"
	InterfaceVideoApp              Interface = "VideoApp"
	InterfaceAlexaAuthorization    Interface = "Alexa.Authorization"
)
