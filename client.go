package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)
type Response struct {
	Pair     string `json:"Pair"`
	Total   string `json:"Total"`
	Volume string    `json:"Volume"`
	Fees string `json:"Fees"`
}
func main() {
	fmt.Println("Calling API...")
	client := &http.Client{}
 	req, err := http.NewRequest("GET", "http://127.0.0.1:5000/v2/0x63607de7ae773638d012561a01383ab8ac321371", nil)
 	if err != nil {
  		fmt.Print(err.Error())
 	}
	//req.Header.Add("Accept", "application/json")
	//req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)

 	//starts creation of PDF
 	begin := time.Now()

	//darkGrayColor := getDarkGrayColor()
	grayColor := getGrayColor()
	//blueColor := getBlueColor()
	header := getHeader()
	contents := getContents(responseObject)

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

	errDocument := m.OutputFileAndClose("./billing.pdf")
	if errDocument != nil {
		fmt.Println("Could not save PDF:", errDocument)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}

func getHeader() []string {
	return []string{"Platform", "Project", "Liquidity", "Fees"}
}

func getContents(pair Response) [][]string {
	return [][]string{
		{"Uniswap", pair.Pair, pair.Total, pair.Fees},
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