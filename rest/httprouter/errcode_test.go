package httprouter

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestGetError(t *testing.T) {
	resp := GetError(ErrInvalidParameterCode, errors.New("xxxxx")).SetOuterErrorMessage("custom outer message").SetRequestParameter(map[string]interface{}{
		"title": "234567890-",
	})
	b, _ := json.Marshal(resp)
	t.Log(string(b))

	resp = GetError(ErrInvalidParameterCode, errors.New("错误来啦"))
	b, _ = json.Marshal(resp)
	t.Log(string(b))
}
