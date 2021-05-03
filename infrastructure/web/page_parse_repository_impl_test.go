package web

import (
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"

	"github.com/s14t284/apple-maitained-bot/domain"
	"github.com/s14t284/apple-maitained-bot/domain/model"
)

const (
	detailMacHTML = `
<div class="as-productinfosection-panel Overview-panel row">
    <div class="as-productinfosection-sidepanel column large-3 small-12">
        <h3 data-autom="sectionTitle">概要</h3>
    </div>
    <div class="as-productinfosection-mainpanel column large-9 small-12">

            <div class="para-list">
            <p>
                2019年11月発売
            </p>
        </div>
        <div class="para-list">
            <p>
                Touch IDセンサーが組み込まれたTouch Bar
            </p>
        </div>
        <div class="para-list">
            <p>
                IPSテクノロジー搭載16インチ（対角）LEDバックライトディスプレイ、3,072 x 1,920ピクセル標準解像度、226ppi、数百万色以上対応
            </p>
        </div>
        <div class="para-list">
            <p>
                16GB 2,666MHz DDR4オンボードメモリ
            </p>
        </div>
        <div class="para-list">
            <p>
                1TB SSD<sup>1</sup>
            </p>
        </div>
        <div class="para-list">
            <p>
                720p FaceTime HDカメラ
            </p>
        </div>
        <div class="para-list as-pdp-lastparalist">
            <p>
                AMD Radeon Pro 5300M（4GB GDDR6メ‍モ‍リ搭載）
            </p>
        </div>
    </div>
</div>
`
	detailMacMiniHTML = `
<div class="as-productinfosection-mainpanel column large-9 small-12">

            <div class="para-list">
            <p>
                2018年10月発売モデル
            </p>
        </div>
        <div class="para-list">
            <p>
                8GB 2,666MHz DDR4 SO-DIMMメモリ
            </p>
        </div>
        <div class="para-list">
            <p>
                512GB PCIeベースSSD<sup>1</sup>
            </p>
        </div>
        <div class="para-list">
            <p>
                Thunderbolt 3ポート（最大40Gbps）x 4
            </p>
        </div>
        <div class="para-list">
            <p>
                Intel UHD Graphics 630
            </p>
        </div>
        <div class="para-list as-pdp-lastparalist">
            <p>
                ギガビットEthernetポート
            </p>
        </div>
    </div>
`
	detailIPadHTML = `
<div class="as-productinfosection-panel Overview-panel row">

    <div class="as-productinfosection-sidepanel column large-3 small-12">
        <h3 data-autom="sectionTitle">概要</h3>
    </div><div class="as-productinfosection-mainpanel column large-9 small-12">

            <div class="para-list">
            <p>
                2015年9月発売モデル
            </p>
        </div>
        <div class="para-list">
            <p>
                Wi-Fi（802.11a/b/g/n/ac）
            </p>
        </div>
        <div class="para-list">
            <p>
                Bluetooth 4.2テクノロジー
            </p>
        </div>
        <div class="para-list">
            <p>
                7.9インチのRetinaディスプレイ
            </p>
        </div>
        <div class="para-list">
            <p>
                8メガピクセルiSightカメラ
            </p>
        </div>
        <div class="para-list">
            <p>
                FaceTime HDカメラ
            </p>
        </div>
        <div class="para-list">
            <p>
                1080p HDビデオ撮影
            </p>
        </div>
        <div class="para-list">
            <p>
                64ビットアーキテクチャ搭載A8チップ
            </p>
        </div>
        <div class="para-list">
            <p>
                M8モーションコプロセッサ
            </p>
        </div>
        <div class="para-list">
            <p>
                10時間のバッテリー駆動時間
            </p>
        </div>
        <div class="para-list">
            <p>
                マルチタッチスクリーン
            </p>
        </div>
        <div class="para-list as-pdp-lastparalist">
            <p>
                304 g、6.1 mm
            </p>
        </div>
    </div>
</div>
`
	detailWatchHTML = `
<div class="as-productinfosection-panel Overview-panel row">
    <div class="as-productinfosection-sidepanel column large-3 small-12">
        <h3 data-autom="sectionTitle">概要</h3>
    </div>
    <div class="as-productinfosection-mainpanel column large-9 small-12">

            <div class="para-list">
            <p>
                2018年9月発売モデル
            </p>
        </div>
        <div class="para-list">
            <p>
                S4（64ビットデュアルコアプロセッサ搭載。S3プロセッサと比べて最大2倍高速）
            </p>
        </div>
        <div class="para-list">
            <p>
                50メートルの耐水性能<sup>1</sup>
            </p>
        </div>
        <div class="para-list">
            <p>
                感圧タッチ対応LTPO OLED Retinaディスプレイ（1,000ニトの輝度）
            </p>
        </div>
        <div class="para-list">
            <p><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">
                Wi-Fi（802.11b / g / n 2.4GHz）
            </font></font></p>
        </div>
        <div class="para-list">
            <p><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">
                Bluetooth 5.0
            </font></font></p>
        </div>
        <div class="para-list">
            <p>
                光学式心拍センサー
            </p>
        </div>
        <div class="para-list">
            <p>
                進化した加速度センサー
            </p>
        </div>
        <div class="para-list">
            <p>
                進化したジャイロスコープ
            </p>
        </div>
        <div class="para-list">
            <p>
                環境光センサー
            </p>
        </div>
        <div class="para-list">
            <p>
                容量16GB<sup>2</sup>
            </p>
        </div>
        <div class="para-list as-pdp-lastparalist">
            <p>
                全面セラミックとサファイアクリスタルの裏蓋
            </p>
        </div>
    </div>
</div>
`
)

