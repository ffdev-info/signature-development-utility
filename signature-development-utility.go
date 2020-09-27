// PRONOM signature development server.

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/exponential-decay/signature-development-utility/pkg/sigdevutil"
)

var (
	port         string
	bootsrapPort string
	https        bool
	cert         string
	key          string
)

// Initialize the variables used for the flags
func init() {
	flag.StringVar(&port, "port", "8080", "port to run the utility on")
	flag.StringVar(&bootsrapPort, "bootstrap", "8000", "port to run the utility on")
	flag.BoolVar(&https, "https", false, "run a HTTPS server")
	flag.StringVar(&cert, "cert", "localhost.crt", "certificate file")
	flag.StringVar(&key, "key", "localhost.key", "private key file")
}

// main is the primary entry point of this code-base.
func main() {
	// Parse application flags.
	flag.Parse()

	// Application consts.
	const staticDir = "./static"
	const routeStandard = "/processStandardSignature"
	const routeContainer = "/processContainerSignature"
	const rootDir = "/"

	// Dynamic application variables.
	var iAmListening = fmt.Sprintf("Good evening caller... listening on :%s..", port)
	var routePort = fmt.Sprintf(":%s", port)

	fs := http.FileServer(http.Dir(staticDir))
	http.Handle(rootDir, fs)
	http.HandleFunc(routeStandard, processStandardSignature)
	http.HandleFunc(routeContainer, processContainerSignature)
	log.Println(iAmListening)

	if !https {
		err := http.ListenAndServe(routePort, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Run as HTTPS.
	fmt.Println("Running with HTTPS...")
	log.Fatal(http.ListenAndServeTLS(routePort, cert, key, nil))
}

// setHeaders will set the HTTP response headers for downloading of a file.
func setHeaders(w http.ResponseWriter, contentLength int, signatureFileName string) {
	const userAgent = "signature-development-utility/2.0 (https://github.com/exponential-decay/signature-development-utility; by @beet_keeper"

	const headerUserAgent = "User-agent"
	const headerDisposition = "Content-Disposition"
	const headerContentType = "Content-Type"
	const headerContentLength = "Content-Length"

	const attachment = "attachment; filename=%s.xml"
	const mime = "application/xml"

	disposition := fmt.Sprintf(attachment, signatureFileName)

	w.Header().Set(headerUserAgent, userAgent)
	w.Header().Set(headerDisposition, disposition)
	w.Header().Set(headerContentType, mime)
	w.Header().Set(headerContentLength, strconv.Itoa(contentLength))
}

// processStandardSignaure processes the input from a standard signature form.
func processStandardSignature(writer http.ResponseWriter, request *http.Request) {
	signatureProcessing(writer, request, false)
}

// processStandardSignaure processes the input from a container signature form.
func processContainerSignature(writer http.ResponseWriter, request *http.Request) {
	signatureProcessing(writer, request, true)
}

// routeSignatureProcessing routes our data the right way and ensures
// it can be processed.
func signatureProcessing(writer http.ResponseWriter, request *http.Request, container bool) {
	switch request.Method {
	case "GET":
	case "POST":
		if err := request.ParseForm(); err != nil {
			log.Printf("ParseForm() err: %#v", err)
			return
		}
		log.Printf("Signature submission, thank you!")
		/* log.Printf(
			"Request from client (request.PostFrom): %#v\n", request.PostForm,
		) */
		if container {
			// Process the form values and return a container signature
			// file.
			signatureFile, fileName := processContainerForm(request.PostForm)
			setHeaders(writer, len(signatureFile), fileName)
			// Stream the content to the client.
			reader := strings.NewReader(signatureFile)
			io.Copy(writer, reader)
			return
		}
		// Process the form values and return a standard signature file.
		signatureFile, fileName := processStandardForm(request.PostForm)
		setHeaders(writer, len(signatureFile), fileName)
		// Stream the content to the client.
		reader := strings.NewReader(signatureFile)
		io.Copy(writer, reader)
		return
	default:
		const onlyPost = "Sorry, only POST methods are supported."
		fmt.Fprintf(writer, onlyPost)
	}
}

// processForm will initiate the processing of a standard signature
// file and return the signature file as a string.
func processStandardForm(form url.Values) (string, string) {
	var signatureFile sigdevutil.SignatureInterface
	signatureFile.ProcessSignature(form)
	const triggers = "triggers"
	if _, ok := form[triggers]; ok {
		// We now want to create sequences for container triggers.
		droid := signatureFile.ToDROID(true)
		return droid.String(), signatureFile.GetFileName()
	}
	droid := signatureFile.ToPHP(bootsrapPort)
	return droid, signatureFile.GetFileName()
}

// processContainerForm will initiate the processing of a container
// signature file and return the signature file as a string.
func processContainerForm(form url.Values) (string, string) {
	var signatureFile sigdevutil.ContainerSignatureInterface
	signatureFile.ProcessSignature(form)
	droid := signatureFile.ToDROIDContainer()
	return droid.String(), signatureFile.GetFileName()
}
