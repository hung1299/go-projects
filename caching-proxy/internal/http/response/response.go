package response

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Copy(c *gin.Context, resp *http.Response) {
	c.Status(resp.StatusCode)

	for k, vs := range resp.Header {
		for _, v := range vs {
			c.Header(k, v)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	c.Writer.Write(body)
}

func ToResponse(bs []byte) *http.Response {
	bf := bytes.NewBuffer(bs)

	resp, err := http.ReadResponse(bufio.NewReader(bf), nil)

	if err != nil {
		fmt.Println("error when write bytes to response")
		return nil
	}

	return resp
}

func ToBytes(r *http.Response) ([]byte, []byte) {
	rb := bytes.Buffer{}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading response body", err)
		return nil, nil
	}
	(*r).Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	err = r.Write(&rb)
	if err != nil {
		fmt.Println("error when write response to bytes", err)
		return nil, nil
	}

	return rb.Bytes(), bodyBytes
}
