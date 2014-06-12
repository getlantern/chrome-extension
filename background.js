var ws = new WebSocket("ws://127.0.0.1:9000/data");
ws.onopen = function() {
    console.log("WebSocket opened")
};
ws.onmessage = function(evt) {
    console.log("Got data", evt);
    var msg = evt.data;
    var number = parseInt(msg);
    var icon = null;
    if (number % 2 == 1) {
        icon = '32on.png';
    } else {
        icon = '32off.png';
    }
    console.log("Setting icon to:", icon);
    chrome.browserAction.setIcon({path: icon});
};
ws.onerror = function(evt) {
    console.log("error", evt);
}
ws.onclose = function() {
    // websocket is closed.
    console.log("WebSocket closed");
};