package edit

type TransitionType string

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

type Transition struct {
	In  TransitionType
	Out TransitionType
}

func NewTransition() *Transition {
	return new(Transition)
}

func (t *Transition) SetIn(in TransitionType) *Transition {
	t.In = in
	return t
}

func (t *Transition) SetOut(out TransitionType) *Transition {
	t.Out = out
	return t
}