func TestLoadMacInformationFromTitle(t *testing.T) {
	a := assert.New(t)
	{
		// 16インチMacBook Proの場合
		page := domain.Page{
			Title:     "16インチMacBook Pro 2.4GHz 8コアIntel Core i9 Retinaディスプレイモデル - スペースグレイ [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://www.apple.com",
		}
		mac := &model.Mac{}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadMacInformationFromTitle(mac, page)
		a.NoError(err)
		expectedName := strings.Replace(page.Title, " [整備済製品]", "", 1)
		a.Equal(float32(16), mac.Inch)
		a.Equal("2.4GHz 8コアIntel Core i9", mac.CPU)
		a.Equal("スペースグレイ", mac.Color)
		a.Equal(30000, mac.Amount)
		a.Equal(expectedName, mac.Name)
		a.Equal("https://www.apple.com", mac.URL)
	}
	{
		// 15.4インチMacBook Proの場合
		// 16インチMacBook Proの場合
		page := domain.Page{
			Title:     "15.4インチMacBook Pro 2.9GHz 6コアIntel Core i9 Retinaディスプレイモデル - シルバー [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://www.apple.com",
		}
		mac := &model.Mac{}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadMacInformationFromTitle(mac, page)
		a.NoError(err)
		expectedName := strings.Replace(page.Title, " [整備済製品]", "", 1)
		a.Equal(mac.Inch, float32(15.4))
		a.Equal(mac.CPU, "2.9GHz 6コアIntel Core i9")
		a.Equal(mac.Color, "シルバー")
		a.Equal(mac.Amount, 30000)
		a.Equal(expectedName, mac.Name)
		a.Equal(mac.URL, "https://www.apple.com")
	}
	{
		// 13.3インチMacBook Proの場合
		page := domain.Page{
			Title:     "13.3インチMacBook Pro 1.4GHzクアッドコアIntel Core i5 Retinaディスプレイモデル - スペースグレイ [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://www.apple.com",
		}
		mac := &model.Mac{}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadMacInformationFromTitle(mac, page)
		a.NoError(err)
		expectedName := strings.Replace(page.Title, " [整備済製品]", "", 1)
		a.Equal(mac.Inch, float32(13.3))
		a.Equal(mac.CPU, "1.4GHzクアッドコアIntel Core i5")
		a.Equal(mac.Color, "スペースグレイ")
		a.Equal(mac.Amount, 30000)
		a.Equal(expectedName, mac.Name)
		a.Equal(mac.URL, "https://www.apple.com")
	}
	{
		mac := &model.Mac{}
		// Mac Miniの場合
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		page := domain.Page{
			Title:     "Mac mini 3.0GHz 6コアIntel Core i5 - スペースグレイ [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://www.apple.com",
		}
		err = pageParser.loadMacInformationFromTitle(mac, page)
		if err != nil {
			t.FailNow()
		}
		expectedName := strings.Replace(page.Title, " [整備済製品]", "", 1)
		a.Equal(mac.Inch, float32(0.0))
		a.Equal(mac.CPU, "3.0GHz 6コアIntel Core i5")
		a.Equal(mac.Color, "スペースグレイ")
		a.Equal(mac.Amount, 30000)
		a.Equal(expectedName, mac.Name)
		a.Equal(mac.URL, "https://www.apple.com")
	}
	{
		// 税込表示にも対応
		mac := &model.Mac{}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		page := domain.Page{
			AmountStr: "30,000円（税込）",
		}
		err = pageParser.loadMacInformationFromTitle(mac, page)
		if err != nil {
			t.FailNow()
		}
		a.Equal(mac.Amount, 30000)
	}
	// TODO: MacBook Airに関するテストを増やす
	// TODO: MacBook以外のMacに関するテストを増やす
	// {
	// 	// Mac Proの場合
	// pageParser, err := initializePageParserImpl()
	// if err != nil {
	// 	t.FailNow()
	// }
	// 	page := domain.Page{
	// 		Title:     "Mac Pro 3.2GHz 16コア Intel Xeon W、Radeon Pro 580X [整備済製品]",
	// 		AmountStr: "30,000円（税別）",
	// 		DetailURL: "https://www.apple.com",
	// 	}
	// 	pageParser.loadMacInformationFromTitle(mac, page)
	// 	a.Equal(mac.Inch, float32(0.0))
	// 	a.Equal(mac.CPU, "3.2GHz 16コア Intel Xeon W、Radeon Pro 580X")
	// 	a.Equal(mac.Color, "")
	// 	a.Equal(mac.Amount, 30000)
	// 	a.Equal(mac.Name, "Mac Pro")
	// 	a.Equal(mac.URL, "https://www.apple.com")
	// }

}

