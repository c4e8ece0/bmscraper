// Package sport contains constants for sport-names
package sport

// First name
const (
	BADMINTON       = 1 << iota
	BANDY           // field hockey
	BASEBALL        //
	BASKETBALL      //
	BEACHVOLLEYBALL //
	BEASAL          //Beach Soccer
	BICYCLE         //
	BOXING          //
	CRICKET         //
	CYBER           //
	DARTS           //
	FLOORBALL       //
	FOOTBALL        // SOCCER
	FOOTY           // SOCCER (Australian rules)
	FUTSAL          //
	GOLF            //
	GRIDIRON        // AMERICAL FOOTBALL
	HANDBALL        //
	HOCKEY          // ICE HOCKEY
	MMA             //
	MOTOR           //
	NETBALL         //
	PESAPALLO       //
	PINGPONG        //
	RUGBY           //
	RUGBYLEAGUE     //
	SNOOKER         //
	TENNIS          //
	VOLLEYBALL      //
	WATERPOLO       //
)

// Second name
const (
	FIELDHOCKEY = BANDY
	BEACHSOCCER = BEASAL
	SOCCER      = FOOTBALL
	AMFOOTBALL  = GRIDIRON
	FOOTBALLAM  = GRIDIRON
	AUFOOTBALL  = FOOTY
	FOOTBALLAU  = FOOTY
)

// Map english translation of sport.CONSTs
func En(t int) string {
	n, e := enstr[t]
	if !e {
		return "Unknown ID"
	}
	return n
}

// Const to english names
var enstr = map[int]string{
	BADMINTON:       "Badminton",
	BANDY:           "Bandy",
	BASEBALL:        "Baseball",
	BASKETBALL:      "Basketball",
	BEACHVOLLEYBALL: "Beach Volleyball",
	BEASAL:          "Beach Soccer",
	BICYCLE:         "Bicycle",
	BOXING:          "Boxing",
	CRICKET:         "Cricket",
	CYBER:           "Cyber",
	DARTS:           "Darts",
	FLOORBALL:       "Floorball",
	FOOTBALL:        "Soccer",
	FOOTY:           "Australian Rules",
	FUTSAL:          "Futsal",
	GOLF:            "Golf",
	GRIDIRON:        "Americal Football",
	HANDBALL:        "Handball",
	HOCKEY:          "Ice Hockey",
	MMA:             "MMA",
	MOTOR:           "Motor",
	NETBALL:         "Netball",
	PESAPALLO:       "Pesapallo",
	PINGPONG:        "Ping Pong",
	RUGBY:           "Rugby",
	RUGBYLEAGUE:     "Rugby League",
	SNOOKER:         "Snooker",
	TENNIS:          "Tennis",
	VOLLEYBALL:      "Volleyball",
	WATERPOLO:       "Water Polo",
}

// Russian translation of sport.CONSTs
func Ru(t int) string {
	n, e := rustr[t]
	if !e {
		return "Неизвестный ID"
	}
	return n
}

// Map const to russian names
var rustr = map[int]string{
	BADMINTON:       "Бадминтон",
	BANDY:           "Хоккей на траве",
	BASEBALL:        "Бейсбол",
	BASKETBALL:      "Баскетбол",
	BEACHVOLLEYBALL: "Пляжный волейбол",
	BEASAL:          "Пляжный футбол",
	BICYCLE:         "Велоспорт",
	BOXING:          "Бокс",
	CRICKET:         "Крикет",
	CYBER:           "Киберспорт",
	DARTS:           "Дартс",
	FLOORBALL:       "Флорбол",
	FOOTBALL:        "Футбол",
	FOOTY:           "Австралийский футбол",
	FUTSAL:          "Футзал",
	GOLF:            "Гольф",
	GRIDIRON:        "Американский футбол",
	HANDBALL:        "Гандбол",
	HOCKEY:          "Хоккей",
	MMA:             "ММА",
	MOTOR:           "Мотоспорт",
	NETBALL:         "Нетбол",
	PESAPALLO:       "Песапалло",
	PINGPONG:        "Пинг Понг",
	RUGBY:           "Рэгби",
	RUGBYLEAGUE:     "Рэгбилиг",
	SNOOKER:         "Снукер",
	TENNIS:          "Теннис",
	VOLLEYBALL:      "Волейбол",
	WATERPOLO:       "Водное поло",
}
