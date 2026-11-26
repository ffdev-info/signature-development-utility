// PRONOM signature development server.

package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ffdev-info/signature-development-utility/internal/logformatter"
	"github.com/ffdev-info/signature-development-utility/pkg/sigdev"
)

var (
	port  string
	dev   bool
	https bool
	cert  string
	key   string
	debug bool

	// Application variables.
	version string
)

// Initialize the variables used for the flags
func init() {
	flag.StringVar(&port, "port", "8080", "port to run the utility on")
	flag.BoolVar(&dev, "dev", false, "use development tools, e.g. non-embedded files")
	flag.BoolVar(&https, "https", false, "run a HTTPS server")
	flag.StringVar(&cert, "cert", "localhost.crt", "certificate file")
	flag.StringVar(&key, "key", "localhost.key", "private key file")
	flag.BoolVar(&debug, "debug", false, "enable debug logging")
}

const createdBy = "created in the signature development utility, https://ffdev.info:"

// getHandlers returns a http handler depending on whether we have
// dev mode set.
func getHandlers() http.Handler {
	const staticDir = "static"
	var httpHandler http.Handler
	if !dev {
		httpHandler = http.FileServer(http.Dir(staticDir))
	} else {
		static, _ := fs.Sub(staticEmbed, staticDir)
		httpHandler = http.FileServer(http.FS(static))
	}
	return httpHandler
}

// main is the primary entry point of this code-base.
func main() {

	// Initialize all the things.
	flag.Parse()
	logformatter.Set("sigdev", true)
	if debug {
		log.Println("debug logging is turned on")
	}

	const routeStandard = "/processStandardSignature"
	const routeContainer = "/processContainerSignature"
	const rootDir = "/"

	// Dynamic application variables.
	var routePort = fmt.Sprintf(":%s", port)

	// Handlers.
	http.Handle(rootDir, getHandlers())
	http.HandleFunc(routeStandard, processStandardSignature)
	http.HandleFunc(routeContainer, processContainerSignature)

	// good night seattle we love you...
	log.Printf("good evening caller... listening on :%s...", port)
	log.Printf("connect with http://0.0.0.0:%s", port)

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
	const userAgent = "signature-development-utility/2.0 (https://github.com/ffdev-info/signature-development-utility; by @beet_keeper"

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
			log.Printf("%#v", err)
			return
		}
		log.Printf("signature submission, thank you!")
		if debug {
			log.Printf(
				"request from client (request.PostFrom): %#v\n", request.PostForm,
			)
		}

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
		signatureFile = fmt.Sprintf("<!-- %s %s -->\n%s", createdBy, version, signatureFile)
		setHeaders(writer, len(signatureFile), fileName)
		// Stream the content to the client.
		reader := strings.NewReader(signatureFile)
		io.Copy(writer, reader)
		return
	default:
		const onlyPost = "sorry, only POST methods are supported."
		fmt.Fprint(writer, onlyPost)
	}
}

// processForm will initiate the processing of a standard signature
// file and return the signature file as a string.
func processStandardForm(form url.Values) (string, string) {
	var signatureFile sigdev.SignatureInterface
	signatureFile.ProcessSignature(form)
	const triggers = "triggers"
	if _, ok := form[triggers]; ok {
		// We now want to create sequences for container triggers.
		droid := signatureFile.ToDROID(true)
		return droid.String(), signatureFile.GetFileName()
	}
	droid := signatureFile.ToDROID(false)
	return droid.String(), signatureFile.GetFileName()
}

// processContainerForm will initiate the processing of a container
// signature file and return the signature file as a string.
func processContainerForm(form url.Values) (string, string) {
	var signatureFile sigdev.ContainerSignatureInterface
	signatureFile.ProcessSignature(form)
	droid := signatureFile.ToDROIDContainer()
	return droid.String(), signatureFile.GetFileName()
}
