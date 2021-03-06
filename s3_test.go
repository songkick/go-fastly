package fastly

import "testing"

func TestClient_S3s(t *testing.T) {
	t.Parallel()

	tv := testVersion(t)

	// Create
	s3, err := testClient.CreateS3(&CreateS3Input{
		Service:         testServiceID,
		Version:         tv.Number,
		Name:            "test-s3",
		BucketName:      "bucket-name",
		AccessKey:       "access_key",
		SecretKey:       "secret_key",
		Path:            "/path",
		Period:          12,
		GzipLevel:       9,
		Format:          "format",
		TimestampFormat: "%Y",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Ensure deleted
	defer func() {
		testClient.DeleteS3(&DeleteS3Input{
			Service: testServiceID,
			Version: tv.Number,
			Name:    "test-s3",
		})

		testClient.DeleteS3(&DeleteS3Input{
			Service: testServiceID,
			Version: tv.Number,
			Name:    "new-test-s3",
		})
	}()

	if s3.Name != "test-s3" {
		t.Errorf("bad name: %q", s3.Name)
	}
	if s3.BucketName != "bucket-name" {
		t.Errorf("bad bucket_name: %q", s3.BucketName)
	}
	if s3.AccessKey != "access_key" {
		t.Errorf("bad access_key: %q", s3.AccessKey)
	}
	if s3.SecretKey != "secret_key" {
		t.Errorf("bad secret_key: %q", s3.SecretKey)
	}
	if s3.Path != "/path" {
		t.Errorf("bad path: %q", s3.Path)
	}
	if s3.Period != 12 {
		t.Errorf("bad period: %q", s3.Period)
	}
	if s3.GzipLevel != 9 {
		t.Errorf("bad gzip_level: %q", s3.GzipLevel)
	}
	if s3.Format != "format" {
		t.Errorf("bad format: %q", s3.Format)
	}
	if s3.TimestampFormat != "%Y" {
		t.Errorf("bad timestamp_format: %q", s3.TimestampFormat)
	}

	// List
	s3s, err := testClient.ListS3s(&ListS3sInput{
		Service: testServiceID,
		Version: tv.Number,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(s3s) < 1 {
		t.Errorf("bad s3s: %v", s3s)
	}

	// Get
	ns3, err := testClient.GetS3(&GetS3Input{
		Service: testServiceID,
		Version: tv.Number,
		Name:    "test-s3",
	})
	if err != nil {
		t.Fatal(err)
	}
	if s3.Name != ns3.Name {
		t.Errorf("bad name: %q", s3.Name)
	}
	if s3.BucketName != ns3.BucketName {
		t.Errorf("bad bucket_name: %q", s3.BucketName)
	}
	if s3.AccessKey != ns3.AccessKey {
		t.Errorf("bad access_key: %q", s3.AccessKey)
	}
	if s3.SecretKey != ns3.SecretKey {
		t.Errorf("bad secret_key: %q", s3.SecretKey)
	}
	if s3.Path != ns3.Path {
		t.Errorf("bad path: %q", s3.Path)
	}
	if s3.Period != ns3.Period {
		t.Errorf("bad period: %q", s3.Period)
	}
	if s3.GzipLevel != ns3.GzipLevel {
		t.Errorf("bad gzip_level: %q", s3.GzipLevel)
	}
	if s3.Format != ns3.Format {
		t.Errorf("bad format: %q", s3.Format)
	}
	if s3.TimestampFormat != ns3.TimestampFormat {
		t.Errorf("bad timestamp_format: %q", s3.TimestampFormat)
	}

	// Update
	us3, err := testClient.UpdateS3(&UpdateS3Input{
		Service: testServiceID,
		Version: tv.Number,
		Name:    "test-s3",
		NewName: "new-test-s3",
	})
	if err != nil {
		t.Fatal(err)
	}
	if us3.Name != "new-test-s3" {
		t.Errorf("bad name: %q", us3.Name)
	}

	// Delete
	if err := testClient.DeleteS3(&DeleteS3Input{
		Service: testServiceID,
		Version: tv.Number,
		Name:    "new-test-s3",
	}); err != nil {
		t.Fatal(err)
	}
}

func TestClient_ListS3s_validation(t *testing.T) {
	var err error
	_, err = testClient.ListS3s(&ListS3sInput{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.ListS3s(&ListS3sInput{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_CreateS3_validation(t *testing.T) {
	var err error
	_, err = testClient.CreateS3(&CreateS3Input{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.CreateS3(&CreateS3Input{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_GetS3_validation(t *testing.T) {
	var err error
	_, err = testClient.GetS3(&GetS3Input{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.GetS3(&GetS3Input{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.GetS3(&GetS3Input{
		Service: "foo",
		Version: "1",
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_UpdateS3_validation(t *testing.T) {
	var err error
	_, err = testClient.UpdateS3(&UpdateS3Input{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.UpdateS3(&UpdateS3Input{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	_, err = testClient.UpdateS3(&UpdateS3Input{
		Service: "foo",
		Version: "1",
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}

func TestClient_DeleteS3_validation(t *testing.T) {
	var err error
	err = testClient.DeleteS3(&DeleteS3Input{
		Service: "",
	})
	if err != ErrMissingService {
		t.Errorf("bad error: %s", err)
	}

	err = testClient.DeleteS3(&DeleteS3Input{
		Service: "foo",
		Version: "",
	})
	if err != ErrMissingVersion {
		t.Errorf("bad error: %s", err)
	}

	err = testClient.DeleteS3(&DeleteS3Input{
		Service: "foo",
		Version: "1",
		Name:    "",
	})
	if err != ErrMissingName {
		t.Errorf("bad error: %s", err)
	}
}
