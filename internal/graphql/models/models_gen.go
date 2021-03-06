// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken" bson:"authToken"`
	User      *User      `json:"user" bson:"user"`
}

type AuthToken struct {
	AccessToken string    `json:"accessToken" bson:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt" bson:"expiredAt"`
}

type Cap struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Product string        `json:"product" bson:"product"`
	Desc    string        `json:"desc" bson:"desc"`
	Year    int           `json:"year" bson:"year"`
	Week    int           `json:"week" bson:"week"`
	Ctime   time.Time     `json:"ctime" bson:"ctime"`
	Utime   time.Time     `json:"utime" bson:"utime"`
}

type CapList struct {
	Code  int    `json:"code" bson:"code"`
	Data  []*Cap `json:"data" bson:"data"`
	Count int    `json:"count" bson:"count"`
}

type CreateCapInput struct {
	Product string `json:"product" bson:"product"`
	Desc    string `json:"desc" bson:"desc"`
}

type CreateFailureInput struct {
	StartTime time.Time `json:"start_time" bson:"start_time"`
	EndTime   time.Time `json:"end_time" bson:"end_time"`
	Duration  *int      `json:"duration" bson:"duration"`
	Product   string    `json:"product" bson:"product"`
	Desc      string    `json:"desc" bson:"desc"`
	Title     *string   `json:"title" bson:"title"`
	Recorder  string    `json:"recorder" bson:"recorder"`
	Level     string    `json:"level" bson:"level"`
}

type CreateOnlineCountInput struct {
	Product string `json:"product" bson:"product"`
	Online  int    `json:"online" bson:"online"`
}

type CreateSloInput struct {
	Product string                 `json:"product" bson:"product"`
	Metrics map[string]interface{} `json:"metrics" bson:"metrics"`
}

type CreateTicketInput struct {
	NormalLt2h   int `json:"normalLt2h" bson:"normalLt2h"`
	AbnormalLt2h int `json:"abnormalLt2h" bson:"abnormalLt2h"`
	NormalGt2h   int `json:"normalGt2h" bson:"normalGt2h"`
	AbnormalGt2h int `json:"abnormalGt2h" bson:"abnormalGt2h"`
}

type DeleteCap struct {
	Success bool `json:"success" bson:"success"`
}

type DeleteFailure struct {
	Success bool `json:"success" bson:"success"`
}

type DeleteOnlineCount struct {
	Success bool `json:"success" bson:"success"`
}

type DeleteSlo struct {
	Success bool `json:"success" bson:"success"`
}

type DeleteTicket struct {
	Success bool `json:"success" bson:"success"`
}

type EmailInput struct {
	From string `json:"from" bson:"from"`
	To   string `json:"to" bson:"to"`
	Data string `json:"data" bson:"data"`
}

type Failure struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	StartTime time.Time     `json:"start_time" bson:"start_time"`
	EndTime   time.Time     `json:"end_time" bson:"end_time"`
	Duration  float64       `json:"duration" bson:"duration"`
	// 业务线
	Product string `json:"product" bson:"product"`
	// 故障描述
	Desc string `json:"desc" bson:"desc"`
	// 故障标题
	Title *string `json:"title" bson:"title"`
	// 故障上报人
	Recorder string    `json:"recorder" bson:"recorder"`
	Level    string    `json:"level" bson:"level"`
	Week     int       `json:"week" bson:"week"`
	Year     int       `json:"year" bson:"year"`
	Ctime    time.Time `json:"ctime" bson:"ctime"`
	Utime    time.Time `json:"utime" bson:"utime"`
}

type FailureItem struct {
	Name string `json:"name" bson:"name"`
	Data []int  `json:"data" bson:"data"`
	Type string `json:"type" bson:"type"`
}

type FailureList struct {
	Code  int        `json:"code" bson:"code"`
	Data  []*Failure `json:"data" bson:"data"`
	Count int        `json:"count" bson:"count"`
}

type FailurePretty struct {
	Series []*FailureItem `json:"series" bson:"series"`
	XAxis  []string       `json:"xAxis" bson:"xAxis"`
	Year   int            `json:"year" bson:"year"`
	Week   int            `json:"week" bson:"week"`
}

type LoginInput struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type OnlineCount struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Ctime   time.Time     `json:"ctime" bson:"ctime"`
	Utime   time.Time     `json:"utime" bson:"utime"`
	Week    int           `json:"week" bson:"week"`
	Year    int           `json:"year" bson:"year"`
	Product string        `json:"product" bson:"product"`
	Online  int           `json:"online" bson:"online"`
}

type OnlineCountAllProduct struct {
	Online   []*int    `json:"online" bson:"online"`
	Products []*string `json:"products" bson:"products"`
	Year     int       `json:"year" bson:"year"`
	Week     int       `json:"week" bson:"week"`
}

type OnlineCountList struct {
	Data  []*OnlineCount `json:"data" bson:"data"`
	Count int            `json:"count" bson:"count"`
	Code  int            `json:"code" bson:"code"`
}

