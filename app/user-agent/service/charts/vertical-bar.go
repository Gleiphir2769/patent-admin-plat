package charts

import "strings"

const verticalBarProfile = `{
  "tooltip": {
    "trigger": "axis",
    "axisPointer": {
      "type": "shadow"
    }
  },
  "grid": {
    "left": "3%",
    "right": "4%",
    "bottom": "3%",
    "containLabel": true
  },
  "xAxis": [
    {
      "type": "value"
    }
  ],
  "yAxis": [
    {
      "type": "category",
      "data": $CATE,
      $ROTATE
    }
  ],
  "series": [
    {
      "type": "bar",
      "stack": "total",
      "label": {
        "show": true
      },
      "emphasis": {
        "focus": "series"
      },
      "data": $DATA
    }
  ]
}`

func genVerticalBarProfile(cate []string, data []int, isRotate bool) string {
	var afterRotateProfile string
	if isRotate {
		afterRotateProfile = strings.Replace(verticalBarProfile, "$ROTATE", ROTATE, 1)
	} else {
		afterRotateProfile = strings.Replace(verticalBarProfile, "$ROTATE", "", 1)
	}
	cateTemp := strListTemplate(cate)
	dataTemp := intListTemplate(data)
	return strings.Replace(
		strings.Replace(afterRotateProfile, "$CATE", cateTemp, 1),
		"$DATA", dataTemp, 1)
}
