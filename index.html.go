package events

var indexHTML = []byte(`<html>
<head>
    <title>Geonet Quake Events Forwarder</title>
</head>
<body>
    <div>
        <h1>Geonet Quake Events Forwarder</h1>
        <h2>This is completely unofficial and unaffiliated with Geonet</h2>
        
        <p>
            This page describes a free service which forwards events from the Geonet API in an evented manner.  This means you
            don't need to poll for quakes anymore, they can be delivered via the following protocols:

            <ul>
                <li>NATS via <pre>nats://quakes.nz:4222</pre> on subjects <pre>geonet.quakes.new</pre> and <pre>geonet.quakes.updated</pre></li>
                <li>WebSockets via <pre>ws://quakes.nz/events</pre></li>
                <li>MQTT via <pre>mqtt://quakes.nz:8883</pre> (COMING SOON)</li>
            </ul>
        </p>

        <p>The source code for this service is available at https://github.com/penguinpowernz/go-geonet-events</p>
    </div>

    <h2>Websockets Test</h2>
    <p>Below is an example of the websockets events running.  When a new earthquake happens you should see it appear below.</p>
    <div id="container"></div>

    <script type="text/javascript" src="http://ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
    <script type="text/javascript">
        $(function () {
            var ws;

            if (window.WebSocket === undefined) {
                $("#container").append("Your browser does not support WebSockets");
                return;
            } else {
                ws = initWS();
            }

            function initWS() {
                var socket = new WebSocket("ws://quakes.nz:80/events"),
                    container = $("#container")
                socket.onopen = function() {
                    container.append("<p>Socket is open, waiting for quakes to happen</p>");
                };
                socket.onmessage = function (e) {
                    container.append("<p><code>" + e.data + "</code></p>");
                }
                socket.onclose = function () {
                    container.append("<p>Socket closed</p>");
                }

                return socket;
            }
        });
    </script>
</body>
</html>`)