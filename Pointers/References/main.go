package References

import (
	"strings"
)

func removeProfanity(message *string) {
	newStr := *message
	newStr = strings.ReplaceAll(newStr, "fubb", "****")
	newStr = strings.ReplaceAll(newStr, "shiz", "****")
	newStr = strings.ReplaceAll(newStr, "witch", "*****")
	*message = newStr
}
