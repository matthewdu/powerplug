#!/usr/bin/env python
import SimpleHTTPServer
import SocketServer
import urllib2

class MyRequestHandler(SimpleHTTPServer.SimpleHTTPRequestHandler):
    def do_GET(self):
        cl_url = self.path[1:]
        contents = urllib2.urlopen(cl_url).read()
        self.wfile.write(contents)

Handler = MyRequestHandler
server = SocketServer.TCPServer(('', 2138), Handler)

server.serve_forever()