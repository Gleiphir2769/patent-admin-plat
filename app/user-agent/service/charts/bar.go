package charts

import "strings"

const barProfile = `{
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
      "type": "category",
      "data": $CATE,
      $ROTATE
    }
  ],
  "yAxis": [
    {
      "type": "value"
    }
  ],
  "series": [
    {
      "name": "Direct",
      "type": "bar",
      "barWidth": "60%",
      "data": $DATA
    }
  ]
}`

const ROTATE = `
      "axisTick": {
        "alignWithLabel": true
      },
      "axisLabel": {
        "interval": 0,
        "rotate": 45
      }`

func genBarProfile(cate []string, data []int, isRotate bool) string {
	var afterRotateProfile string
	if isRotate {
		afterRotateProfile = strings.Replace(barProfile, "$ROTATE", ROTATE, 1)
	} else {
		afterRotateProfile = strings.Replace(barProfile, "$ROTATE", "", 1)
	}
	cateTemp := strListTemplate(cate)
	dataTemp := intListTemplate(data)
	return strings.Replace(
		strings.Replace(afterRotateProfile, "$CATE", cateTemp, 1),
		"$DATA", dataTemp, 1)
}
