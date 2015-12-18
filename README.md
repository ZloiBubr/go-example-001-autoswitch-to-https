Go is awesome language and Beego is cool
========================================

Welcome to Siarhei Hladkou examples. 

**Features:**

- this example will help you to easy build up auto routing from `http` protocol to `https` protocol
- i used couple of articles and read beego documentation to made this sample (beego documentation is still buggy :) )
- project has been configured with IntelliJ IDEA 15 Community edition


SETUP
------------
1. download copy of files to your #GOPATH/src/go-example-001-autoswitch-to-https folder
2. generate your personal https certificate files xert.pem and key.pem and replace in root folder
```
	$ go run /usr/local/go/src/crypto/tls/generate_cert.go -host="127.0.0.1"
```
3. restore packages via
```
    $ go get github.com/astaxie/beego

    $ go get github.com/smartystreets/goconvey/convey

    $ go get github.com/jtolds/gls
```
4. packages should appear at your $GOPATH/src/github.com folder
3. build and run
4. open <http://127.0.0.1:8080> page in a browser
5. follow the "Go to Login page: Login Page Link"
6. if everythis is successed you will appear at Secure HTTPS page
