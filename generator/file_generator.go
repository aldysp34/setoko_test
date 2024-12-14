package generator

type ModelData struct {
	Key            string
	Value          string
	IsTotalPayment bool
}

type ListModelData struct {
	HeaderData string
	ModelData  []ModelData
}

type Paper struct {
	MarginSetup                MarginSetup
	RectSetup                  RectSetup
	TransformSetup             TransformSetup
	LineHt                     float64
	TotalPaymentFont           FontSize
	ValueFont                  FontSize
	FooterSetup                FooterSetup
	WLogo1                     LogoSetup
	WLogo2                     LogoSetup
	TransactionTextSetup       TransactionTextSetup
	TransactionTextSetupResize TransactionTextSetup
	BottomSetup                BottomSetup
	ValueCellSetup             CellSetup
	HeaderSetup                HeaderSetup
}

type FontSize struct {
	ValueFontSize  float64
	HeaderFontSize float64
}

type MarginSetup struct {
	LMargin float64
	TMargin float64
	RMargin float64
}

type RectSetup struct {
	X float64
	Y float64
	W float64
	H float64
}

type TransformSetup struct {
	X struct {
		A float64
		B float64
	}
	Y struct {
		A float64
		B float64
	}
	TextX struct {
		A float64
		B float64
	}
	TextY struct {
		A float64
		B float64
	}
	Angle float64
	I     struct {
		Min float64
		Max float64
	}
	J struct {
		Min float64
		Max float64
	}
}

type LogoSetup struct {
	X    float64
	Y    float64
	W    float64
	H    float64
	Flow bool
}

type FooterSetup struct {
	Y           float64
	WordSpacing float64
	FontSize    float64
}

type TransactionTextSetup struct {
	FontSize   float64
	UpperSpace float64
	LowerSpace float64
}

type BottomSetup struct {
	BottomLimit      float64
	BottomLimitMinus float64
	FontSize         float64
}

type CellSetup struct {
	W1         float64
	W2         float64
	WMultiCell float64
	H1         float64
	H2         float64
	HMultiCell float64
	Ln1        float64
	Ln2        float64
}

type HeaderSetup struct {
	Space1   float64
	Space2   float64
	W        float64
	H        float64
	X        float64
	Y        float64
	FontSize float64
}

func GetPaperA5() Paper {
	var paper Paper

	paper = Paper{
		MarginSetup: MarginSetup{
			LMargin: 6.35,
			TMargin: 6.35,
			RMargin: 6.35,
		},
		RectSetup: RectSetup{
			X: 12.7,
			Y: 12.7,
			W: 148 - 25.4,
			H: 210 - 25.4,
		},
		TransformSetup: TransformSetup{
			X: struct {
				A float64
				B float64
			}{A: 6.35, B: 30},
			Y: struct {
				A float64
				B float64
			}{A: 15, B: 10.7},
			TextX: struct {
				A float64
				B float64
			}{A: 6.35, B: 30},
			TextY: struct {
				A float64
				B float64
			}{A: 15, B: 10.7},
			Angle: 30,
			I: struct {
				Min float64
				Max float64
			}{
				Min: 0.04,
				Max: 5,
			},
			J: struct {
				Min float64
				Max float64
			}{
				Min: 0.9,
				Max: 18.5,
			},
		},
		WLogo1: LogoSetup{
			X:    20.05,
			Y:    19.05,
			W:    50,
			H:    5.5,
			Flow: false,
		},
		WLogo2: LogoSetup{
			X:    37.1,
			Y:    19.05,
			W:    15,
			H:    5.5,
			Flow: false,
		},
		LineHt: 5.5,
		TotalPaymentFont: FontSize{
			ValueFontSize:  12,
			HeaderFontSize: 12,
		},
		ValueFont: FontSize{
			ValueFontSize:  10,
			HeaderFontSize: 10,
		},
		FooterSetup: FooterSetup{
			Y:           -27.5,
			WordSpacing: 1,
			FontSize:    7,
		},
		TransactionTextSetup: TransactionTextSetup{
			FontSize:   13,
			UpperSpace: 8,
			LowerSpace: 20,
		},
		TransactionTextSetupResize: TransactionTextSetup{
			FontSize:   11,
			UpperSpace: 6,
			LowerSpace: 20,
		},
		BottomSetup: BottomSetup{
			BottomLimit:      40,
			BottomLimitMinus: 35,
			FontSize:         15,
		},
		ValueCellSetup: CellSetup{
			W1:         10,
			W2:         1,
			WMultiCell: 99.3,
			H2:         5.5,
			H1:         5.5,
			HMultiCell: 5.5,
			Ln1:        1.5,
			Ln2:        0.1,
		},
		HeaderSetup: HeaderSetup{
			Space1:   1,
			Space2:   12.5,
			W:        0,
			H:        5.5,
			X:        22.4,
			Y:        3.5,
			FontSize: 7,
		},
	}

	return paper
}
