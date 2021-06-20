package maine

import (
	"fmt"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func maine() {
	begin := time.Now()

	//darkGrayColor := getDarkGrayColor()
	grayColor := getGrayColor()
	//blueColor := getBlueColor()
	header := getHeader()
	contents := getContents()

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)

	m.RegisterHeader(func() {
		m.Row(20, func() {

			m.ColSpace(6)
		})
	})

	
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Report liquidity", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 4, 2, 3,3},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 4, 2, 3,3},
		},
		Align:                consts.Center,
		AlternatedBackground: &grayColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})


	err := m.OutputFileAndClose("./billing.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}

func getHeader() []string {
	return []string{"Platform", "Project", "Liquidity", "Fees"}
}

func getContents() [][]string {
	return [][]string{
		{"Uniswap", "ETH/NIF", "$12000", "200"},
		{"Uniswap", "ETH/DODO", "$11000","150"},
		{"Uniswap", "ETH/WILD", "$15000","100"},
	}
}

func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getGrayColor() color.Color {
	return color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getBlueColor() color.Color {
	return color.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor() color.Color {
	return color.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}

