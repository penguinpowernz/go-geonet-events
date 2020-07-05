package events

var version string

var indexHTML = []byte(`<html>
<head>
    <title>New Zealand Earthquake Events Forwarder</title>
    <style type="text/css">
        body { font-family:  Arial, Helvetica, sans-serif; padding: 10 200 10 200; color: white; background-color: #181a1b;}
        .code { font-family: 'Courier New', Courier, monospace; background-color: black; color: white;  padding: 3 5 3 5; border-radius: 5px;}
        ul li { padding: 10 5 10 5; }
        a { font-weight: bold;text-decoration: none;}
        a:visited { color:cyan;}
        a:hover { color: lightblue; }
    </style>
</head>
<body>
    <div>
        <h1>Geonet Quake Events Forwarder v` + version + `</h1>
        <h2>This is completely unofficial and unaffiliated with Geonet NZ</h2>
        
        <p>
            This page describes a free (beta) service which forwards events from the <a href="https://api.geonet.org.nz/">Geonet API</a> in an evented manner.  This means you
            don't need to poll the <a href="https://api.geonet.org.nz/">Geonet API</a> for quakes anymore, they can be delivered via the following protocols:

            <ul>
                <li>NATS via <span class="code">nats://quakes.nz:4222</span> on subjects <span class="code">geonet.quakes.new</span> and <span class="code">geonet.quakes.updated</span> (username of <span class="code">client</span> and empty password)</li>
                <li>WebSockets via <span class="code">ws://quakes.nz/events</span></li>
                <li>MQTT via <span class="code">mqtt://quakes.nz:8883</span> (COMING SOON)</li>
            </ul>
        </p>

        <p>
            The source code for this service is available at <a href="https://github.com/penguinpowernz/go-geonet-events">https://github.com/penguinpowernz/go-geonet-events</a>. 
            We would love and appreciate contributions to the code but would prefer if this was the only running instance of the code given that it polls Geonets API every second.
        </p>

        <p>If you find any bugs or issues with the service please report them on that repositories <a href="https://github.com/penguinpowernz/go-geonet-events/issues">issue tracker</a></p>
    </div>

    <h2>Event types</h2>

    <p>
        There are two event types; <span class="code">new</span> and <span class="code">updated</span>. The latter is only sent when there is an update to a previous
        quake that was sent out.  This is usually due to revisions to the magnitude and depth of the quakes, the <span class="code">quality</span> field can be used to determine the usefulness
        of the event.
    </p>

    <div style="overflow: hidden"></div>
        <div style="float: right; width: 45%">
            <h3>New payload</h3>
            <script src="https://gist.github.com/penguinpowernz/93ea0468d592456ae596a232bf217a4c.js"></script>
        </div>

        <div style=" width: 45%">
            <h3>Update payload</h3>
            <script src="https://gist.github.com/penguinpowernz/6c55fcdd603e3d59953d0155e01482a3.js"></script>
        </div>
    </div>

    <div>
        Example code of how to use this can be found in the projects <a href="https://github.com/penguinpowernz/go-geonet-events">README</a>.

        <h2>Websockets Test</h2>
        <p>Below is an example of the websockets events running.  When a new earthquake happens you should see it appear below.</p>
        <div id="container" class="code" style="padding: 10px; padding-left: 20px;"></div>
    </div>

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
