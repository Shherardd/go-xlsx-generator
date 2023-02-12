package styles

import "github.com/xuri/excelize/v2"

var rightAlignment = &excelize.Alignment{
	Horizontal: "right",
}

var centerAlignment = &excelize.Alignment{
	Horizontal: "center",
}

var bigCenteredFont = &excelize.Font{
	VertAlign: "center",
	Family:    "Arial",
	Size:      14,
}

var borderTopThick = excelize.Border{
	Type:  "top",
	Style: 5,
	Color: "636363",
}
var borderBorderBottomThick = excelize.Border{
	Type:  "bottom",
	Style: 5,
	Color: "72b834",
}

var boldFont = &excelize.Font{
	Bold:   true,
	Family: "Arial",
}
