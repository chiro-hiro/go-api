package exports

import (
	"encoding/json"
)

//ExportAPI export a function to API with API wrapper
func ExportAPI(data interface{}, err error) (result []byte) {
	type apiResonse struct {
		Success bool
		Message string
		Data    interface{}
	}
	var response apiResonse
	if data != nil {
		response.Data = data
	}
	if err != nil {
		response.Message = err.Error()
		response.Success = false
	} else {
		response.Message = ""
		response.Success = true
	}
	result, e := json.Marshal(response)
	if e != nil {
		return []byte("{\"Success\":false,\"Message\":\"" + e.Error() + "\",\"Data\":null}")
	}
	return
}
