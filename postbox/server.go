package postbox

import (
	"fmt"
	"github.com/boltdb/bolt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"time"
)

func ServeHttp(addr string) error {
	http.HandleFunc("/", handler)
	log.Infof("Server listening to %s", addr)
	return http.ListenAndServe(addr, nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(writer, "")

	rawRequest, err := httputil.DumpRequest(request, true)
	if err != nil {
		log.Errorf("failed dumping request: %v", err)
	}

	boltClient.Open()
	defer boltClient.Close()

	err = boltClient.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("postbox"))
		id, err := bucket.NextSequence()
		if err != nil {
			log.Errorf("failed getting bucket next sequence: %v", err)
			return err
		}
		val, err := marshalLetter(&Letter{
			ID: id,
			Value: string(rawRequest),
			Ipaddr: request.RemoteAddr,
			Time: time.Now().Unix(),
		})
		if err != nil {
			log.Errorf("failed when marshalling letter: %v", err)
			return err
		}
		bucket.Put(itob(id), val)
		return nil
	})

	if err != nil {
		log.Errorf("failed updating DB: %v", err)
	}
}
