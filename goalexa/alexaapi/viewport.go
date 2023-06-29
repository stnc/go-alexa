package alexaapi

type Viewport struct {
	Type ViewportType `json:"type,omitempty"`

	//
	//
	// type = "" or "APL"
	Dpi                int                     `json:"dpi,omitempty"`
	Touch              []ViewportTouchStyle    `json:"touch,omitempty"`
	Keyboard           []ViewportKeyboardStyle `json:"keyboard,omitempty"`
	Video              *ViewportVideoSpec      `json:"video,omitempty"`
	Experiences        []ViewportExperience    `json:"experiences,omitempty"`
	Mode               ViewportMode            `json:"mode,omitempty"`
	Shape              ViewportShape           `json:"shape,omitempty"`
	PixelHeight        int                     `json:"pixelHeight,omitempty"`
	PixelWidth         int                     `json:"pixelWidth,omitempty"`
	CurrentPixelHeight int                     `json:"currentPixelHeight,omitempty"`
	CurrentPixelWidth  int                     `json:"currentPixelWidth,omitempty"`

	//
	//
	// type = "APLT"
	Format            ViewportFormat             `json:"format,omitempty"`
	Id                string                     `json:"id,omitempty"`
	InterSegments     []ViewportInterSegment     `json:"interSegments,omitempty"`
	LineLength        int                        `json:"lineLength,omitempty"`
	LineCount         int                        `json:"lineCount,omitempty"`
	SupportedProfiles []ViewportSupportedProfile `json:"supportedProfiles,omitempty"`
}

type ViewportSupportedProfile string

const (
	ViewportSupportedProfileUnspecified        ViewportSupportedProfile = ""
	ViewportSupportedProfileFourCharacterClock ViewportSupportedProfile = "FOUR_CHARACTER_CLOCK"
)

type ViewportInterSegment struct {
	X          int    `json:"x,omitempty"`
	Y          int    `json:"y,omitempty"`
	Characters string `json:"characters,omitempty"`
}

type ViewportFormat string

const (
	ViewportFormatUnspecified  ViewportFormat = ""
	ViewportFormatSevenSegment ViewportFormat = "SEVEN_SEGMENT"
)

type ViewportType string

const (
	ViewportTypeUnspecified ViewportType = ""
	ViewportTypeApl         ViewportType = "APL"
	ViewportTypeAplt        ViewportType = "APLT"
)

type ViewportVideoSpec struct {
	Codecs []Codec `json:"codecs,omitempty"`
}

type Codec string

const (
	CodecUnspecified Codec = ""
	CodecH_264_41    Codec = "H_264_41"
	CodecH_264_42    Codec = "H_264_42"
)

type ViewportKeyboardStyle string

const (
	ViewportKeyboardStyleUnspecified ViewportKeyboardStyle = ""
	ViewportKeyboardStyleDirection   ViewportKeyboardStyle = "DIRECTION"
)

type ViewportTouchStyle string

const (
	ViewportTouchStyleUnspecified ViewportTouchStyle = ""
	ViewportTouchStyleSingle      ViewportTouchStyle = "SINGLE"
)

type ViewportExperience struct {
	CanRotate bool `json:"canRotate,omitempty"`
	CanResize bool `json:"canResize,omitempty"`
}

type ViewportMode string

const (
	ViewportModeUnspecified ViewportMode = ""
	ViewportModeHub         ViewportMode = "HUB"
	ViewportModeTv          ViewportMode = "TV"
	ViewportModePc          ViewportMode = "PC"
	ViewportModeMobile      ViewportMode = "MOBILE"
	ViewportModeAuto        ViewportMode = "AUTO"
)

type ViewportShape string

const (
	ViewportShapeUnspecified ViewportShape = ""
	ViewportShapeRound       ViewportShape = "ROUND"
	ViewportShapeRectangle   ViewportShape = "RECTANGLE"
)
