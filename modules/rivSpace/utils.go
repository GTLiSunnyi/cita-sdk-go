package rivSpace

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/GTLiSunnyi/cita-sdk-go/types"
)

func SendRivSpaceRequest(header types.GrpcRequestHeader, params map[string]interface{}, address string) ([]byte, error) {
	b1, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, address, bytes.NewReader(b1))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Authorization", header.XAuthorization)
	req.Header.Set("app-user-code", header.AppUserCode)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
