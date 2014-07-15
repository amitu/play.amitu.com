import sys
import urllib, urllib2

BASE = "http://localhost:8877"

def dopost(path, data):
	req = urllib2.Request(
		url = BASE + path, 
		data = data,
		headers = {
			"User-Agent": "represent"
		}
	)
	resp = urllib2.urlopen(req)
	print resp.info()
	print resp.read()

def compile(filename):
	params = {
		"version": 2,
		"body": file(filename).read()
	}
	dopost("/compile", urllib.urlencode(params))

def share(filename):
	dopost("/share", file(filename).read())

def main():
	if sys.argv[1] == "compile":
		compile(sys.argv[2])
	elif sys.argv[1] == "share":
		share(sys.argv[2])
	else:
		print "share or compile filename"

if __name__ == "__main__":
	main()