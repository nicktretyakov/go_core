// Code generated by sqlc. DO NOT EDIT.

package models

import (
	"fmt"
	"time"
)

type ArticlesStatus string

const (
	ArticlesStatusPUBLISHED ArticlesStatus = "PUBLISHED"
	ArticlesStatusDRAFT     ArticlesStatus = "DRAFT"
)

func (e *ArticlesStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ArticlesStatus(s)
	case string:
		*e = ArticlesStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ArticlesStatus: %T", src)
	}
	return nil
}

type AppsCountriesDetailed struct {
	ID            int32          `json:"id"`
	Countrycode   string         `json:"countrycode"`
	Countryname   string         `json:"countryname"`
	Currencycode  string `json:"currencycode"`
	Fipscode      string `json:"fipscode"`
	Isonumeric    string `json:"isonumeric"`
	North         string `json:"north"`
	South         string `json:"south"`
	East          string `json:"east"`
	West          string `json:"west"`
	Capital       string `json:"capital"`
	Continentname string `json:"continentname"`
	Continent     string `json:"continent"`
	Languages     string `json:"languages"`
	Isoalpha3     string `json:"isoalpha3"`
	Geonameid     int32  `json:"geonameid"`
}

type Article struct {
	ID int32 `json:"id"`
	// (DC2Type:json)
	Title string `json:"title"`
	Slug  string `json:"slug"`
	// (DC2Type:json)
	Content   string         `json:"content"`
	Image     string         `json:"image"`
	Status    ArticlesStatus `json:"status"`
	Date      time.Time      `json:"date"`
	Featured  bool           `json:"featured"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt time.Time      `json:"deleted_at"`
	Visits    int64          `json:"visits"`
}

type ArticleTag struct {
	ID        int32        `json:"id"`
	ArticleID int32        `json:"article_id"`
	TagID     int32        `json:"tag_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Ban struct {
	ID          int64          `json:"id"`
	Description string `json:"description"`
	Level       int32  `json:"level"`
	UserID      int64          `json:"user_id"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Card struct {
	ID            int64          `json:"id"`
	UserID        int64          `json:"user_id"`
	First6        string         `json:"first6"`
	Last4         string         `json:"last4"`
	ExpiryMonth   string         `json:"expiry_month"`
	ExpiryYear    string         `json:"expiry_year"`
	CardType      string `json:"card_type"`
	IssuerCountry string `json:"issuer_country"`
	IssuerName    string `json:"issuer_name"`
	Token         string         `json:"token"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	Main          bool           `json:"main"`
}

