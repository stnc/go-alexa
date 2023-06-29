package alexaapi

//
//
// Interface: Display

type DisplayImage struct {
	ContentDescription string               `json:"contentDescription,omitempty"`
	Sources            []DisplayImageSource `json:"sources,omitempty"`
}

type DisplayImageSource struct {
	Url          string `json:"url"`
	Size         string `json:"size,omitempty"`
	WidthPixels  uint32 `json:"widthPixels,omitempty"`
	HeightPixels uint32 `json:"heightPixels,omitempty"`
}
