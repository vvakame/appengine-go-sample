package sample

import (
	"appengine/aetest"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestArticleServiceGet(t *testing.T) {
	opt := &aetest.Options{AppID: "unittest", StronglyConsistentDatastore: true}

	inst, err := aetest.NewInstance(opt)
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer inst.Close()

	r, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create r: %v", err)
	}
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	if w.Code != 200 {
		b, _ := ioutil.ReadAll(w.Body)
		t.Fatalf("unexpected %d, expected 200, body=%s", w.Code, string(b))
	}
}
