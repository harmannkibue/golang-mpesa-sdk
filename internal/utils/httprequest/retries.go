package httprequest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	originalBody []byte
	err          error
)

// RetryDo when given a request it executes the request client with the specified redoes -.
func RetryDo(req *http.Request, maxRetries int, timeout time.Duration,
	backoffStrategy []time.Duration) (*http.Response, error) {

	if req != nil && req.Body != nil {
		originalBody, err = copyBody(req.Body)
		resetBody(req, originalBody)
	}

	if err != nil {
		return nil, err
	}

	AttemptLimit := maxRetries

	if AttemptLimit <= 0 {
		AttemptLimit = 1
	}

	client := http.Client{
		Timeout: timeout,
	}

	var resp *http.Response

	// number of retries
	for i := 1; i <= AttemptLimit; i++ {
		resp, err = client.Do(req)

		if err != nil {
			return nil, err
		}

		log.Printf("%d RETRYING STATUS CODE %d ON URL %s", i, resp.StatusCode, resp.Request.URL)

		// The status code is withing the 400-499 range and thus contains error message -.
		if err == nil && resp.StatusCode >= 400 && resp.StatusCode < 500 {
			var r interface{}

			// Checking for error 404 specifically since we know its mainly caused by the url being invalid -.
			if resp.StatusCode == 404 {
				return nil, fmt.Errorf("resource url %s not found", req.URL)
			}

			decodeErr := json.NewDecoder(resp.Body).Decode(&r)

			if decodeErr != nil {
				return nil, fmt.Errorf("error: %w", decodeErr)
			}

			rData, er := json.Marshal(r)

			if er != nil {
				return nil, er
			}

			return nil, errors.New(string(rData))

		} else if err == nil && resp.StatusCode < 400 {
			log.Printf("SUCCESSFULLY SENT REQUEST TO %s", req.URL)
			return resp, nil
		}

		// If retrying, release
		if resp != nil {
			resp.Body.Close()
		}

		// resetting body
		if req.Body != nil {
			resetBody(req, originalBody)
		}

		time.Sleep(backoffStrategy[i-1] + 1*time.Microsecond)
	}

	// Here, it means that retrying is useless -.
	return nil, fmt.Errorf("something went wrong please try again later")
}

// Copying the body so that the original body with resources can be released -.
func copyBody(src io.ReadCloser) ([]byte, error) {
	b, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, errors.New("Error reading the request body ")
	}
	src.Close()
	return b, nil
}

// Resetting in order to close the request body to avoid keeping it open all the time -.
func resetBody(request *http.Request, originalBody []byte) {
	request.Body = io.NopCloser(bytes.NewBuffer(originalBody))
	request.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewBuffer(originalBody)), nil
	}
}
