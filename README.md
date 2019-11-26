# Mongo Go Sever

	Mongo Go Server is a server housing a Mongo Database with the server sided code written in Golang

## Downloading the source code

	git clone at https://sayaninspiron@bitbucket.org/qinvent/mongogolanglib.git

	or,

	click on the link https://sayaninspiron@bitbucket.org/qinvent/mongogolanglib.git to download the code

## Directory organization

	Create a directory named gocode and inside the code create three other directories src, pkg and bin 

	Inside src create a directory named mongogolib 

	Place all the files inside mongogolib

	The file organization is expected to be as:

		.../gocode/
		       pkg/
		       bin/
		       src/
		         /mongogolib/
		                pkg/
		       			bin/
		       			src/
		       			others/
		                readme.md

	set your GOPATH as ".../gocode"

	Install mongo packages with commands

		go get go.mongodb.org/mongo-driver/mongo
		go.mongodb.org/mongo-driver/bson
		go.mongodb.org/mongo-driver/mongo/options


## Running the code 

Upload the swagger file "APIDoc.yaml" (i.e. API documentation for the API interface) to your stand alone swagger editor 

Navigate to .../gocode/src/mongogolib/src in your command prompt and run the following commands:

	go build main.go
	./main.exe
	(give permission to the firewall prompt if any and interact from your swagger UI)