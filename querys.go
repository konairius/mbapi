package mbapi

import "fmt"

type Query interface {
	Path() string   //The Relative Query Path on the server
	Method() string //http Mathod i.e. GET, POST, DELETE, ...
}

type QueryField interface{}

type ListUsersQuery struct{}

func (l *ListUsersQuery) Path() string {
	return "/Users/Public"
}

func (l *ListUsersQuery) Method() string {
	return "GET"
}

type LoginQuery struct {
	Username QueryField `encoding:"para" type:"string" optional:"false"`
	Password QueryField `encoding:"para" type:"string" optional:"false"`
}

func (l *LoginQuery) Path() string {
	return "/Users/AuthenticateByName"
}

func (l *LoginQuery) Method() string {
	return "POST"
}

type ItemQuery struct {
	UserId                         QueryField `encoding:"path" type:"string" optional:"false"`
	Person                         QueryField `encoding:"para" type:"string" optional:"true"`
	PersonTypes                    QueryField `encoding:"para" type:"[]string" optional:"true" sep:","`
	Studios                        QueryField `encoding:"para" type:"[]string" optional:"true" sep:"|"`
	Artists                        QueryField `encoding:"para" type:"[]string" optional:"true" sep:"|"`
	Albums                         QueryField `encoding:"para" type:"[]string" optional:"true" sep:"|"`
	AllGenres                      QueryField `encoding:"para" type:"[]string" optional:"true" sep:"|"`
	Ids                            QueryField `encoding:"para" type:"[]string" optional:"true" sep:","`
	VideoTypes                     QueryField `encoding:"para" type:"[]string" optional:"true" sep:"," options:"videofile,dvd,bluray,iso"`
	Is3D                           QueryField `encoding:"para" type:"bool" optional:"true"`
	SeriesStatus                   QueryField `encoding:"para" type:"[]string" optional:"true" sep:","`
	NameStartsWithOrGreater        QueryField `encoding:"para" type:"string" optional:"true"`
	NameStartsWith                 QueryField `encoding:"para" type:"string" optional:"true"`
	NameLessThan                   QueryField `encoding:"para" type:"string" optional:"true"`
	AlbumArtistStartsWithOrGreater QueryField `encoding:"para" type:"string" optional:"true"`
	AirDays                        QueryField `encoding:"para" type:"[]time.Weekday" optional:"true" sep:","`
	MinOfficalRating               QueryField `encoding:"para" type:"string" optional:"true"`
	MaxOfficalRating               QueryField `encoding:"para" type:"string" optional:"true"`
	HasThemeSong                   QueryField `encoding:"para" type:"string" optional:"true"`
	HasThemeVideo                  QueryField `encoding:"para" type:"string" optional:"true"`
	HasSubtitles                   QueryField `encoding:"para" type:"string" optional:"true"`
	HasSpecialFeature              QueryField `encoding:"para" type:"string" optional:"true"`
	HasTrailer                     QueryField `encoding:"para" type:"string" optional:"true"`
	AdjacentTo                     QueryField `encoding:"para" type:"string" optional:"true"`
	MinIndexNumber                 QueryField `encoding:"para" type:"int" optional:"true"`
	MaxPlayers                     QueryField `encoding:"para" type:"int" optional:"true"`
	ParentIndexNumber              QueryField `encoding:"para" type:"int" optional:"true"`
	HasParentalRating              QueryField `encoding:"para" type:"bool" optional:"true"`
	IsHD                           QueryField `encoding:"para" type:"bool" optional:"true"`
	LocationTypes                  QueryField `encoding:"para" type:"[]string" optional:"true" sep:","`
	ExcludeLocationTypes           QueryField `encoding:"para" type:"[]string" optional:"true" sep:","`
	IsMissing                      QueryField `encoding:"para" type:"bool" optional:"true"`
	IsUnaired                      QueryField `encoding:"para" type:"bool" optional:"true"`
	IsVirtualUnaired               QueryField `encoding:"para" type:"bool" optional:"true"`
	MinCommunityRating             QueryField `encoding:"para" type:"int" optional:"true"`
	MinCriticRating                QueryField `encoding:"para" type:"int" optional:"true"`
	AiredDuringSeason              QueryField `encoding:"para" type:"int" optional:"true"`
	HasOverview                    QueryField `encoding:"para" type:"bool" optional:"true"`
	HasImdbId                      QueryField `encoding:"para" type:"bool" optional:"true"`
	HasTmdbId                      QueryField `encoding:"para" type:"bool" optional:"true"`
	HasTvdbId                      QueryField `encoding:"para" type:"bool" optional:"true"`
	IsYearMismatched               QueryField `encoding:"para" type:"bool" optional:"true"`
	IsInBoxSet                     QueryField `encoding:"para" type:"bool" optional:"true"`
	IsLocked                       QueryField `encoding:"para" type:"bool" optional:"true"`
	IsUnidentified                 QueryField `encoding:"para" type:"bool" optional:"true"`
	IsPlaceHolder                  QueryField `encoding:"para" type:"bool" optional:"true"`
	HasOfficialRating              QueryField `encoding:"para" type:"bool" optional:"true"`
	CollapseBoxSetItems            QueryField `encoding:"para" type:"bool" optional:"true"`
	StartIndex                     QueryField `encoding:"para" type:"int" optional:"true"`
	Limit                          QueryField `encoding:"para" type:"int" optional:"true"`
	Recursive                      QueryField `encoding:"para" type:"bool" optional:"true"`
	SortOrder                      QueryField `encoding:"para" type:"string" optional:"true" options:"Ascending,Descending"`
	ParentId                       QueryField `encoding:"para" type:"string" optional:"true"`
	Fields                         QueryField `encoding:"para" type:"[]string" optional:"true" sep:"," options:"Budget,Chapters,CriticRatingSummary,DateCreated,Genres,HomePageUrl,IndexOptions,MediaSources,MediaStreams,Overview,ParentId,Path,People,ProviderIds,PrimaryImageAspectRatio,Revenue,SortName,Studios,Taglines"`
	ExcludeItemTypes               QueryField `encoding:"para" type:"[]string" optional:"true" sep:","`
	IncludeItemTypes               QueryField `encoding:"para" type:"[]string" optional:"true" sep:","`
	Filters                        QueryField `encoding:"para" type:"[]string" optional:"true" sep:"," options:"IsFolder,IsNotFolder,IsUnplayed,IsPlayed,IsFavorite,IsResumable,Likes,Dislikes"`
	MediaTypes                     QueryField `encoding:"para" type:"[]string" optional:"true" sep:","`
	ImageTypes                     QueryField `encoding:"para" type:"[]string" optional:"true" sep:","`
	SortBy                         QueryField `encoding:"para" type:"[]string" optional:"true" sep:"," options:"Album,AlbumArtist,Artist,Budget,CommunityRating,CriticRating,DateCreated,DatePlayed,PlayCount,PremiereDate,ProductionYear,SortName,Random,Revenue,Runtime"`
	IsPlayed                       QueryField `encoding:"para" type:"bool" optional:"true"`
	Genres                         QueryField `encoding:"para" type:"[]string" optional:"true" sep:"|"`
	OfficialRatings                QueryField `encoding:"para" type:"[]string" optional:"true" sep:"|"`
	Tags                           QueryField `encoding:"para" type:"[]string" optional:"true" sep:"|"`
	Years                          QueryField `encoding:"para" type:"[]string" optional:"true" sep:","`
	EnableImages                   QueryField `encoding:"para" type:"bool" optional:"true"`
	ImageTypeLimit                 QueryField `encoding:"para" type:"int" optional:"true"`
	EnableImageTypes               QueryField `encoding:"para" type:"[]string" optional:"true" sep:","`
}

