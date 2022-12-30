package charts

import "strings"

const lineSmoothProfile = `{
  "xAxis": {
    "type": "category",
    "data": $CATE,
	$ROTATE
  },
  "yAxis": {
    "type": "value"
  },
  "series": [
    {
      "data": $DATA,
      "type": "line",
      "smooth": true
    }
  ]
}`

func genLineSmoothProfile(cate []string, data []int, isRotate bool) string {
	var afterRotateProfile string
	if isRotate {
		afterRotateProfile = strings.Replace(lineSmoothProfile, "$ROTATE", ROTATE, 1)
	} else {
		afterRotateProfile = strings.Replace(lineSmoothProfile, "$ROTATE", "", 1)
	}
	cateTemp := strListTemplate(cate)
	dataTemp := intListTemplate(data)
	return strings.Replace(
		strings.Replace(afterRotateProfile, "$CATE", cateTemp, 1),
		"$DATA", dataTemp, 1)
}
