package alexaapi

//
//
// Interface: AudioPlayer

const (
	RequestTypeAudioPlayerPlaybackStarted        RequestType = "AudioPlayer.PlaybackStarted"
	RequestTypeAudioPlayerPlaybackFinished       RequestType = "AudioPlayer.PlaybackFinished"
	RequestTypeAudioPlayerPlaybackStopped        RequestType = "AudioPlayer.PlaybackStopped"
	RequestTypeAudioPlayerPlaybackNearlyFinished RequestType = "AudioPlayer.PlaybackNearlyFinished"
	RequestTypeAudioPlayerPlaybackFailed         RequestType = "AudioPlayer.PlaybackFailed"
)

type AudioPlayerContext struct {
	PlayerActivity       AudioPlayerActivity `json:"playerActivity"`
	OffsetInMilliseconds int64               `json:"offsetInMilliseconds,omitempty"`
	Token                string              `json:"token,omitempty"`
}

type AudioPlayerActivity string

const (
	AudioPlayerActivityUnspecified    = ""
	AudioPlayerActivityIdle           = "IDLE"
	AudioPlayerActivityPlaying        = "PLAYING"
	AudioPlayerActivityPaused         = "PAUSED"
	AudioPlayerActivityFinished       = "FINISHED"
	AudioPlayerActivityBufferUnderrun = "BUFFER_UNDERRUN"
)

type AudioPlayerPlayBehavior string

const (
	AudioPlayerPlayBehaviorUnspecified     AudioPlayerPlayBehavior = ""
	AudioPlayerPlayBehaviorReplaceAll      AudioPlayerPlayBehavior = "REPLACE_ALL"
	AudioPlayerPlayBehaviorEnqueue         AudioPlayerPlayBehavior = "ENQUEUE"
	AudioPlayerPlayBehaviorReplaceEnqueued AudioPlayerPlayBehavior = "REPLACE_ENQUEUED"
)

type AudioPlayerClearQueueBehavior string

const (
	AudioPlayerClearQueueBehaviorUnspecified   AudioPlayerClearQueueBehavior = ""
	AudioPlayerClearQueueBehaviorClearEnqueued AudioPlayerClearQueueBehavior = "CLEAR_ENQUEUED"
	AudioPlayerClearQueueBehaviorClearAll      AudioPlayerClearQueueBehavior = "CLEAR_ALL"
)

type AudioItemMetadata struct {
	Title           string        `json:"title,omitempty"`
	Subtitle        string        `json:"subtitle,omitempty"`
	Art             *DisplayImage `json:"art,omitempty"`
	BackgroundImage *DisplayImage `json:"backgroundImage,omitempty"`
}

type AudioItemCaptionDataType string

const (
	AudioItemCaptionDataTypeUnspecified AudioItemCaptionDataType = ""
	AudioItemCaptionDataTypeWebvtt      AudioItemCaptionDataType = "WEBVTT"
)

type AudioItemCaptionData struct {
	Type    AudioItemCaptionDataType `json:"type,omitempty"`
	Content string                   `json:"content,omitempty"`
}

type AudioItemStream struct {
	Url                   string                `json:"url"`
	Token                 string                `json:"token"`
	ExpectedPreviousToken string                `json:"expectedPreviousToken,omitempty"`
	OffsetInMilliseconds  uint64                `json:"offsetInMilliseconds"`
	CaptionData           *AudioItemCaptionData `json:"captionData,omitempty"`
}

type AudioPlayerAudioItem struct {
	Stream   AudioItemStream    `json:"stream"`
	Metadata *AudioItemMetadata `json:"metadata,omitempty"`
}

//
//
//

const (
	DirectiveTypeAudioPlayerPlay       DirectiveType = "AudioPlayer.Play"
	DirectiveTypeAudioPlayerStop       DirectiveType = "AudioPlayer.Stop"
	DirectiveTypeAudioPlayerClearQueue DirectiveType = "AudioPlayer.ClearQueue"
)

//
//
// Directive: AudioPlayer.Play

type DirectiveAudioPlayerPlay struct {
	Type         DirectiveType           `json:"type"`
	PlayBehavior AudioPlayerPlayBehavior `json:"playBehavior,omitempty"`
	AudioItem    *AudioPlayerAudioItem   `json:"audioItem,omitempty"`
}

func CreateDirectiveAudioPlayerPlay(
	behavior AudioPlayerPlayBehavior,
	streamUrl string,
	token string,
	prevToken *string,
	offsetMs uint64,
) *DirectiveAudioPlayerPlay {
	streamObj := AudioItemStream{
		Url:                  streamUrl,
		Token:                token,
		OffsetInMilliseconds: offsetMs,
	}
	if prevToken != nil {
		streamObj.ExpectedPreviousToken = *prevToken
	}
	return &DirectiveAudioPlayerPlay{
		Type:         DirectiveTypeAudioPlayerPlay,
		PlayBehavior: behavior,
		AudioItem: &AudioPlayerAudioItem{
			Stream: streamObj,
		},
	}
}

//
//
// Directive: AudioPlayer.Stop

type DirectiveAudioPlayerStop struct {
	Type DirectiveType `json:"type"`
}

func CreateDirectiveAudioPlayerStop() *DirectiveAudioPlayerStop {
	return &DirectiveAudioPlayerStop{
		Type: DirectiveTypeAudioPlayerStop,
	}
}

//
//
// Directive: AudioPlayer.ClearQueue

type DirectiveAudioPlayerClearQueue struct {
	Type          DirectiveType                 `json:"type"`
	ClearBehavior AudioPlayerClearQueueBehavior `json:"clearBehavior,omitempty"`
}

func CreateDirectiveAudioPlayerClearQueue(
	clearBehavior AudioPlayerClearQueueBehavior,
) *DirectiveAudioPlayerClearQueue {
	return &DirectiveAudioPlayerClearQueue{
		Type:          DirectiveTypeAudioPlayerClearQueue,
		ClearBehavior: clearBehavior,
	}
}
