This is a proof of concept showing a Lantern extension with a popup that can
interact with a locally running backend and a page served from that locally
running backend.

The webpage lives at http://localhost:9000/

The extension can be installed through chrome://extensions using the 
"Load Unpacked Extension" button.

The server requires go.  Run it with `go run server.go`.

The server maintains a single number as its state.  Clicking the increment
button increments that number by one, at which point the following should happen
automatically:

1. Displayed number on webpage changes
2. Displayed number in tooltip changes
3. Lantern icon changes.  If the number is odd, the icon lights up.  If even, it
   becomes unlit.

Here are some of the core concepts tested in this POC:

1. Backend using websockets
2. Local page and extension popup can both communicate with the same local
   backend
3. Based on data received from the backend, the extension can update its action
   icon in the browser toolbar.
4. Interestingly, the extension can use an iframe to show content served from
   the local backend, meaning that we can display dynamically generated content
   in the popup (though in reality we'll be doing all our templating in the
   front-end anyway), but more importantly it means that we make some types of
   updates to how the popup works without having to deploy a new version of the
   extension to the web store!  That could be huge.

This demo happens to use ractive.js because Ox finds it easier to use than
angularjs, but that doesn't mean that we need to (or even should) use ractive.js
in the production extension.  If we're reusing a lot of our existing ui, it
probably makes sense to stick with angularjs.