type Category struct {
	ID        int32         `json:"id"`
	ParentID  int32 `json:"parent_id"`
	Lft       int32 `json:"lft"`
	Rgt       int32 `json:"rgt"`
	Depth     int32 `json:"depth"`
	Name      string        `json:"name"`
	Slug      string        `json:"slug"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}

type Chat struct {
	ID            int64         `json:"id"`
	LastMessageID int64 `json:"last_message_id"`
	CreatedAt                 time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type Comment struct {
	ID              int64        `json:"id"`
	CommentableType string       `json:"commentable_type"`
	CommentableID   int64        `json:"commentable_id"`
	CommenterType   string       `json:"commenter_type"`
	CommenterID     int64        `json:"commenter_id"`
	Comment         string       `json:"comment"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Complaint struct {
	ID        int64          `json:"id"`
	UserID    int64          `json:"user_id"`
	SenderID  int64          `json:"sender_id"`
	Reason    int32          `json:"reason"`
	Status    int32          `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Image     string `json:"image"`
}

type ComplaintReason struct {
	ID int64 `json:"id"`
	// (DC2Type:json)
	Reason    string       `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeviceSession struct {
	ID        int64        `json:"id"`
	Uuid      string       `json:"uuid"`
	UserID    int64        `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Dislike struct {
	ID int64 `json:"id"`
	// user_id
	UserID          int64        `json:"user_id"`
	DislikeableType string       `json:"dislikeable_type"`
	DislikeableID   int64        `json:"dislikeable_id"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type DmcaComplaint struct {
	ID          int64        `json:"id"`
	Url         string       `json:"url"`
	Email       string       `json:"email"`
	Description string       `json:"description"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FailedJob struct {
	ID         int64     `json:"id"`
	Connection string    `json:"connection"`
	Queue      string    `json:"queue"`
	Payload    string    `json:"payload"`
	Exception  string    `json:"exception"`
	FailedAt   time.Time `json:"failed_at"`
}

type Faq struct {
	ID      int64 `json:"id"`
	Type    int32 `json:"type"`
	ThemeID int64 `json:"theme_id"`
	// (DC2Type:json)
	Name string `json:"name"`
	// (DC2Type:json)
	Text      string       `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FaqTheme struct {
	ID int64 `json:"id"`
	// (DC2Type:json)
	Name      string       `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Gift struct {
	ID        int64        `json:"id"`
	UserID    int64        `json:"user_id"`
	Name      string       `json:"name"`
	Phone     string       `json:"phone"`
	Address   string       `json:"address"`
	Gift      int32        `json:"gift"`
	Status    int32        `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Image struct {
	ID        int64           `json:"id"`
	UserID    int64           `json:"user_id"`
	Src       string  `json:"src"`
	Blured    string  `json:"blured"`
	Price     float64 `json:"price"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type Like struct {
	ID int64 `json:"id"`
	// user_id
	UserID       int64        `json:"user_id"`
	LikeableType string       `json:"likeable_type"`
	LikeableID   int64        `json:"likeable_id"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Link struct {
	ID        int64        `json:"id"`
	Link      string       `json:"link"`
	Name      string       `json:"name"`
	UserID    int64        `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LinkVisit struct {
	ID       int64        `json:"id"`
	LinkID   int64        `json:"link_id"`
	Visits   int64        `json:"visits"`
	HitsDate time.Time `json:"hits_date"`
}

type Message struct {
	ID          int64          `json:"id"`
	ChatID      int64          `json:"chat_id"`
	SenderID    int64          `json:"sender_id"`
	Message     string `json:"message"`
	ReplyTo     int64  `json:"reply_to"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	ReadAt      time.Time   `json:"read_at"`
	Type        int32          `json:"type"`
	RelatedType string `json:"related_type"`
	RelatedID   int64  `json:"related_id"`
	Paid        bool           `json:"paid"`
	IsGreeting  bool           `json:"is_greeting"`
}

type Migration struct {
	ID        int32  `json:"id"`
	Migration string `json:"migration"`
	Batch     int32  `json:"batch"`
}

type ModelHasPermission struct {
	PermissionID int64  `json:"permission_id"`
	ModelType    string `json:"model_type"`
	ModelID      int64  `json:"model_id"`
}

type ModelHasRole struct {
	RoleID    int64  `json:"role_id"`
	ModelType string `json:"model_type"`
	ModelID   int64  `json:"model_id"`
}

type Money struct {
	ID             int64          `json:"id"`
	ToID           int64          `json:"to_id"`
	Amount         float64        `json:"amount"`
	OurAmount      float64        `json:"our_amount"`
	Reason         int32          `json:"reason"`
	PaymentSystem  int32          `json:"payment_system"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	Status         int32          `json:"status"`
	ReferrerAmount string `json:"referrer_amount"`
	Uuid           string `json:"uuid"`
	CardID         int64  `json:"card_id"`
	FromID         int64          `json:"from_id"`
	BalanceAfter   string         `json:"balance_after"`
}

type Notification struct {
	ID             string       `json:"id"`
	Type           string       `json:"type"`
	NotifiableType string       `json:"notifiable_type"`
	NotifiableID   int64        `json:"notifiable_id"`
	Data           string       `json:"data"`
	ReadAt         time.Time `json:"read_at"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type OldWithdrawal struct {
	ID                int64          `json:"id"`
	UserID            int64          `json:"user_id"`
	Amount            string `json:"amount"`
	Status            int32          `json:"status"`
	PaidAt            time.Time   `json:"paid_at"`
	Method            string `json:"method"`
	Details           string `json:"details"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	Fio               string `json:"fio"`
	UserBill          string `json:"user_bill"`
	BankBik           string `json:"bank_bik"`
	CorrespondentBill string `json:"correspondent_bill"`
	Inn               string `json:"inn"`
	Kpp               string `json:"kpp"`
	IbanBsb           string `json:"iban_bsb"`
	SwiftAba          string `json:"swift_aba"`
	CountryID         int64  `json:"country_id"`
	Addres            string `json:"addres"`
}

type PaidVideo struct {
	ID        int64        `json:"id"`
	UserID    int64        `json:"user_id"`
	VideoID   int64        `json:"video_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Participant struct {
	ID        int64        `json:"id"`
	UserID    int64        `json:"user_id"`
	ChatID    int64        `json:"chat_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PasswordReset struct {
	Email     string       `json:"email"`
	Token     string       `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

type Payment struct {
	ID          int64           `json:"id"`
	UserID      int64           `json:"user_id"`
	FollowingID int64   `json:"following_id"`
	Amount      float64 `json:"amount"`
	Currency    string          `json:"currency"`
	Type        int32   `json:"type"`
	Status      int32           `json:"status"`
	Description string  `json:"description"`
	Uuid        string  `json:"uuid"`
	Recurrent   bool            `json:"recurrent"`
	CreatedAt                 time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type Permission struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	GuardName string       `json:"guard_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type QuestionTheme struct {
	ID        int64        `json:"id"`
	Theme     string       `json:"theme"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ReserveImage struct {
	ID        int64           `json:"id"`
	UserID    int64           `json:"user_id"`
	Src       string  `json:"src"`
	Blured    string  `json:"blured"`
	Price     float64 `json:"price"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type ReserveVideo struct {
	ID                        int64           `json:"id"`
	UserID                    int64           `json:"user_id"`
	Privacy                   int32           `json:"privacy"`
	Top                       bool            `json:"top"`
	Price                     float64 `json:"price"`
	Title                     string  `json:"title"`
	Description               string  `json:"description"`
	OriginalName              string  `json:"original_name"`
	Thumbnail                 string  `json:"thumbnail"`
	Disk                      string  `json:"disk"`
	Path                      string          `json:"path"`
	ConvertedForDownloadingAt time.Time    `json:"converted_for_downloading_at"`
	ConvertedForStreamingAt   time.Time    `json:"converted_for_streaming_at"`
	CreatedAt                 time.Time    `json:"created_at"`
	UpdatedAt                 time.Time    `json:"updated_at"`
	SharesCount               int32           `json:"shares_count"`
	Duration                  int32   `json:"duration"`
	Ban                       bool            `json:"ban"`
}

type Role struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	GuardName string       `json:"guard_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleHasPermission struct {
	PermissionID int64 `json:"permission_id"`
	RoleID       int64 `json:"role_id"`
}

type RussianRequisite struct {
	ID                int64          `json:"id"`
	Fio               string `json:"fio"`
	UserBill          string `json:"user_bill"`
	BankBik           string `json:"bank_bik"`
	CorrespondentBill string `json:"correspondent_bill"`
	Inn               string `json:"inn"`
	Kpp               string `json:"kpp"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	UserID            int64          `json:"user_id"`
	DeletedAt         time.Time   `json:"deleted_at"`
}

type Tag struct {
	ID int32 `json:"id"`
	// (DC2Type:json)
	Name      string       `json:"name"`
	Slug      string       `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type TempSetting struct {
	ID        int64        `json:"id"`
	Key       string       `json:"key"`
	Value     string       `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Transaction struct {
	ID         int64          `json:"id"`
	UserID     int64          `json:"user_id"`
	ReceiverID int64          `json:"receiver_id"`
	Amount     float64        `json:"amount"`
	Status     int32          `json:"status"`
	Currency   string         `json:"currency"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Desc       string `json:"desc"`
}


type UserFollower struct {
	ID             int32           `json:"id"`
	FollowingID    int64           `json:"following_id"`
	FollowerID     int64           `json:"follower_id"`
	AcceptedAt     time.Time    `json:"accepted_at"`
	CreatedAt                 time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	Price          float64 `json:"price"`
	SubscriptionID string  `json:"subscription_id"`
	Active         bool            `json:"active"`
	NextPaymentAt  time.Time    `json:"next_payment_at"`
	DebtDays       int32           `json:"debt_days"`
	DeletedAt      time.Time    `json:"deleted_at"`
	OnlyOur        bool            `json:"only_our"`
	Trial          bool            `json:"trial"`
}

type UserQuestion struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Theme     int64        `json:"theme"`
	Message   string       `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Vacancy struct {
	ID            int64        `json:"id"`
	Name          string       `json:"name"`
	Duty          string       `json:"duty"`
	Qualification string       `json:"qualification"`
	Link          string       `json:"link"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Verification struct {
	ID           int64          `json:"id"`
	UserID       int64          `json:"user_id"`
	Selfie       string `json:"selfie"`
	Doc          string `json:"doc"`
	Status       int32          `json:"status"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	RandomPhoto  string `json:"random_photo"`
	RandomNumber int32  `json:"random_number"`
	Description  string `json:"description"`
}

type Video struct {
	ID                        int64           `json:"id"`
	UserID                    int64           `json:"user_id"`
	Privacy                   int32           `json:"privacy"`
	Top                       bool            `json:"top"`
	Price                     float64 `json:"price"`
	Title                     string  `json:"title"`
	Description               string  `json:"description"`
	OriginalName              string  `json:"original_name"`
	Thumbnail                 string  `json:"thumbnail"`
	Disk                      string  `json:"disk"`
	Path                      string          `json:"path"`
	ConvertedForDownloadingAt time.Time    `json:"converted_for_downloading_at"`
	ConvertedForStreamingAt   time.Time    `json:"converted_for_streaming_at"`
	CreatedAt                 time.Time    `json:"created_at"`
	UpdatedAt                 time.Time    `json:"updated_at"`
	SharesCount               int32           `json:"shares_count"`
	Duration                  int32   `json:"duration"`
	Ban                       bool            `json:"ban"`
	ThumbnailID               int64   `json:"thumbnail_id"`
	Views                     int64           `json:"views"`
	LikesCount                int32           `json:"likes_count"`
	Src                       string  `json:"src"`
	RestrictedForRf           bool            `json:"restricted_for_rf"`
}

type Withdrawal struct {
	ID             int64         `json:"id"`
	Amount         string        `json:"amount"`
	UserID         int64         `json:"user_id"`
	RequisiteID    int64         `json:"requisite_id"`
	RequisiteType  int32         `json:"requisite_type"`
	Status         int32         `json:"status"`
	PaidBy         int64 `json:"paid_by"`
	PaidAt         time.Time  `json:"paid_at"`
	CreatedAt                 time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      time.Time  `json:"deleted_at"`
	StartPayment   int64         `json:"start_payment"`
	EndPayment     int64         `json:"end_payment"`
	FullAmount     string        `json:"full_amount"`
	BalanceAfter   string        `json:"balance_after"`
	OriginalAmount string        `json:"original_amount"`
	Type           int32         `json:"type"`
}

type WorldRequisite struct {
	ID                 int64          `json:"id"`
	Fio                string `json:"fio"`
	IbanBsb            string `json:"iban_bsb"`
	SwiftAba           string `json:"swift_aba"`
	Address            string `json:"address"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	UserID             int64          `json:"user_id"`
	DeletedAt          time.Time   `json:"deleted_at"`
	CorrespondentSwift string `json:"correspondent_swift"`
	CorrespondentBill  string `json:"correspondent_bill"`
}
