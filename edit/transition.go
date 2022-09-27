package edit

type TransitionIn string

const (
	TransitionFade               TransitionIn = "fade"
	TransitionReveal             TransitionIn = "fade"
	TransitionWipeLeft           TransitionIn = "wipeLeft"
	TransitionWipeRight          TransitionIn = "wipeRight"
	TransitionSlideLeft          TransitionIn = "slideLeft"
	TransitionSlideRight         TransitionIn = "slideRight"
	TransitionSlideUp            TransitionIn = "slideUp"
	TransitionSlideDown          TransitionIn = "slideDown"
	TransitionCarouselLeft       TransitionIn = "carouselLeft"
	TransitionCarouselRight      TransitionIn = "carouselRight"
	TransitionCarouselUp         TransitionIn = "carouselUp"
	TransitionCarouselDown       TransitionIn = "carouselDown"
	TransitionShuffleTopRight    TransitionIn = "shuffleTopRight"
	TransitionShuffleRightBottom TransitionIn = "shuffleRightBottom"
	TransitionShuffleBottomRight TransitionIn = "shuffleBottomRight"
	TransitionShuffleBottomLeft  TransitionIn = "shuffleBottomLeft"
	TransitionShuffleRightTop    TransitionIn = "shuffleRightTop"
	TransitionShuffleLeftBottom  TransitionIn = "shuffleLeftBottom"
	TransitionShuffleLeftTop     TransitionIn = "shuffleLeftTop"
	TransitionZoom               TransitionIn = "zoom"
	TransitionShuffleTopLeft     TransitionIn = "shuffleTopLeft"
)
