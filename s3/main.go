package main

import (
	"s3/app"
)

func main() {
	// consider multi AZ S3 servers
	// consider creating creating small token based auth + ACL
	app.Init()
}