func (i *ItemQuery) Method() string {
	return "GET"
}

func (i *ItemQuery) Path() string {
	return "/Users/{{UserId}}/Items"
}

type VideoStreamQuery struct {
	Container string

	VideoCodec             QueryField `encoding:"para" type:"string" optional:"true" sep:"," options:"h264,mpeg4,theora,vpx,wmv"`
	VideoBitRate           QueryField `encoding:"para" type:"int" optional:"true"`
	AudioStreamIndex       QueryField `encoding:"para" type:"int" optional:"true"`
	VideoStreamIndex       QueryField `encoding:"para" type:"int" optional:"true"`
	SubtitleStreamIndex    QueryField `encoding:"para" type:"int" optional:"true"`
	Width                  QueryField `encoding:"para" type:"int" optional:"true"`
	Height                 QueryField `encoding:"para" type:"int" optional:"true"`
	MaxWidth               QueryField `encoding:"para" type:"int" optional:"true"`
	MaxHeight              QueryField `encoding:"para" type:"int" optional:"true"`
	MaxRefFrames           QueryField `encoding:"para" type:"int" optional:"true"`
	MaxVideoBitDepth       QueryField `encoding:"para" type:"int" optional:"true"`
	Framerate              QueryField `encoding:"para" type:"float64" optional:"true"`
	MaxFramerate           QueryField `encoding:"para" type:"float64" optional:"true"`
	Profile                QueryField `encoding:"para" type:"string" optional:"true"`
	Level                  QueryField `encoding:"para" type:"string" optional:"true"`
	SubtitleDeliveryMethod QueryField `encoding:"para" type:"string" optional:"true"`
	EnableAutoStreamCopy   QueryField `encoding:"para" type:"bool" optional:"true"`
	Cabac                  QueryField `encoding:"para" type:"bool" optional:"true"`
	Id                     QueryField `encoding:"path" type:"string" optional:"false"`
	MediaSourceId          QueryField `encoding:"para" type:"string" optional:"false"`
	DeviceId               QueryField `encoding:"para" type:"string" optional:"true"`
	AudioCodec             QueryField `encoding:"para" type:"string" optional:"true" sep:"," options:"aac,mp3,vorbis,wma"`
	StartTimeTicks         QueryField `encoding:"para" type:"int" optional:"true"`
	AudioBitRate           QueryField `encoding:"para" type:"int" optional:"true"`
	AudioChannels          QueryField `encoding:"para" type:"int" optional:"true"`
	AudioSampleRate        QueryField `encoding:"para" type:"int" optional:"true"`
	Static                 QueryField `encoding:"para" type:"bool" optional:"true"`
	DeviceProfileId        QueryField `encoding:"para" type:"string" optional:"true"`
}

func (v *VideoStreamQuery) Method() string {
	return "GET"
}

func (v *VideoStreamQuery) Path() string {
	if v.Container != "" {
		return fmt.Sprintf("/Videos/{{Id}}/stream.%s", v.Container)
	}
	return "/Videos/{{Id}}/stream"
}
