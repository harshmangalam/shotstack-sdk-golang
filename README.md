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
SetTransform([*edit.Transform](#transform)) | A transformation lets you modify the visual properties of a clip. Available transformations are [edit.RotateTransformation](#rotatetransformation), [edit.SkewTransformation](#skewtransformation) and [edit.FlipTransformation](#fliptransformation). Transformations can be combined to create interesting new shapes and effects. | -

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
		SetOffset((offset))
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
