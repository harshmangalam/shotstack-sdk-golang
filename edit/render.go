package edit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	shotstack "github.com/harshmangalam/shotstack-sdk-golang"
)

type ClipFilter string
type ClipEffect string
type Position string
type ClipFit string
type TransitionType string

const (
	Boost     ClipFilter = "boost"
	Contrast  ClipFilter = "contrast"
	Darken    ClipFilter = "darken"
	Greyscale ClipFilter = "greyscale"
	Lighten   ClipFilter = "lighten"
	Muted     ClipFilter = "muted"
	Negative  ClipFilter = "negative"
)

const (
	ZoomIn     ClipEffect = "zoomIn"
	ZoomOut    ClipEffect = "zoomOut"
	SlideLeft  ClipEffect = "slideLeft"
	SlideRight ClipEffect = "slideRight"
	SlideUp    ClipEffect = "slideUp"
	SlideDown  ClipEffect = "slideDown"
)

const (
	Top         Position = "top"
	TopRight    Position = "topRight"
	Right       Position = "right"
	BottomRight Position = "bottomRight"
	Bottom      Position = "bottom"
	BottomLeft  Position = "bottomLeft"
	Left        Position = "left"
	TopLeft     Position = "topLeft"
	Center      Position = "center"
)

const (
	FitCover   ClipFit = "cover"
	FitContain ClipFit = "contain"
	FitCrop    ClipFit = "crop"
	FitNone    ClipFit = "none"
)

const (
	TransitionFade               TransitionType = "fade"
	TransitionReveal             TransitionType = "fade"
	TransitionWipeLeft           TransitionType = "wipeLeft"
	TransitionWipeRight          TransitionType = "wipeRight"
	TransitionSlideLeft          TransitionType = "slideLeft"
	TransitionSlideRight         TransitionType = "slideRight"
	TransitionSlideUp            TransitionType = "slideUp"
	TransitionSlideDown          TransitionType = "slideDown"
	TransitionCarouselLeft       TransitionType = "carouselLeft"
	TransitionCarouselRight      TransitionType = "carouselRight"
	TransitionCarouselUp         TransitionType = "carouselUp"
	TransitionCarouselDown       TransitionType = "carouselDown"
	TransitionShuffleTopRight    TransitionType = "shuffleTopRight"
	TransitionShuffleRightBottom TransitionType = "shuffleRightBottom"
	TransitionShuffleBottomRight TransitionType = "shuffleBottomRight"
	TransitionShuffleBottomLeft  TransitionType = "shuffleBottomLeft"
	TransitionShuffleRightTop    TransitionType = "shuffleRightTop"
	TransitionShuffleLeftBottom  TransitionType = "shuffleLeftBottom"
	TransitionShuffleLeftTop     TransitionType = "shuffleLeftTop"
	TransitionZoom               TransitionType = "zoom"
	TransitionShuffleTopLeft     TransitionType = "shuffleTopLeft"
)

type Crop struct {
	Top    float32 `json:"top"`
	Bottom float32 `json:"bottom"`
	Left   float32 `json:"left"`
	Right  float32 `jight:"right"`
}

// edit

func (e *Edit) PostRender(config *shotstack.Config) interface{} {
	Url := fmt.Sprintf("https://api.shotstack.io/%v/render", config.Env)
	jsonData, err := json.MarshalIndent(e, "", "   ")

	fmt.Println(string(jsonData))

	if err != nil {
		panic(err)
	}

	headers := map[string][]string{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
		"x-api-key":    {config.ApiKey},
	}

	data := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("POST", Url, data)

	if err != nil {
		panic(err)
	}
	req.Header = headers

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	stringResp := string(bodyBytes)
	return stringResp

}
