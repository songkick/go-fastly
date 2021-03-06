package fastly

import (
	"sync"
	"testing"
)

// testClient is the test client.
var testClient = DefaultClient()

// testServiceID is the ID of the testing service.
var testServiceID = "7i6HN3TK9wS159v2gPAZ8A"

// testVersionLock is a lock around version creation because the Fastly API
// kinda dies on concurrent requests to create a version.
var testVersionLock sync.Mutex

// testVersion is a new, blank version suitable for testing.
func testVersion(t *testing.T) *Version {
	testVersionLock.Lock()
	defer testVersionLock.Unlock()

	v, err := testClient.CreateVersion(&CreateVersionInput{
		Service: testServiceID,
	})
	if err != nil {
		t.Fatal(err)
	}
	return v
}
