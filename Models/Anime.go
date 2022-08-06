package models

import (
	config "animerest/Config"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

type UserAnime struct {
	Id     int                     `pg:",pk"`
	Site   config.ApiConfiguration `pg:",unique"`
	UserId uuid.UUID               `pg:",pk,type:uuid"`
	User   *User                   `pg:"rel:has-one"`
}

type Anime struct {
	Id          int           `json:"id"`
	Series      string        `json:"series"`
	CreatedDate time.Duration `json:"createdDate"`
	Code        string        `json:"code"`
	Year        string        `json:"year"`
	Poster      string        `json:"poster"`
	Names       AnimeNames    `json:"names"`
	Description string        `json:"description"`
	Team        *AnimeTeam    `json:"team"`
	Voiceover   string        `json:"voiceOver"`
	Genres      []string      `json:"genres"`
}

func (u *Anime) UnmarshalJSON(data []byte) error {
	var aux Anilibria

	err := json.Unmarshal(data, &aux)

	if err != nil {
		return err
	} else {
		u.Id = aux.Id
		u.Code = aux.Code
		u.CreatedDate = aux.LastChange
		u.Description = aux.Description
		u.Genres = aux.Genres
		u.Names = AnimeNames{
			Ru: aux.Names.Ru,
			En: aux.Names.En,
		}
		if len(aux.Player.Playlist) > 0 {
			for k := range aux.Player.Playlist {
				u.Series = k + "-" + strconv.Itoa(len(aux.Player.Playlist))
				break
			}
		}
		u.Poster = aux.Posters.Origins.Url
		u.Team = &AnimeTeam{
			Voice:      aux.Team.Voice,
			Editing:    aux.Team.Editing,
			Decor:      aux.Team.Decor,
			Translator: aux.Team.Translator,
		}
		u.Year = strconv.Itoa(aux.Season.Year)
	}

	return nil
}

type AnimeSkips struct {
	Opening string `json:"opening"`
	Ending  string `json:"ending"`
}

type AnimeNames struct {
	Ru string `json:"ru"`
	Kz string `json:"kz"`
	En string `json:"en"`
}

type AnimeTeam struct {
	Voice      []string `json:"voice"`
	Editing    []string `json:"editing"`
	Decor      []string `json:"decor"`
	Translator []string `json:"translator"`
}

type Anilibria struct {
	Announce interface{} `json:"announce"`
	Blocked  struct {
		Bakanim bool `json:"bakanim"`
		Blocked bool `json:"blocked"`
	} `json:"blocked"`
	Code        string        `json:"code"`
	Description string        `json:"description"`
	Genres      []string      `json:"genres"`
	Id          int           `json:"id"`
	InFavorites int           `json:"in_favorites"`
	LastChange  time.Duration `json:"last_change"`
	Names       struct {
		Alternative string `json:"alternative"`
		En          string `json:"en"`
		Ru          string `json:"ru"`
	} `json:"names"`
	AlternativePlayer string `json:"alternative_player"`
	Host              string `json:"host"`
	Player            struct {
		Playlist map[string]struct {
			CreatedTimestamp time.Duration `json:"created_timestamp"`
			Hls              struct {
				Fhd string `json:"fhd"`
				Hd  string `json:"hd"`
				Sd  string `json:"sd"`
			} `json:"hls"`
			Preview string `json:"preview"`
			Serie   int    `json:"serie"`
			Skips   struct {
				Ending  []int `json:"ending"`
				Opening []int `json:"opening"`
			} `json:"skips"`
		} `json:"playlist"`
		Series struct {
			First  int    `json:"first"`
			Last   int    `json:"last"`
			String string `json:"string"`
		} `json:"series"`
	} `json:"player"`
	Posters struct {
		Medium struct {
			RawBase64File interface{} `json:"row_base64_file"`
			Url           string      `json:"string"`
		} `json:"medium"`
		Origins struct {
			RawBase64File interface{} `json:"row_base64_file"`
			Url           string      `json:"string"`
		} `json:"origins"`
		Small struct {
			RawBase64File interface{} `json:"row_base64_file"`
			Url           string      `json:"string"`
		} `json:"small"`
	} `json:"posters"`
	Season struct {
		Code    int8   `json:"code"`
		String  string `json:"string"`
		WeekDay int8   `json:"week_day"`
		Year    int    `json:"year"`
	} `json:"season"`
	Status struct {
		Code   int8   `json:"code"`
		String string `json:"string"`
	} `json:"status"`
	Team struct {
		Decor      []string `json:"decor"`
		Editing    []string `json:"editing"`
		Timing     []string `json:"timing"`
		Translator []string `json:"translator"`
		Voice      []string `json:"voice"`
	} `json:"team"`
	Torrents struct {
		List []struct {
			Downloads int8        `json:"download"`
			Hash      string      `json:"hash"`
			Leechers  int8        `json:"leechers"`
			Metadata  interface{} `json:"metadata"`
			Quality   struct {
				Encoder    string      `json:"encoder"`
				LqAudio    interface{} `json:"lq_audio"`
				Resolution string      `json:"resolution"`
				String     string      `json:"string"`
				Type       string      `json:"type"`
			} `json:"quality"`
			RawBase64File interface{} `json:"row_base64_file"`
			Seeders       int         `json:"seeders"`
			Series        struct {
				First  int    `json:"series"`
				Last   int    `json:"last"`
				String string `json:"string"`
			} `json:"series"`
			TorrentId         int           `json:"torrent_id"`
			TotalSize         int           `json:"total_size"`
			UploadedTimestamp time.Duration `json:"uploaded_timestamp"`
			Url               string        `json:"url"`
		} `json:"list"`
		Type struct {
			Code       int    `json:"code"`
			FullString string `json:"full_string"`
			Length     int    `json:"length"`
			Series     int    `json:"series"`
			String     string `json:"string"`
		} `json:"type"`
	}
	Updated time.Duration `json:"updated"`
}

type AnimeFilter struct {
	Id   int                     `json:"id"`
	Code string                  `json:"code"`
	T    config.ApiConfiguration `jsoN:"t"`
}
