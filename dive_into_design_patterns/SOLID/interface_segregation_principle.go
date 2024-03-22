package solid

type CloudHostingProvider interface {
	createServer(region string)
	listServer(region string)
}

type CDNProvider interface {
	getCDNAddress()
}

type CloudStorageProvider interface {
	storeFile(name string)
	getFile(name string)
}

type Amazon struct {
	CloudHostingProvider
	CDNProvider
	CloudStorageProvider
}

type Dropbox struct {
	CloudStorageProvider
}

var amz Amazon
var dbx Dropbox

func sample() {
	amz.createServer("west")
	dbx.getFile("insect_1.png")
}
