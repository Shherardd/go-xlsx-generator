package styles

import "github.com/xuri/excelize/v2"

var DataMoneyFormat = &excelize.Style{
	NumFmt: 4,
}

var HeadStyles = &excelize.Style{
	Font:      bigCenteredFont,
	Alignment: centerAlignment,
}

var HeaderStyles = &excelize.Style{
	Border: []excelize.Border{
		borderTopThick, borderBorderBottomThick,
	},
	Alignment: centerAlignment,
	Font:      boldFont,
}

var TotalesStyles = &excelize.Style{
	NumFmt: 4,
	Border: []excelize.Border{
		borderTopThick, borderBorderBottomThick,
	},
	Alignment: rightAlignment,
	Font:      boldFont,
}
