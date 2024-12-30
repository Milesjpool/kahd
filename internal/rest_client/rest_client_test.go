package rest

// func TestRestClient_Get(t *testing.T) {
// 	t.Run("Get_delegatesToHTTPClient", func(t *testing.T) {
// 		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			if r.URL.Path == "/resource" {
// 				w.WriteHeader(http.StatusOK)
// 			} else {
// 				w.WriteHeader(http.StatusNotFound)
// 			}
// 		}))
// 		defer ts.Close()

// 		client := NewClient(ts.URL)

// 		actual, err := client.Get("/resource")

// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusOK, actual.)
// 	})
// }

// func TestRestClient_GetResource(t *testing.T) {
// 	type resource struct {
// 		Resource string `json:"resource"`
// 	}

// 	//TODO: we shouldn't need a real server - just a mocked response data-stream.
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.Path == "/resource" {
// 			w.Write([]byte(`{"resource": "123"}`))
// 		} else {
// 			w.WriteHeader(http.StatusNotFound)
// 		}
// 	}))
// 	defer ts.Close()

// 	t.Run("GetResource_DecodesValidResponseBody", func(t *testing.T) {
// 		actual, err := GetResource[resource](&http.Client{}, ts.URL+"/resource")

// 		assert.NoError(t, err)
// 		assert.Equal(t, &resource{Resource: "123"}, actual)
// 	})

// 	t.Run("GetResource_ErrorsWhenNotFound", func(t *testing.T) {
// 		resourceId := ts.URL + "/unknown"
// 		_, err := GetResource[resource](&http.Client{}, resourceId)

// 		assert.ErrorIs(t, err, NotFound{resourceId})
// 	})
// }
