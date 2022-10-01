# Shotstack Golang SDK

Golang SDK for [Shotstack](http://shotstack.io), the cloud video editing API.

Shotstack is a cloud based video editing platform that enables the editing of videos using code.

The platform uses an API and a JSON format for specifying how videos should be edited and what assets and titles should be used.

A server based render farm takes care of rendering the videos allowing multiple videos to be created simultaneously.

## Contents <!-- omit in toc -->

- [Using the Golang SDK](#using-the-golang-sdk)
  - [Installation](#installation)
  - [Video Editing](#video-editing)
    - [Video Editing Example](#video-editing-example)
    - [Status Check Example](#status-check-example)
  - [Video Editing Schemas](#video-editing-schemas)
    - [Edit](#edit)
	- [Timeline](#timeline)
	- [SoundTrack](#soundTrack)
	- [Font](#font)
	- [Track](#track)
	- [Clip](#clip)
	- [VideoAsset](#videoAsset)
	- [ImageAsset](#imageAsset)
	- [TitleAsset](#titleAsset)
	- [HtmlAsset](#htmlAsset)
	- [AudioAsset](#audioAsset)
	- [LumaAsset](#lumaAsset)
	- [Transition](#transition)
	- [Offset](#offset)
	- [Crop](#crop)
	- [Transform](#transform)
	- [RotateTransform](#rotatetransform)
	- [SkewTransform](#skewtransform)


# Using the Golang SDK

### Installation

#### required go v1.18+

```bash
go install github.com/harshmangalam/shotstack-sdk-golang@latest
```

## Video Editing

The Shotstack SDK enables programmatic video editing via the Edit API `render` endpoint. Add required schema using declarative api provided by library and render video.

### Video Editing Example

The example below trims the start of a video clip and plays it for 8 seconds. The edit is prepared using the SDK models
and then sent to the API for rendering.

```go
package main

import (
	"fmt"

	shotstack "github.com/harshmangalam/shotstack-sdk-golang"
	"github.com/harshmangalam/shotstack-sdk-golang/edit"
)

func main() {

	// create new configuration by adding apikey and env
	config := shotstack.
		NewConfig().
		SetApiKey("STOCKSTACK_API_KEY").
		SetEnv(shotstack.Stage)

	// create video asset
	videoAsset := edit.
		NewVideoAsset().
		SetSrc("https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/skater.hd.mp4").
		SetTrim(3)

	// create clips slice
	var clips []edit.Clip

	// create video clip
	videoClip := edit.NewClip().
		SetStart(0).
		SetLength(8).
		SetAsset(videoAsset)

	clips = append(clips, *videoClip)

	// create tracks slice
	var tracks []edit.Track

	// create track
	track := edit.NewTrack().SetClips(&clips)
	tracks = append(tracks, *track)

	// create output
	output := edit.NewOutput().
		SetFormat(edit.Mp4).
		SetResolution(edit.ResolutionSd)

	// create timeline

	timeline := edit.NewTimeline().SetTracks(&tracks)
	// post render

	res, err := edit.NewEdit().
		SetOutput(output).
		SetTimeline(timeline).
		PostRender(config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Message)
	fmt.Println(res.Response.Id)
	fmt.Println(res.Response.Message)
	fmt.Println(res.Success)

}

```

### Status Check Example

The example request below can be called a few seconds after the render above is posted. It will return the status of
the render, which can take several seconds to process.

```go
package main

import (
	"encoding/json"
	"fmt"

	shotstack "github.com/harshmangalam/shotstack-sdk-golang"
	"github.com/harshmangalam/shotstack-sdk-golang/edit"
)

func main() {

	// create new configuration by adding apikey and env
	config := shotstack.
		NewConfig().
		SetApiKey("SHOTSTACK_API_KEY").
		SetEnv(shotstack.Stage)

	id := "bcffc816-71eb-437f-a368-ec7aa9e2cc08"
	res, err := edit.GetRender(id, config)
	if err != nil {
		fmt.Println(err)
	}

  // print json response
	data, _ := json.MarshalIndent(res, "", "   ")
	fmt.Println(string(data))

	// you can access all render response data here 

	// res.Message
	// res.Response.Created
	// res.Response.Data.Output.AspectRatio
	// res.Response.Id
	// res.Response.Owner
	// ....

}

```





## Video Editing Schemas

The following schemas are used to prepare a video edit.

### Edit

An **Edit** defines the arrangement of a video on a timeline, an audio edit or an image design and the output format.

#### Example:

```go
	edit.
		NewEdit().
		SetTimeline(timeline).
		SetOutput(output).
		SetMerges(merges).
		SetCallback("https://my-server.com/edit/callback").
		SetDisk(edit.Local).
		PostRender(config)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---:
NewEdit() | initialize new edit api return *edit.Edit. | Y 
SetTimeline([*edit.Timeline](#timeline)) | A timeline represents the contents of a video edit over time, an audio edit over time, in seconds, or an image layout. A timeline consists of layers called tracks. Tracks are composed of titles, images, audio, html or video segments referred to as clips which are placed along the track at specific starting point and lasting for a specific amount of time. | -
SetOutput([*edit.Output](#output)) | The output format, render range and type of media to generate. | Y
SetMerges([*[]edit.Merge](#mergefield)) | An array of key/value pairs that provides an easy way to create templates with placeholders. The placeholders can be used to find and replace keys with values. For example you can search for the placeholder `{{NAME}}` and replace it with the value `Jane`. | -
SetCallback(string) | An optional webhook callback URL used to receive status notifications when a render completes or fails. See [webhooks](https://shotstack.io/docs/guide/architecting-an-application/webhooks/) for  more details. | -
SetDisk(edit.Disk) | The disk type to use for storing footage and assets for each render. See [disk types](https://shotstack.io/docs/guide/architecting-an-application/disk-types/) for more details. [default to `edit.Local`] <ul><li>`edit.Local` - optimized for high speed rendering with up to 512MB storage</li><li>`edit.Mount` - optimized for larger file sizes and longer videos with 5GB for source footage and 512MB for output render</li></ul> | -
PostRender(*Config) | Pass configuration containig apiKey and environment return *edit.QueuedResponse | Y

-----


### Timeline

A **Timeline** represents the contents of a video edit over time, an audio edit over time, in seconds, or an image layout. A timeline consists of layers called tracks. Tracks are composed of titles, images, audio, html or video segments referred to as clips which are placed along the track at specific starting point and lasting for a specific amount of time.

#### Example:

```go
timeline := edit.
		NewTimeline().
		SetSoundTrack(soundTrack).
		SetBackground("#000000").
		SetFonts(fonts).
		SetTracks(tracks).
		SetCache(true)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewTimeline() | Create new timeline return *edit.Timeline. | Y
SetSoundTrack([*edit.SoundTrack](#soundtrack)) | A music or audio soundtrack file in mp3 format. | -
SetBackground(string) | A hexadecimal value for the timeline background colour. Defaults to `#000000` (black). | -
SetFonts([*[]edit.Font](#font)) | An array of custom fonts to be downloaded for use by the HTML assets. | -
SetTracks([*[]edit.Track[]](#track)) | A timeline consists of an array of tracks, each track containing clips. Tracks are layered on top of each other in the same order they are added to the array with the top most track layered over the top of those below it. Ensure that a track containing titles is the top most track so that it is displayed above videos and images. | Y
SetCache(bool) | Disable the caching of ingested source footage and assets. See  [caching](https://shotstack.io/docs/guide/architecting-an-application/caching) for more details. [default to `true`] | -

---




### SoundTrack

A music or audio file in mp3 format that plays for the duration of the rendered video or the length of the audio file, which ever is shortest.

#### Example:

```go
	soundTrack := edit.
		NewSoundTrack().
		SetSrc("https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/music/disco.mp3").
		SetEffect(edit.FadeIn).
		SetVolume(1)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewSoundTrack | Create new sound track return *edit.SoundTrack. | Y
SetSrc(string) | The URL of the mp3 audio file. The URL must be publicly accessible or include credentials. | Y
SetEffect(AudioEffect) | The effect to apply to the audio file <ul><li>`FadeIn` - fade volume in only</li><li>`FadeOut` - fade volume out only</li><li>`FadeInFadeOut` - fade volume in and out</li></ul> | -
SetVolume(float32) | Set the volume for the soundtrack between 0 and 1 where 0 is muted and 1 is full volume (defaults to `1`). | -

---



### Font

Download a custom font to use with the HTML asset type, using the font name in the CSS or font tag. See our [custom fonts](https://shotstack.io/learn/html-custom-fonts/) getting started guide for more details.

#### Example:

```go
	font := edit.
		NewFont().
		SetSrc("https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/fonts/OpenSans-Regular.ttf")
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewFont() | Create new font and return *edit.Font. | Y
SetSrc(string) | The URL of the font file. The URL must be publicly accessible or include credentials. | Y

---


### Track

A track contains an array of clips. Tracks are layered on top of each other in the order in the array. The top most track will render on top of those below it.

#### Example:

```go
	track := edit.
		NewTrack().
		SetClips(clips)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewTrack() | Create new track and return *edit.Track | Y
SetClips([*[]edit.Clip](#clip)) | An array of Clips comprising of TitleClip, ImageClip or VideoClip. | Y

---



### Clip

A **Clip** is a container for a specific type of asset, i.e. a title, image, video, audio or html. You use a Clip to define when an asset will display on the timeline, how long it will play for and transitions, filters and effects to apply to it.

#### Example:

```go
	clip := edit.
		NewClip().
		SetAsset(asset).
		SetStart(2).
		SetLength(5).
		SetFit(edit.FitCrop).
		SetScale(0).
		SetPosition(edit.Center).
		SetOffset(offset).
		SetTransition(transition).
		SetEffect(edit.ZoomIn).
		SetFilter(edit.Greyscale).
		SetOpacity(1).
		SetTransform(transform)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewClip() | Create new clip and return *edit.Clip. | Y
SetAsset(any) | The type of asset to display for the duration of this Clip. Value must be one of: <ul><li>[edit.VideoAsset](#videoasset)</li><li>[edit.ImageAsset](#imageasset)</li><li>[edit.TitleAsset](#titleasset)</li><li>[edit.HtmlAsset](#htmlasset)</li><li>[edit.AudioAsset](#audioasset)</li><li>[edit.LumaAsset](#lumaasset)</li></ul>  | Y
SetStart(float32) | The start position of the Clip on the timeline, in seconds. | Y
SetLength(float32) | The length, in seconds, the Clip should play for. | Y
SetFit(ClipFit) | Set how the asset should be scaled to fit the viewport using one of the following options [default to `FitCrop`]: <ul><li>`FitCover` - stretch the asset to fill the viewport without maintaining the aspect ratio.</li><li>`FitContain` - fit the entire asset within the viewport while maintaining the original aspect ratio.</li><li>`FitCrop` - scale the asset to fill the viewport while maintaining the aspect ratio. The asset will be cropped if it exceeds the bounds of the viewport.</li><li>`FitNone` - preserves the original asset dimensions and does not apply any scaling.</li></ul>| -
SetScale(float32) | Scale the asset to a fraction of the viewport size - i.e. setting the scale to 0.5 will scale asset to half the size of the viewport. This is useful for picture-in-picture video and  scaling images such as logos and watermarks. | -
SetPosition(Position) | Place the asset in one of nine predefined positions of the viewport. This is most effective for when the asset is scaled and you want to position the element to a specific position [default to `Center`].<ul><li>`Top` - top (center)</li><li>`TopRight` - top right</li><li>`Right` - right (center)</li><li>`BottomRight` - bottom right</li><li>`Bottom` - bottom (center)</li><li>`BottomLeft` - bottom left</li><li>`Left` - left (center)</li><li>`TopLeft` - top left</li><li>`Center` - center</li></ul> | -
SetOffset([*edit.Offset](#offset)) | Offset the location of the asset relative to its position on the viewport. The offset distance is relative to the width of the viewport - for example an x offset of 0.5 will move the asset half the viewport width to the right. | -
SetTransition([*edit.Transition](#transition)) | In and out transitions for a clip - i.e. fade in and fade out | -
SetEffect(ClipEffect) | A motion effect to apply to the Clip. <ul><li>`ZoomIn` - slow zoom in</li><li>`ZoomOut` - slow zoom out</li><li>`SlideLeft` - slow slide (pan) left</li><li>`SlideRight` - slow slide (pan) right</li><li>`SlideUp` - slow slide (pan) up</li><li>`SlideDown` - slow slide (pan) down</li></ul>| -
SetFilter(ClipFilter) | A filter effect to apply to the Clip. <ul><li>`Boost` - boost contrast and saturation</li><li>`Contrast` - increase contrast</li><li>`Darken` - darken the scene</li><li>`Greyscale` - remove colour</li><li>`Lighten` - lighten the scene</li><li>`Muted` - reduce saturation and contrast</li><li>`Negative` - invert colors</li></ul> | -
SetOpacity(float32) | Sets the opacity of the Clip where 1 is opaque and 0 is transparent. [default to `1`] | -
SetTransform([*edit.Transform](#transform)) | A transformation lets you modify the visual properties of a clip. Available transformations are [edit.RotateTransformation](#rotatetransform), [edit.SkewTransformation](#skewtransformation) and [edit.FlipTransformation](#fliptransformation). Transformations can be combined to create interesting new shapes and effects. | -

---


### VideoAsset

The **VideoAsset** is used to create video sequences from video files. The src must be a publicly accessible URL to a video
resource such as an mp4 file.

#### Example:

```go
	videoAsset := edit.
		NewVideoAsset().
		SetSrc("https://shotstack-assets.s3.aws.com/mountain.mp4").
		SetTrim(5). 
		SetVolume(0.5). 
		SetCrop(crop)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewVideoAsset() | Create new video asset and return *edit.VideoAsset. | Y
SetSrc(string) | The video source URL. The URL must be publicly accessible or include credentials. | Y
SetTrim(float32) | The start trim point of the video clip, in seconds (defaults to 0). Videos will start from the in trim point. The video will play until the file ends or the Clip length is reached. | -
SetVolume(float32) | Set the volume for the video clip between 0 and 1 where 0 is muted and 1 is full volume (defaults to 0). | -
SetCrop([*edit.Crop](#crop)) | Crop the sides of an asset by a relative amount. The size of the crop is specified using a scale between 0 and 1, relative to the screen width - i.e. a left crop of 0.5 will crop half of the asset from the left, a top crop of 0.25 will crop the top by quarter of the asset. | -

---


### ImageAsset

The **ImageAsset** is used to create video from images to compose an image. The src must be a publicly accessible URL to an image resource such as a jpg or png file.

#### Example:

```go
	imageAsset := edit.
		NewImageAsset().
		SetSrc("https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/images/earth.jpg").
		SetCrop(crop)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewImageAsset() | Create new image asset and return *edit.ImageAsset. | Y
SetSrc(string) | The image source URL. The URL must be publicly accessible or include credentials. | Y
SetCrop([*edit.Crop](#crop)) | Crop the sides of an asset by a relative amount. The size of the crop is specified using a scale between 0 and 1, relative to the screen width - i.e. a left crop of 0.5 will crop half of the asset from the left, a top crop of 0.25 will crop the top by quarter of the asset. | -

---



### TitleAsset

The **TitleAsset** clip type lets you create video titles from a text string and apply styling and positioning.

#### Example:

```go
	titleAsset := edit.
		NewTitleAsset().
		SetText("My Title").
		SetStyle(edit.Minimal).
		SetColor("#ffffff").
		SetSize(edit.SizeMedium).
		SetBackground("#000000").
		SetPosition(edit.Center).
		SetOffset(offset)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewTitleAsset() | Create new title asset and return *edit.TitleAsset | Y
SetText(string) | The title text string. | Y
SetStyle(TitleAssetStyle) | Uses a preset to apply font properties and styling to the title. <ul><li>`Minimal`</li><li>`Blockbuster`</li><li>`Vogue`</li><li>`Sketchy`</li><li>`Skinny`</li><li>`Chunk`</li><li>`ChunkLight`</li><li>`Marker`</li><li>`Future`</li><li>`Subtitle`</li></ul> | -
SetColor(string) | Set the text color using hexadecimal color notation. Transparency is supported by setting the first two characters of the hex string (opposite to HTML),  i.e. #80ffffff will be white with  50% transparency [default to `#ffffff`]. | - 
SetSize(TitleSize) | Set the relative size of the text using predefined sizes from xx-small to xx-large [default to `SizeMedium`]. <ul><li>`SizeXxSmall`</li><li>`SizeXSmall`</li><li>`SizeSmall`</li><li>`SizeMedium`</li><li>`SizeLarge`</li><li>`SizeXLarge`</li><li>`SizeXxLarge`</li></ul> | -
SetBackground(string) | Apply a background color behind the text. Set the text color using hexadecimal color notation. Transparency is supported by setting the first two characters of the hex string (opposite to HTML),  i.e. #80ffffff will be white with 50% transparency. Omit to use transparent background. | -
SetPosition(Position) | Place the title in one of nine predefined positions of the viewport [default to `Center`]. <ul><li>`Top` - top (center)</li><li>`TopRight` - top right</li><li>`Right` - right (center)</li><li>`BottomRight` - bottom right</li><li>`Bottom` - bottom (center)</li><li>`BottomLeft` - bottom left</li><li>`Left` - left (center)</li><li>`TopLeft` - top left</li><li>`Center` - center</li></ul> | -
SetOffset([*edit.Offset](#offset)) | Offset the location of the title relative to its position on the screen. | -

---




### HtmlAsset

The **HtmlAsset** clip type lets you create text based layout and formatting using HTML and CSS. You can also set the height and width of a bounding box for the HTML content to sit within. Text and elements will wrap within the bounding box.

#### Example:

```go
	htmlAsset := edit.
		NewHtmlAsset().
		SetHtml("<p>Hello <b>World</b></p>").
		SetCss("p { color: #ffffff; } b { color: #ffff00; }").
		SetWidth(400).
		SetHeight(200).
		SetBackground("transparent").
		SetPosition(edit.Center)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---:
NewHtmlAsset() | Create new html asset and return [*edit.HtmlAsset](#htmlAsset) | Y
SetHtml(string) | The HTML text string. See list of [supported HTML tags](https://shotstack.io/docs/guide/architecting-an-application/html-support/#supported-html-tags). | Y
SetCss(string) | The CSS text string to apply styling to the HTML. See list of  [support CSS properties](https://shotstack.io/docs/guide/architecting-an-application/html-support/#supported-css-properties). | -
SetWidth(int) | Set the width of the HTML asset bounding box in pixels. Text will wrap to fill the bounding box. | -
SetHeight(int) | Set the height of the HTML asset bounding box in pixels. Text and elements will be masked if they exceed the  height of the bounding box. | -
SetBackground(string) | Apply a background color behind the HTML bounding box using. Set the text color using hexadecimal  color notation. Transparency is supported by setting the first two characters of the hex string  (opposite to HTML), i.e. #80ffffff will be white with 80% transparency [default to `transparent`]. | - 
SetPosition(Position) | Place the HTML in one of nine predefined positions within the HTML area [default to `Center`]. <ul><li>`Top` - top (center)</li><li>`TopRight` - top right</li><li>`Right` - right (center)</li><li>`BottomRight` - bottom right</li><li>`Bottom` - bottom (center)</li><li>`BottomLeft` - bottom left</li><li>`Left` - left (center)</li><li>`TopLeft` - top left</li><li>`Center` - center</li></ul> | -

---



### AudioAsset

The **AudioAsset** is used to add sound effects and audio at specific intervals on the timeline. The src must be a 
publicly accessible URL to an audio resource such as an mp3 file.

#### Example:

```go

	audioAsset := edit.
		NewAudioAsset().
		SetSrc("https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/music/unminus/lit.mp3").
		SetTrim(2).
		SetVolume(0.5).
		SetEffect(edit.FadeInFadeOut)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewAudioAsset() | Create new audio asset and return *edit.AudioAsset. | Y
SetSrc(string) | The audio source URL. The URL must be publicly accessible or include credentials. | Y
SetTrim(float32) | The start trim point of the audio clip, in seconds (defaults to 0). Audio will start from the trim point. The audio will play until the file ends or the Clip length is reached. | -
SetVolume(float32) | Set the volume for the audio clip between 0 and 1 where 0 is muted and 1 is full volume (defaults to 1). | -
SetEffect(AudioEffect) | The effect to apply to the audio asset: <ul><li>`FadeIn` - fade volume in only</li><li>`FadeOut` - fade volume out only</li><li>`FadeInFadeOut` - fade volume in and out</li></ul> | -

---


### LumaAsset

The **LumaAsset** is used to create luma matte masks, transitions and effects between other assets. A luma matte is a grey scale image or animated video where the black areas are transparent and the white areas solid. The luma matte animation should be provided as an mp4 video file. The src must be a publicly accessible URL to the file.

#### Example:

```go
	lumaAsset := edit.
		NewLumaAsset().
		SetSrc("https://shotstack-assets.s3-ap-southeast-2.amazonaws.com/examples/luma-mattes/paint-left.mp4").
		SetTrim(5)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewLumaAsset() | Create new luma asset and return *edit.LumaAsset | Y
SetSrc(string) | The luma matte source URL. The URL must be publicly accessible or include credentials. | Y
SetTrim(float32) | The start trim point of the luma matte clip, in seconds (defaults to 0). Videos will start from the in trim point. A luma matte video will play until the file ends or the Clip length is reached. | -

---



### Transition

The **Transition** clip type lets you define in and out transitions for a clip - i.e. fade in and fade out

#### Example:

```go
	transition := edit.
		NewTransition().
		SetIn(edit.TransitionFade).
		SetOut(edit.TransitionFade)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewTransition() | Create new transition and return *edit.Transition | Y
SetIn(TransitionType) | The transition in. Available transitions are:   <ul><li>`TransitionFade` - fade in</li><li>`TransitionReveal` - reveal from left to right</li><li>`TransitionWipeLeft` - fade across screen to the left</li><li>`TransitionWipeRight` - fade across screen to the right</li><li>`TransitionSlideLeft` - move slightly left and fade in</li><li>`TransitionSlideRight` - move slightly right and fade in</li><li>`TransitionSlideUp` - move slightly up and fade in</li><li>`TransitionSlideDown` - move slightly down and fade in</li><li>`TransitionCarouselLeft` - slide in from right to left</li><li>`TransitionCarouselRight` - slide in from left to right</li><li>`TransitionCarouselUp` - slide in from bottom to top</li><li>`TransitionCarouselDown` - slide in from top to bottom</li><li>`TransitionShuffleTopRight` - rotate in from top right</li><li>`TransitionShuffleRightTop` - rotate in from right top</li><li>`TransitionShuffleRightBottom` - rotate in from right bottom</li><li>`TransitionShuffleBottomRight` - rotate in from bottom right</li><li>`TransitionShuffleBottomLeft` - rotate in from bottom left</li><li>`TransitionShuffleLeftBottom` - rotate in from left bottom</li><li>`TransitionShuffleLeftTop` - rotate in from left top</li><li>`TransitionShuffleTopLeft` - rotate in from top left</li><li>`TransitionZoom` - fast zoom in</li></ul>  | -
SetOut(TransitionType) | The transition out. Available transitions are:  <ul><li>`TransitionFade` - fade in</li><li>`TransitionReveal` - reveal from left to right</li><li>`TransitionWipeLeft` - fade across screen to the left</li><li>`TransitionWipeRight` - fade across screen to the right</li><li>`TransitionSlideLeft` - move slightly left and fade in</li><li>`TransitionSlideRight` - move slightly right and fade in</li><li>`TransitionSlideUp` - move slightly up and fade in</li><li>`TransitionSlideDown` - move slightly down and fade in</li><li>`TransitionCarouselLeft` - slide in from right to left</li><li>`TransitionCarouselRight` - slide in from left to right</li><li>`TransitionCarouselUp` - slide in from bottom to top</li><li>`TransitionCarouselDown` - slide in from top to bottom</li><li>`TransitionShuffleTopRight` - rotate in from top right</li><li>`TransitionShuffleRightTop` - rotate in from right top</li><li>`TransitionShuffleRightBottom` - rotate in from right bottom</li><li>`TransitionShuffleBottomRight` - rotate in from bottom right</li><li>`TransitionShuffleBottomLeft` - rotate in from bottom left</li><li>`TransitionShuffleLeftBottom` - rotate in from left bottom</li><li>`TransitionShuffleLeftTop` - rotate in from left top</li><li>`TransitionShuffleTopLeft` - rotate in from top left</li><li>`TransitionZoom` - fast zoom in</li></ul>  | -

---



### Offset

Offsets the position of an asset horizontally or vertically by a relative distance.

#### Example:

```go
	offset := edit.
		NewOffset().
		SetX(0.1).
		SetY(-0.2)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewOffset() | Create new offset and return *edit.Offset. | Y
SetX(float32) | Offset an asset on the horizontal axis (left or right), range varies from -1 to 1. Positive numbers move the asset right, negative left. For all assets except titles the distance moved is relative to the width  of the viewport - i.e. an X offset of 0.5 will move the asset half the  screen width to the right. [default to `0`] | -
SetY(float32) | Offset an asset on the vertical axis (up or down), range varies from -1 to 1. Positive numbers move the asset up, negative down. For all assets except titles the distance moved is relative to the height of the viewport - i.e. an Y offset of 0.5 will move the asset up half the screen height. [default to `0`] | -

---





### Crop

Crop the sides of an asset by a relative amount. The size of the crop is specified using a scale between 0 and 1, relative to the screen width - i.e a left crop of 0.5 will crop half of the asset from the left, a top crop of 0.25 will crop the top by quarter of the asset.

#### Example:

```go
	crop := edit.
		NewCrop().
		SetTop(0.15).
		SetBottom(0.15).
		SetLeft(0).
		SetRight(0)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewCrop() | Create new crop and return *edit.Crop.  |Y
SetTop(float32) | Crop from the top of the asset | -
SetBottom(float32) | Crop from the bottom of the asset | -
SetLeft(float32) | Crop from the left of the asset | -
SetRight(float32) | Crop from the right of the asset | -

---

### Transform

Apply one or more transformations to a clip. **Transformations** alter the visual properties of a clip and can be combined to create new shapes and effects.

#### Example:

```go
	transform := edit.
		NewTransform().
		SetRotate(rotate).
		SetSkew(skew).
		SetFlip(flip)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewTransform() | Create new transformation and return *edit.Transform | Y
SetRotate([*edit.RotateTransform](#rotatetransform)) | Rotate a clip by the specified angle in degrees. Rotation origin is set based on the clips `position`. | -
SetSkew([*edit.SkewTransform](#skewtransform)) | Skew a clip so its edges are sheared at an angle. Use values between 0 and 3. Over 3 the clip will be skewed almost flat. | -
setFlip([*edit.FlipTransform](#fliptransform)) | Flip a clip vertically or horizontally. Acts as a mirror effect of the clip along the selected plane. | -

---



### RotateTransform

Rotate a clip by the specified angle in degrees. Rotation origin is set based on the clips `position`.

#### Example:

```go
	rotateTransform := edit.
		NewRotateTransform().
		SetAngle(45)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewRotateTransform() | Create new rotate transform and return *edit.RotateTransform | Y
SetAngle(int) | The angle to rotate the clip. Can be 0 to 360, or 0 to -360. Using a positive number rotates the clip clockwise, negative numbers counter-clockwise. | -

---



### SkewTransform

Skew a clip so its edges are sheared at an angle. Use values between 0 and 3. Over 3 the clip will be skewed almost flat.

#### Example:

```go
	skewTransform := edit.
		NewSkewTransform().
		SetX(0.5).
		SetY(0.5)
```

#### Methods:

Method | Description | Required
:--- | :--- | :---: 
NewSkewTransform() | Create new skew transform and return *edit.SkewTransform | Y
SetX(float32) | Skew the clip along it&#39;s x axis. [default to `0`] | -
SetY(float32) | Skew the clip along it&#39;s y axis. [default to `0`] | -

---
