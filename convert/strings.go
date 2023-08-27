package convert

import "strconv"

func ValueOf(value any) string {
	return value.(string)
}

func FormatBool(boolean bool) string {
	if boolean {
		return "true"
	}
	return "false"
}

func FormatInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatFloat(float float64) string {
	return strconv.FormatFloat(float, 'f', -1, 64)
}
