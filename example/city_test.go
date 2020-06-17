package example

import (
	"log"
	"testing"
)

var cities Citys

func init(){
	cities = Citys{
		City{`北海道`, `旭川市`, 339605, false},
		City{`北海道`, `釧路市`, 174742, false},
		City{`北海道`, `江別市`, 120636, false},
		City{`北海道`, `札幌市`, 1952356, true},
		City{`北海道`, `小樽市`, 121924, false},
		City{`北海道`, `帯広市`, 169327, false},
		City{`北海道`, `苫小牧市`, 172737, false},
		City{`北海道`, `函館市`, 265979, false},
		City{`北海道`, `北見市`, 121226, false},
		City{`青森県`, `弘前市`, 177411, false},
		City{`青森県`, `青森市`, 287648, true},
		City{`青森県`, `八戸市`, 231257, false},
		City{`岩手県`, `一関市`, 121583, false},
		City{`岩手県`, `奥州市`, 119422, false},
		City{`岩手県`, `盛岡市`, 297631, true},
		City{`宮城県`, `石巻市`, 147214, false},
		City{`宮城県`, `仙台市`, 1082159, true},
		City{`宮城県`, `大崎市`, 133391, false},
		City{`秋田県`, `横手市`, 92197, false},
		City{`秋田県`, `秋田市`, 315814, true},
		City{`山形県`, `山形市`, 253832, true},
		City{`山形県`, `酒田市`, 106244, false},
		City{`山形県`, `鶴岡市`, 129652, false},
		City{`福島県`, `いわき市`, 350237, false},
		City{`福島県`, `会津若松市`, 124062, false},
		City{`福島県`, `郡山市`, 335444, false},
		City{`福島県`, `福島市`, 294247, true},
		City{`茨城県`, `つくば市`, 226963, false},
		City{`茨城県`, `水戸市`, 270783, true},
		City{`茨城県`, `日立市`, 185054, false},
		City{`栃木県`, `宇都宮市`, 518594, true},
		City{`栃木県`, `小山市`, 166760, false},
		City{`栃木県`, `栃木市`, 159211, false},
		City{`群馬県`, `伊勢崎市`, 208814, false},
		City{`群馬県`, `高崎市`, 370884, false},
		City{`群馬県`, `前橋市`, 336154, true},
		City{`群馬県`, `太田市`, 219807, false},
		City{`埼玉県`, `さいたま市`, 1263979, true},
		City{`埼玉県`, `越谷市`, 337498, false},
		City{`埼玉県`, `熊谷市`, 198742, false},
		City{`埼玉県`, `春日部市`, 232709, false},
		City{`埼玉県`, `所沢市`, 340386, false},
		City{`埼玉県`, `上尾市`, 225196, false},
		City{`埼玉県`, `川越市`, 350745, false},
		City{`埼玉県`, `川口市`, 578112, false},
		City{`埼玉県`, `草加市`, 247034, false},
		City{`千葉県`, `市原市`, 274656, false},
		City{`千葉県`, `市川市`, 481732, false},
		City{`千葉県`, `松戸市`, 483480, false},
		City{`千葉県`, `千葉市`, 971882, true},
		City{`千葉県`, `船橋市`, 622890, false},
		City{`千葉県`, `柏市`, 413954, false},
		City{`千葉県`, `八千代市`, 193152, false},
		City{`千葉県`, `流山市`, 174373, false},
		City{`東京都`, `町田市`, 432348, false},
		City{`東京都`, `調布市`, 229061, false},
		City{`東京都`, `八王子市`, 577513, false},
		City{`東京都`, `府中市`, 260274, false},
		City{`神奈川県`, `横須賀市`, 406586, false},
		City{`神奈川県`, `横浜市`, 3724844, true},
		City{`神奈川県`, `茅ヶ崎市`, 239348, false},
		City{`神奈川県`, `厚木市`, 225714, false},
		City{`神奈川県`, `小田原市`, 194086, false},
		City{`神奈川県`, `川崎市`, 1475213, false},
		City{`神奈川県`, `相模原市`, 720780, false},
		City{`神奈川県`, `大和市`, 232922, false},
		City{`神奈川県`, `藤沢市`, 423894, false},
		City{`神奈川県`, `平塚市`, 258227, false},
		City{`新潟県`, `上越市`, 196987, false},
		City{`新潟県`, `新潟市`, 810157, true},
		City{`新潟県`, `長岡市`, 275133, false},
		City{`富山県`, `高岡市`, 172125, false},
		City{`富山県`, `富山市`, 418686, true},
		City{`石川県`, `金沢市`, 465699, true},
		City{`石川県`, `小松市`, 106919, false},
		City{`石川県`, `白山市`, 109287, false},
		City{`福井県`, `坂井市`, 90280, false},
		City{`福井県`, `福井市`, 265904, true},
		City{`山梨県`, `甲斐市`, 74386, false},
		City{`山梨県`, `甲府市`, 193125, true},
		City{`長野県`, `松本市`, 243293, false},
		City{`長野県`, `上田市`, 156827, false},
		City{`長野県`, `長野市`, 377598, true},
		City{`岐阜県`, `岐阜市`, 406735, true},
		City{`岐阜県`, `大垣市`, 159879, false},
		City{`静岡県`, `沼津市`, 195633, false},
		City{`静岡県`, `静岡市`, 704989, true},
		City{`静岡県`, `磐田市`, 167210, false},
		City{`静岡県`, `浜松市`, 797980, false},
		City{`静岡県`, `富士市`, 248399, false},
		City{`愛知県`, `安城市`, 184140, false},
		City{`愛知県`, `一宮市`, 380868, false},
		City{`愛知県`, `岡崎市`, 381051, false},
		City{`愛知県`, `春日井市`, 306508, false},
		City{`愛知県`, `豊橋市`, 374765, false},
		City{`愛知県`, `豊川市`, 182436, false},
		City{`愛知県`, `豊田市`, 422542, false},
		City{`愛知県`, `名古屋市`, 2295638, true},
		City{`三重県`, `四日市市`, 311031, false},
		City{`三重県`, `津市`, 279886, true},
		City{`三重県`, `鈴鹿市`, 196403, false},
		City{`滋賀県`, `草津市`, 137247, false},
		City{`滋賀県`, `大津市`, 340973, true},
		City{`京都府`, `宇治市`, 184678, false},
		City{`京都府`, `京都市`, 1475183, true},
		City{`大阪府`, `茨木市`, 280033, false},
		City{`大阪府`, `岸和田市`, 194911, false},
		City{`大阪府`, `高槻市`, 351829, false},
		City{`大阪府`, `堺市`, 839310, false},
		City{`大阪府`, `寝屋川市`, 237518, false},
		City{`大阪府`, `吹田市`, 374468, false},
		City{`大阪府`, `大阪市`, 2691185, true},
		City{`大阪府`, `東大阪市`, 502784, false},
		City{`大阪府`, `八尾市`, 268800, false},
		City{`大阪府`, `豊中市`, 395479, false},
		City{`大阪府`, `枚方市`, 404152, false},
		City{`大阪府`, `和泉市`, 186109, false},
		City{`兵庫県`, `伊丹市`, 196883, false},
		City{`兵庫県`, `加古川市`, 267435, false},
		City{`兵庫県`, `神戸市`, 1537272, true},
		City{`兵庫県`, `西宮市`, 487850, false},
		City{`兵庫県`, `尼崎市`, 452563, false},
		City{`兵庫県`, `姫路市`, 535664, false},
		City{`兵庫県`, `宝塚市`, 224903, false},
		City{`兵庫県`, `明石市`, 293409, false},
		City{`奈良県`, `橿原市`, 124111, false},
		City{`奈良県`, `奈良市`, 360310, true},
		City{`和歌山県`, `田辺市`, 74770, false},
		City{`和歌山県`, `和歌山市`, 364154, true},
		City{`鳥取県`, `鳥取市`, 193717, true},
		City{`鳥取県`, `米子市`, 149313, false},
		City{`島根県`, `出雲市`, 171938, false},
		City{`島根県`, `松江市`, 206230, true},
		City{`岡山県`, `岡山市`, 719474, true},
		City{`岡山県`, `倉敷市`, 477118, false},
		City{`岡山県`, `津山市`, 103746, false},
		City{`広島県`, `呉市`, 228552, false},
		City{`広島県`, `広島市`, 1194034, true},
		City{`広島県`, `東広島市`, 192907, false},
		City{`広島県`, `福山市`, 464811, false},
		City{`山口県`, `宇部市`, 169429, false},
		City{`山口県`, `下関市`, 268517, false},
		City{`山口県`, `岩国市`, 136757, false},
		City{`山口県`, `山口市`, 197422, true},
		City{`山口県`, `周南市`, 144842, false},
		City{`山口県`, `防府市`, 115942, false},
		City{`徳島県`, `阿南市`, 73019, false},
		City{`徳島県`, `徳島市`, 258554, true},
		City{`香川県`, `丸亀市`, 110010, false},
		City{`香川県`, `高松市`, 420748, true},
		City{`愛媛県`, `今治市`, 158114, false},
		City{`愛媛県`, `松山市`, 514865, true},
		City{`高知県`, `高知市`, 337190, true},
		City{`高知県`, `南国市`, 47982, false},
		City{`福岡県`, `久留米市`, 304552, false},
		City{`福岡県`, `飯塚市`, 129146, false},
		City{`福岡県`, `福岡市`, 1538681, true},
		City{`福岡県`, `北九州市`, 961286, false},
		City{`佐賀県`, `佐賀市`, 236372, true},
		City{`佐賀県`, `唐津市`, 122785, false},
		City{`長崎県`, `佐世保市`, 255439, false},
		City{`長崎県`, `長崎市`, 429508, true},
		City{`長崎県`, `諫早市`, 138078, false},
		City{`熊本県`, `熊本市`, 740822, true},
		City{`熊本県`, `八代市`, 127472, false},
		City{`大分県`, `大分市`, 478146, true},
		City{`大分県`, `別府市`, 122138, false},
		City{`宮崎県`, `延岡市`, 125159, false},
		City{`宮崎県`, `宮崎市`, 401138, true},
		City{`宮崎県`, `都城市`, 165029, false},
		City{`鹿児島県`, `鹿屋市`, 103608, false},
		City{`鹿児島県`, `鹿児島市`, 599814, true},
		City{`鹿児島県`, `霧島市`, 125857, false},
		City{`沖縄県`, `那覇市`, 319435, true},
	}
}


