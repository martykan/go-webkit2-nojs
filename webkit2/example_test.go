package webkit2_test

import (
	"fmt"
	"runtime"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sourcegraph/go-webkit2/webkit2"
)

func Example() {
	runtime.LockOSThread()
	gtk.Init(nil)

	webView := webkit2.NewWebView()
	defer webView.Destroy()

	webView.Connect("load-failed", func() {
		fmt.Println("Load failed.")
	})
	webView.Connect("load-changed", func(_ *glib.Object, i int) {
		loadEvent := webkit2.LoadEvent(i)
		switch loadEvent {
		case webkit2.LoadFinished:
			fmt.Println("Load finished.")
			fmt.Printf("Title: %q\n", webView.Title())
			fmt.Printf("URI: %s\n", webView.URI())
		}
	})

	glib.IdleAdd(func() bool {
		webView.LoadURI("https://status.github.com/")
		return false
	})

	gtk.Main()

	// output:
	// Load finished.
	// Title: "GitHub System Status"
	// URI: https://status.github.com/
	// Hostname (from JavaScript): "status.github.com"
}
