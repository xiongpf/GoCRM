package comm
import (
    "encoding/json"
)

func JsonMarshal(resp *DwzResp) string{
    jsonStr, err := json.Marshal(resp)
    if err != nil {
        //log.Error("SwitchAction format json - %s", err.Error())
        return ""
    }

    return string(jsonStr)
}