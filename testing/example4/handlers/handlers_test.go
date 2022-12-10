// go test -run TestSendJSON -race -cpu 16

// Sample test to show how to test the execution of an internal endpoint.
package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"goplayground/testing/example4/handlers"
)

const succeed = "\u2713"
const failed = "\u2717"

func init() {
	handlers.Routes()
}

// TestSendJSON testing the sendjson internal endpoint.
func TestSendJSON(t *testing.T) {
	url := "/sendjson"
	statusCode := 200

	t.Log("Given the need to test the SendJSON endpoint.")
	{
		r := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)

		testID := 0
		t.Logf("\tTest %d:\tWhen checking %q for status code %d", testID, url, statusCode)
		{
			if w.Code != 200 {
				t.Fatalf("\t%s\tTest %d:\tShould receive a status code of %d for the response. Received[%d].", failed, testID, statusCode, w.Code)
			}
			t.Logf("\t%s\tTest %d:\tShould receive a status code of %d for the response.", succeed, testID, statusCode)

			var u struct {
				Name  string
				Email string
			}

			if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to decode the response.", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to decode the response.", succeed, testID)

			if u.Name == "Test" {
				t.Logf("\t%s\tTest %d:\tShould have \"Test\" for Name in the response.", succeed, testID)
			} else {
				t.Errorf("\t%s\tTest %d:\tShould have \"Test\" for Name in the response : %q", failed, testID, u.Name)
			}

			if u.Email == "user@test.com" {
				t.Logf("\t%s\tTest %d:\tShould have \"user@test.com\" for Email in the response.", succeed, testID)
			} else {
				t.Errorf("\t%s\tTest %d:\tShould have \"user@test.com\" for Email in the response : %q", failed, testID, u.Email)
			}
		}
	}
}