func TestLoadMacInformationFromDetailHTML(t *testing.T) {
	a := assert.New(t)
	{
		// 正常系
		// MacBookProの場合
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(detailMacHTML))
		mac := &model.Mac{}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadMacInformationFromDetailHTML(mac, doc)
		a.NoError(err)
		a.Equal(getYearAndMonth(2019, 11), mac.ReleaseDate)
		a.Equal(true, mac.TouchBar)
		a.Equal(1000, mac.Storage)
	}
	{
		// 正常系
		// Mac miniの場合
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(detailMacMiniHTML))
		mac := &model.Mac{}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadMacInformationFromDetailHTML(mac, doc)
		a.NoError(err)
		a.Equal(getYearAndMonth(2018, 10), mac.ReleaseDate)
		a.Equal(512, mac.Storage)
	}
}

func TestLoadIPadInformationFromTitle(t *testing.T) {
	a := assert.New(t)
	ipad := &model.IPad{}
	{
		// IPad Proの場合
		// インチ数がタイトルに含まれているが、インチは詳細ページから取得する
		page := domain.Page{
			Title:     "12.9インチiPad Pro Wi-Fi + Cellular 512GB - スペースグレイ（第2世代） [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://www.apple.com",
		}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadIPadInformationFromTitle(ipad, page)
		a.NoError(err)
		expectedName := strings.Replace(page.Title, " [整備済製品]", "", 1)
		a.Equal(ipad.Color, "スペースグレイ")
		a.Equal(expectedName, ipad.Name)
		a.Equal(ipad.Storage, 512)
		a.Equal(ipad.Amount, 30000)
		a.Equal(ipad.URL, "https://www.apple.com")

	}
	{
		// IPad Airの場合
		page := domain.Page{
			Title:     "iPad Air Wi-Fiモデル 64GB - ゴールド [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://www.apple.com",
		}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadIPadInformationFromTitle(ipad, page)
		a.NoError(err)
		expectedName := strings.Replace(page.Title, " [整備済製品]", "", 1)
		a.Equal(ipad.Color, "ゴールド")
		a.Equal(expectedName, ipad.Name)
		a.Equal(ipad.Storage, 64)
		a.Equal(ipad.Amount, 30000)
		a.Equal(ipad.URL, "https://www.apple.com")
	}
	{
		// IPad miniの場合
		page := domain.Page{
			Title:     "iPad mini 4 Wi-Fi 128GB - スペースグレイ [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://www.apple.com",
		}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadIPadInformationFromTitle(ipad, page)
		a.NoError(err)
		expectedName := strings.Replace(page.Title, " [整備済製品]", "", 1)
		a.Equal(ipad.Color, "スペースグレイ")
		a.Equal(expectedName, ipad.Name)
		a.Equal(ipad.Storage, 128)
		a.Equal(ipad.Amount, 30000)
		a.Equal(ipad.URL, "https://www.apple.com")
	}
	{
		// 通常IPadの場合
		page := domain.Page{
			Title:     "iPad Wi-Fi 1TB - シルバー（第7世代） [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://www.apple.com",
		}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadIPadInformationFromTitle(ipad, page)
		a.NoError(err)
		expectedName := strings.Replace(page.Title, " [整備済製品]", "", 1)
		a.Equal(ipad.Color, "シルバー")
		a.Equal(expectedName, ipad.Name)
		a.Equal(ipad.Storage, 1000)
		a.Equal(ipad.Amount, 30000)
		a.Equal(ipad.URL, "https://www.apple.com")
	}
}

