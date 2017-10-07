package hello

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	//"google.golang.org/api/iterator"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/", handler)
}

// demo struct holds information needed to run the various demo functions.
type demo struct {
	client     *storage.Client
	bucketName string
	bucket     *storage.BucketHandle

	w   io.Writer
	ctx context.Context
	// cleanUp is a list of filenames that need cleaning up at the end of the demo.
	cleanUp []string
	// failed indicates that one or more of the demo steps failed.
	failed bool
}

func (d *demo) errorf(format string, args ...interface{}) {
	d.failed = true
	fmt.Fprintln(d.w, fmt.Sprintf(format, args...))
	log.Errorf(d.ctx, format, args...)
}

// handler is the main demo entry point that calls the GCS operations.
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	//[START get_default_bucket]
	// Use `dev_appserver.py --default_gcs_bucket_name GCS_BUCKET_NAME`
	// when running locally.
	bucket, err := file.DefaultBucketName(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
	}
	//[END get_default_bucket]

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to create client: %v", err)
		return
	}
	defer client.Close()

	buf := &bytes.Buffer{}
	d := &demo{
		w:          buf,
		ctx:        ctx,
		client:     client,
		bucket:     client.Bucket(bucket),
		bucketName: bucket,
	}

	path := r.URL.Path
	host := lookupDomainDirectory(r.URL.Host, r.URL)
	if appengine.IsDevAppServer() {
		host = os.Getenv("DEV_SERVER_DOMAIN")
	}

	if strings.HasSuffix(path, "/") {
		path += "index.html"
	}

	log.Infof(ctx, "Host: %s, %s", host, path)
	d.readFile(fmt.Sprintf("%s%s", host, path))
	extension := ".html"
	file_parts := strings.Split(path, ".")
	if len(file_parts) > 0 {
		extension = file_parts[len(file_parts)-1]
	}
	mime_type := mime.TypeByExtension(fmt.Sprintf(".%s", extension))
	log.Infof(ctx, "extension: %s, mime_type: %s", extension, mime_type)
	w.Header().Set("Content-Type", mime_type)
	if d.failed {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	buf.WriteTo(w)
}

func lookupDomainDirectory(domain string, request_url *url.URL) string {
	return_dir := domain
	dom := strings.ToUpper(domain)
	dom = strings.Replace(dom, "-", "", -1)
	dom = strings.Replace(dom, ".", "_", -1)

	env_dir := os.Getenv(dom)
	if env_dir == "" {
		env_dir = request_url.Query().Get("domain")
	}
	if env_dir != "" {
		return_dir = env_dir
	}

	return return_dir
}

//[START read]
// readFile reads the named file in Google Cloud Storage.
func (d *demo) readFile(fileName string) {

	rc, err := d.bucket.Object(fileName).NewReader(d.ctx)
	if err != nil {
		d.errorf("readFile: unable to open file from bucket %q, file %q: %v", d.bucketName, fileName, err)
		return
	}
	defer rc.Close()
	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		d.errorf("readFile: unable to read data from bucket %q, file %q: %v", d.bucketName, fileName, err)
		return
	}

	fmt.Fprintf(d.w, "%s\n", slurp)
}

//[END read]
