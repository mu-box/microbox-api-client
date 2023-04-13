
package client

import (
	"errors"
	"fmt"
	"regexp"
)

// HandleError takes an error returned from an API call, break it down and
// return important information regarding the error. The Microbox API returns
// custom errors in some instances that need to have very specific handlers.
func HandleError(err error) error {

	// if it's a micrbbox-api.Error we have special things we want to do...
	if e, ok := err.(APIError); ok {

		//
		switch e.Code {

		// Unauthorized, Forbidden, Not Found, Internal Server Error, Bad Gateway
		case 401, 403, 404, 500, 502:
			return errors.New(e.Body)

		// Unprocessable Entity -
		case 422:

			// separate the custom 422 error from the message (ex. {"upgrade":["Cannot exceed free limit"]})
			subMatch := regexp.MustCompile(`^\{\s*\"(.*)\"\s*\:\s*\[\s*\"(.*)\"\s*\]\s*\}$`).FindStringSubmatch(e.Body)
			if subMatch == nil {
				panic(e.Body)
			}

			return errors.New(fmt.Sprintf("[utils/api] %d %v - %v", 422, subMatch[1], subMatch[2]))

		//
		default:
			return errors.New(fmt.Sprintf("[utils/api] Unhandled API error - %v", err))
		}

		// ...if not, just write to the log
	} else {
		return errors.New(fmt.Sprintf("[utils/api] Unhandled error - %v", err))
	}
}