func TestCity(t *testing.T) {
	// total
	log.Println("total count: ",  cities.Count())
	log.Println("total population: ",  cities.Select(`Population`).AsInt().Sum())
	log.Println("max population: ",  cities.Select(`Population`).AsInt().Max())
	log.Println("sort desc population: ",  cities.SortBy(func(a City,b City)bool{return a.Population > b.Population})[0])
	log.Println("sort asc population: ",  cities.SortBy(func(a City,b City)bool{return a.Population < b.Population})[0])

	// Where
	capitals := cities.Where(func(x City)bool{return x.IsCapital})
	log.Println("total count: ",  capitals.Count())
	log.Println("total population: ",  capitals.Select(`Population`).AsInt().Sum())
	log.Println("max population: ",  capitals.Select(`Population`).AsInt().Max())
	log.Println("sort desc population: ",  capitals.SortBy(func(a City,b City)bool{return a.Population > b.Population})[0])
	log.Println("sort asc population: ",  capitals.SortBy(func(a City,b City)bool{return a.Population < b.Population})[0])

	// Under 1,000,000 - over 500,000
	under1M := cities.Where(func(x City)bool{return 50*10000 < x.Population && x.Population < 100*10000})
	log.Println("total count: ",  under1M.Count())
	log.Println("total population: ",  under1M.Select(`Population`).AsInt().Sum())
	log.Println("max population: ",  under1M.Select(`Population`).AsInt().Min())
	log.Println("max population: ",  under1M.Select(`Population`).AsInt().Max())
	log.Println("sort desc population: ",  under1M.SortBy(func(a City,b City)bool{return a.Population > b.Population})[0])
	log.Println("sort asc population: ",  under1M.SortBy(func(a City,b City)bool{return a.Population < b.Population})[0])

	// 北海道
	hokkaido := cities.Where(func(x City)bool{return x.PrefectureName== `北海道`})
	log.Println("total count: ",  hokkaido.Count())
	log.Println("total population: ",  hokkaido.Select(`Population`).AsInt().Sum())
	log.Println("sort desc population: ",  hokkaido.SortBy(func(a City,b City)bool{return a.Population > b.Population})[0])
	log.Println("sort asc population: ",  hokkaido.SortBy(func(a City,b City)bool{return a.Population < b.Population})[0])

}


