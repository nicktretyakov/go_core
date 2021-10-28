package models

import (
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

// User struct
type User struct {
	ID                    int64           `json:"id"`
	Name                  string  `json:"name"`
	Username              string  `json:"username"`
	Balance               float64         `json:"balance"`
	Avatar                string  `json:"avatar"`
	//Country               string  `json:"country"`
	Cover                 string  `json:"cover"`
	About                 string  `json:"about"`
	Url                   string  `json:"url"`
	Wishlist              string  `json:"wishlist"`
	Email                 string          `json:"email"`
	EmailVerifiedAt       time.Time    `json:"email_verified_at"`
	Password              string          `json:"-"`
	RememberToken         string  `json:"remember_token"`
	CreatedAt                 time.Time    `json:"created_at"`
	UpdatedAt             time.Time    `json:"updated_at"`
	FollowPrice           float64 `json:"follow_price"`
	FreeDays              bool    `json:"free_days"`
	ButtonText            string  `json:"button_text"`
	IsFirstTime           bool            `json:"is_first_time"`
	CloudToken            string  `json:"cloud_token"`
	Adult                 bool            `json:"adult"`
	SocID                 string  `json:"soc_id"`
	Phone                 string  `json:"phone"`
	HideFols              bool            `json:"hide_fols"`
	ProfileType           int32           `json:"profile_type"`
	ReferrerID            int64   `json:"referrer_id"`
	SubscriberUid         string  `json:"subscriber_uid"`
	PushType              int32   `json:"push_type"`
	Verified              bool            `json:"verified"`
	AllowMassSend         bool            `json:"allow_mass_send"`
	AllowGreeting         bool            `json:"allow_greeting"`
	GreetingText          string  `json:"greeting_text"`
	AuthToken             string  `json:"-"`
	PushWeb               bool            `json:"push_web"`
	PushEmail             bool            `json:"push_email"`
	PushTelegram          bool            `json:"push_telegram"`
	TelegramStarted       bool            `json:"telegram_started"`
	Percent               int32   `json:"percent"`
	LastOnlineAt          time.Time       `json:"last_online_at"`
	Ban                   bool            `json:"ban"`
	EmailConfirmationCode string  `json:"email_confirmation_code"`
	Ip                    string  `json:"ip"`
	PercentAsRef          int32   `json:"percent_as_ref"`
	ServicePercentAsRef   int32   `json:"service_percent_as_ref"`
	ShowVideo             bool            `json:"show_video"`
	CountryID             int32   `json:"country_id"`
	Lang                  string          `json:"lang"`
	OnlyOur               bool            `json:"only_our"`
	Currency              string          `json:"currency"`
	RefsFreeMonth         bool            `json:"refs_free_month"`
	FollowersCount        int32           `json:"followers_count"`
	VideosCount           int32           `json:"videos_count"`
	FollowingsCount       int32           `json:"followings_count"`
	UnreadMessages        int32           `json:"unread_messages"`
	FriendsCount          int32           `json:"friends_count"`
	LikesCount            int32           `json:"likes_count"`
	IsReferrer            bool            `json:"is_referrer"`
	HasCards              bool            `json:"has_cards"`
	EnteredLink           int64   `json:"entered_link"`
	RestrictedForRf       bool            `json:"restricted_for_rf"`
	ProfileViews          int64           `json:"profile_views"`
	MinWithdrawal         string  `json:"min_withdrawal"`
}

func (user *User) AfterCreate(tx *gorm.DB) (err error) {
	user.Username = strconv.FormatInt(user.ID, 10)
	user.Name = strconv.FormatInt(user.ID, 10)
	if !strings.Contains(user.Email, "@") {
		user.Email = strconv.FormatInt(user.ID, 10) + "@me"
	}
	user.CountryID = 0
	user.Adult = true
	user.FollowPrice = 6.66
	user.CountryID = 0
	user.Currency = "USD"
	user.Lang = "RU"
	user.Ip = ""

	tx.Save(user)
	return
}