func TestLoadIPadInformationFromDetailHTML(t *testing.T) {
	a := assert.New(t)
	ipad := &model.IPad{}
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(detailIPadHTML))
	{
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadIPadInformationFromDetailHTML(ipad, doc)
		a.NoError(err)
		a.Equal(getYearAndMonth(2015, 9), ipad.ReleaseDate)
		a.Equal("8メガピクセルiSightカメラ", ipad.Camera)
		a.Equal(float32(7.9), ipad.Inch)
	}

}
func TestLoadWatchInformationFromTitle(t *testing.T) {
	a := assert.New(t)
	watch := &model.Watch{}
	{
		// Apple Watch Series 4（GPS + Cellularモデル）の場合
		page := domain.Page{
			Title:     "Apple Watch Series 4（GPS + Cellularモデル）- 44mmシルバーアルミニウムケースとホワイトスポーツバンド [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://www.apple.com",
		}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadWatchInformationFromTitle(watch, page)
		a.NoError(err)
		expectedName := strings.Replace(page.Title, " [整備済製品]", "", 1)
		a.Equal(watch.Color, "シルバー")
		a.Equal(expectedName, watch.Name)
		a.Equal(watch.Amount, 30000)
		a.Equal(watch.URL, "https://www.apple.com")
	}
	{
		// Apple Watch Series 4（GPSモデル）の場合
		// GPSとGPS + Cellularは変更できるため、取得していない
		page := domain.Page{
			Title:     "Apple Watch Series 4（GPSモデル）- 44mmゴールドアルミニウムケースとピンクサンドスポーツバンド [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://www.apple.com",
		}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadWatchInformationFromTitle(watch, page)
		a.NoError(err)
		expectedName := strings.Replace(page.Title, " [整備済製品]", "", 1)
		a.Equal(watch.Color, "ゴールド")
		a.Equal(expectedName, watch.Name)
		a.Equal(watch.Amount, 30000)
		a.Equal(watch.URL, "https://www.apple.com")
	}
	{
		// Apple Watch Series Nike+ Series 4の場合
		page := domain.Page{
			Title:     "Apple Watch Nike+ Series 4（GPS + Cellularモデル）- 40mmスペースグレイアルミニウムケースとアンスラサイト/ブラックNikeスポーツバンド [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://www.apple.com",
		}
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadWatchInformationFromTitle(watch, page)
		a.NoError(err)
		expectedName := strings.Replace(page.Title, " [整備済製品]", "", 1)
		a.Equal(watch.Color, "スペースグレイ")
		a.Equal(expectedName, watch.Name)
		a.Equal(watch.Amount, 30000)
		a.Equal(watch.URL, "https://www.apple.com")
	}
}

func TestLoadWatchInformationFromDetailHTML(t *testing.T) {
	a := assert.New(t)
	watch := &model.Watch{}
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(detailWatchHTML))
	{
		pageParser, err := initializePageParserImpl()
		if err != nil {
			t.FailNow()
		}
		err = pageParser.loadWatchInformationFromDetailHTML(watch, doc)
		a.NoError(err)
		a.Equal(getYearAndMonth(2018, 9), watch.ReleaseDate)
		a.Equal(16, watch.Storage)
	}
}

func initializePageParserImpl() (*PageParseRepositoryImpl, error) {
	return NewPageParseRepositoryImpl()
}

func getYearAndMonth(year, month int) time.Time {
	timeZone, _ := time.LoadLocation("Asia/Tokyo")
	return time.Date(year, time.Month(month), 1, 9, 0, 0, 0, timeZone)
}
