# Overview

This is a baseline rserve instance allowing addition packages to be installed.  The go client is a command prompt sending messages to the rserve instance over TCP.

## Get started

Build and run the Rserve instance

<pre>
docker build . -t rserve
# See rserve.config if the default configuration needs modification. https://www.rforge.net/Rserve/doc.html#conf
./run.sh
</pre>

### Go client

<pre>
cd go_client
go run .
</pre>
