var ractive = new Ractive({
  // The `el` option can be a node, an ID, or a CSS selector.
  el: 'container',

  // We could pass in a string, but for the sake of convenience
  // we're passing the ID of the <script> tag above.
  template: '#template',

  // Here, we're passing in some initial data
  data: { number: null }
});

var ws = new WebSocket("ws://localhost:9000/data");
ws.onopen = function() {
    console.log("WebSocket opened")
};
ws.onmessage = function(evt) {
    console.log("Got data", evt);
    var msg = evt.data;
    ractive.set("number", parseInt(msg));
};
ws.onerror = function(evt) {
    console.log("error", evt);
}
ws.onclose = function() {
    // websocket is closed.
    console.log("WebSocket closed");
};

document.getElementById("increment").onclick = function() {
  console.log("Called increment");
  ws.send(1);
}