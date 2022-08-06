package models

import (
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
)

type User struct {
	Id           uuid.UUID         `pg:",pk,type:uuid"`
	Uid          int               `pg:",unique" json:"uid"`
	Username     string            `pg:",unique" json:"username"`
	FirstName    string            `pg:",notnull" json:"firstName"`
	LastName     string            `pg:",notnull" json:"lastName"`
	MiddleName   string            `json:"middleName"`
	Rating       int8              `json:"rating"`
	Favorites    []UserFavorite    `pg:",rel:has-many" json:"favorites"`
	WatchLater   []UserWatchLater  `pg:",rel:has-many" json:"watchLater"`
	Achievements []UserAchievement `pg:",rel:has-many" json:"achievements"`
	Email        string            `pg:",unique" json:"email"`
	Friends      []UserFriend      `json:"friends"`
	Role         UserRole          `json:"role"`
	Messages     []UserMessage     `json:"messages"`
	Password     string            `json:"-"`
}

type UserChat struct {
	Id       uuid.UUID     `pg:",pk,type:uuid"`
	UserId   uuid.UUID     `pg:",type:uuid"`
	User     *User         `pg:",rel:has-one"`
	Messages []UserMessage `pg:",rel:has-many"`
}

type UserMessage struct {
	Id      uuid.UUID `pg:",pk,type:uuid"`
	UserId  uuid.UUID `pg:",type:uuid"`
	User    *User     `pg:",rel:has-one" json:"user"`
	Message string    `json:"message"`
	ChatId  uuid.UUID `pg:",type:uuid"`
	Chat    *UserChat `pg:",rel:has-one" json:"chat"`
}

type UserFriend struct {
	UserId   uuid.UUID          `pg:",pk,type:uuid"`
	User     *User              `pg:",rel:has-one"`
	FriendId uuid.UUID          `pg:",pk,type:uuid"`
	Friend   *User              `pg:",rel:has-one"`
	Relation UserFriendRelation `json:"relation"`
}

type UserAchievement struct {
	UserId uuid.UUID `pg:",pk,type:uuid"`
	User   *User     `pg:",rel:has-one"`
}

type UserFavorite struct {
	UserId  uuid.UUID `pg:",pk,type:uuid"`
	User    *User     `pg:",rel:has-one"`
	AnimeId uuid.UUID `pg:",pk,type:uuid"`
	Anime   *Anime    `pg:",anime"`
}

type UserWatchLater struct {
	UserId  uuid.UUID `pg:",pk,type:uuid"`
	User    *User     `pg:",rel:has-one"`
	AnimeId uuid.UUID `pg:",pk,type:uuid"`
	Anime   *Anime    `pg:",anime"`
}

type UserFriendRelation = string

const (
	New        UserFriendRelation = "new"
	BestFriend UserFriendRelation = "bestFriend"
	Boyfriend  UserFriendRelation = "boyfriend"
	Girlfriend UserFriendRelation = "girlfriend"
)

type UserRole = string

const (
	Admin   UserRole = "admin"
	Member  UserRole = "member"
	Donater UserRole = "donater"
)

// ----------------------------------------------------------------

type UserLoginDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserRegisterDto struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	MiddleName      string `json:"middleName"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UserClaims struct {
	jwt.StandardClaims
	Uid int `json:"uid"`
}
