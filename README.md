# head-first-design-pattern-go
Implement head-first-design-pattern using golang.

Develop env: mac

Some cases in the tutorial depend on specific Java packages. In order to implement with Golang, I have made some adjustments. Trying to restore the effects in the tutorial while following the aim pattern. The changes are listed below:


### chapter-11-proxy
virtual-proxy: implements an album cover viewer. Imageproxy controls network communication and display a "Loading" label during the process. [fyne](https://github.com/fyne-io/fyne) for UI rendering.

<img src="doc/album.png" alt="alt text" width="300" height="350">

### chapter-12-compound
dj: implements the beats control and view display with mvc pattern. The model utilizes [beep](https://github.com/gopxl/beep) for playing audio. The View utilizes [fyne](https://github.com/fyne-io/fyne) for UI rendering.

<img src="doc/djview.png" alt="alt text" width="300" height="350">