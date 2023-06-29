package alexaapi

//
//
// Interface: VideoApp

const (
	DirectiveTypeVideoAppLaunch DirectiveType = "VideoApp.Launch"
)

type VideoItem struct {
	Source   string             `json:"source"`
	Metadata *VideoItemMetadata `json:"metadata,omitempty"`
}

type VideoItemMetadata struct {
	Title    string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
}

func CreateDirectiveVideoAppLaunch(
	streamUrl string,
	title string,
	subtitle string,
) *DirectiveVideoAppLaunch {
	return &DirectiveVideoAppLaunch{
		Type: DirectiveTypeVideoAppLaunch,
		VideoItem: &VideoItem{
			Source: streamUrl,
			Metadata: &VideoItemMetadata{
				Title:    title,
				Subtitle: subtitle,
			},
		},
	}
}

//
//
// Directive: VideoApp.Launch

type DirectiveVideoAppLaunch struct {
	Type      DirectiveType `json:"type"`
	VideoItem *VideoItem    `json:"videoItem"`
}