type Order struct {
	Normal       int `json:"normal" bson:"normal"`
	Abnormal     int `json:"abnormal" bson:"abnormal"`
	NormalLt2h   int `json:"normalLt2h" bson:"normalLt2h"`
	AbnormalLt2h int `json:"abnormalLt2h" bson:"abnormalLt2h"`
	NormalGt2h   int `json:"normalGt2h" bson:"normalGt2h"`
	AbnormalGt2h int `json:"abnormalGt2h" bson:"abnormalGt2h"`
}

type Slo struct {
	ID      bson.ObjectId          `json:"id" bson:"_id,omitempty"`
	Product string                 `json:"product" bson:"product"`
	Metrics map[string]interface{} `json:"metrics" bson:"metrics"`
	Year    int                    `json:"year" bson:"year"`
	Week    int                    `json:"week" bson:"week"`
	Ctime   time.Time              `json:"ctime" bson:"ctime"`
	Utime   time.Time              `json:"utime" bson:"utime"`
}

type SloList struct {
	Code  int    `json:"code" bson:"code"`
	Data  []*Slo `json:"data" bson:"data"`
	Count int    `json:"count" bson:"count"`
}

type SloPretty struct {
	Product string    `json:"product" bson:"product"`
	Sloes   []float64 `json:"sloes" bson:"sloes"`
	Legend  []string  `json:"legend" bson:"legend"`
	Year    int       `json:"year" bson:"year"`
	Week    int       `json:"week" bson:"week"`
}

type Ticket struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Order *Order        `json:"order" bson:"order"`
	Week  int           `json:"week" bson:"week"`
	Year  int           `json:"year" bson:"year"`
	Ctime time.Time     `json:"ctime" bson:"ctime"`
	Utime time.Time     `json:"utime" bson:"utime"`
}

type TicketEchartData struct {
	Name  string `json:"name" bson:"name"`
	Value int    `json:"value" bson:"value"`
}

type TicketList struct {
	Data  []*Ticket `json:"data" bson:"data"`
	Count int       `json:"count" bson:"count"`
	Code  int       `json:"code" bson:"code"`
}

type TicketPretty struct {
	Orders []*TicketEchartData `json:"orders" bson:"orders"`
	Legend []string            `json:"legend" bson:"legend"`
	Week   int                 `json:"week" bson:"week"`
	Year   int                 `json:"year" bson:"year"`
}

type UpdateCap struct {
	Success bool `json:"success" bson:"success"`
}

type UpdateCapInput struct {
	Product string `json:"product" bson:"product"`
	Desc    string `json:"desc" bson:"desc"`
}

type UpdateFailure struct {
	Success bool `json:"success" bson:"success"`
}

type UpdateFailureInput struct {
	StartTime time.Time `json:"start_time" bson:"start_time"`
	EndTime   time.Time `json:"end_time" bson:"end_time"`
	Duration  *int      `json:"duration" bson:"duration"`
	// 业务线
	Product string `json:"product" bson:"product"`
	// 故障描述
	Desc string `json:"desc" bson:"desc"`
	// 故障标题
	Title *string `json:"title" bson:"title"`
	// 故障上报人
	Recorder string `json:"recorder" bson:"recorder"`
	Level    string `json:"level" bson:"level"`
}

type UpdateOnlineCount struct {
	Success bool `json:"success" bson:"success"`
}

type UpdateOnlineCountInput struct {
	Product string `json:"product" bson:"product"`
	Online  int    `json:"online" bson:"online"`
}

type UpdateSlo struct {
	Success bool `json:"success" bson:"success"`
}

type UpdateSloInput struct {
	Product string                 `json:"product" bson:"product"`
	Metrics map[string]interface{} `json:"metrics" bson:"metrics"`
}

type UpdateTicket struct {
	Success bool `json:"success" bson:"success"`
}

type UpdateTicketInput struct {
	NormalLt2h   int `json:"normalLt2h" bson:"normalLt2h"`
	AbnormalLt2h int `json:"abnormalLt2h" bson:"abnormalLt2h"`
	NormalGt2h   int `json:"normalGt2h" bson:"normalGt2h"`
	AbnormalGt2h int `json:"abnormalGt2h" bson:"abnormalGt2h"`
}

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name"`
	Username string        `json:"username" bson:"username"`
	Location string        `json:"location" bson:"location"`
	Password string        `json:"password" bson:"password"`
	Abbr     string        `json:"abbr" bson:"abbr"`
	Email    string        `json:"email" bson:"email"`
	Openhab  string        `json:"openhab" bson:"openhab"`
}

type UserInput struct {
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Location string `json:"location" bson:"location"`
	Abbr     string `json:"abbr" bson:"abbr"`
	Email    string `json:"email" bson:"email"`
	Openhab  string `json:"openhab" bson:"openhab"`
}

type UserUpdate struct {
	Name     *string `json:"name" bson:"name"`
	Username *string `json:"username" bson:"username"`
	Password *string `json:"password" bson:"password"`
	Location *string `json:"location" bson:"location"`
	Abbr     *string `json:"abbr" bson:"abbr"`
	Email    *string `json:"email" bson:"email"`
	Openhab  *string `json:"openhab" bson:"openhab"`
}
