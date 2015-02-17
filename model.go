package mbapi

import "time"

type Response interface{}

/////////////////
//Item Response//
/////////////////

type ItemsResult struct {
	Items            []BaseItemDto
	TotalRecordCount int
}

func (i *ItemsResult) Type() string {
	return "ItemsResult"
}

type BaseItemDto struct {
	Name                            string
	ServerId                        string
	Id                              string
	PlaylistItemId                  string
	DateCreated                     time.Time ////,optional
	DateLastMediaAdded              time.Time ////,optional
	ExtraType                       string    ////,optional) = ['Clip' or 'Trailer' or 'BehindTheScenes' or 'DeletedScene' or 'Interview' or 'Scene' or 'Sample' or 'ThemeSong' or 'ThemeVideo'],
	AirsBeforeSeasonNumber          int       ////,optional
	AirsAfterSeasonNumber           int       ////,optional
	AirsBeforeEpisodeNumber         int       ////,optional
	AbsoluteEpisodeNumber           int       ////,optional
	DisplaySpecialsWithSeasons      bool      ////,optional
	PreferredMetadataLanguage       string
	PreferredMetadataCountryCode    string
	AwardSummary                    string
	Metascore                       float32 //,optional
	IsUnidentified                  bool    //,optional
	AnimeSeriesIndex                int     //,optional
	SupportsSync                    bool    //,optional
	DvdSeasonNumber                 int     //,optional
	DvdEpisodeNumber                float32 //,optional
	SortName                        string
	ForcedSortName                  string
	Video3DFormat                   string    //,optional) = ['HalfSideBySide' or 'FullSideBySide' or 'FullTopAndBottom' or 'HalfTopAndBottom'],
	PremiereDate                    time.Time //,optional
	ExternalUrls                    []ExternalUrl
	MediaSources                    []MediaSourceInfo
	CriticRating                    float32 //,optional
	GameSystem                      string
	CriticRatingSummary             string
	MultiPartGameFiles              []string
	Path                            string
	OfficialRating                  string
	CustomRating                    string
	ChannelId                       string
	ChannelName                     string
	Overview                        string
	ShortOverview                   string
	TmdbCollectionName              string
	Taglines                        []string
	Genres                          []string
	SeriesGenres                    []string
	CommunityRating                 float32 //,optional
	VoteCount                       int     //,optional
	CumulativeRunTimeTicks          int64   //,optional
	OriginalRunTimeTicks            int64   //,optional
	RunTimeTicks                    int64   //,optional
	RecursiveUnplayedItemCount      int     //,optional
	PlayAccess                      string  //) = ['Full' or 'None'],
	AspectRatio                     string
	ProductionYear                  int  //,optional
	SeasonCount                     int  //,optional
	Players                         int  //,optional
	IsPlaceHolder                   bool //,optional
	IndexNumber                     int  //,optional
	IndexNumberEnd                  int  //,optional
	ParentIndexNumber               int  //,optional
	RemoteTrailers                  []MediaUrl
	SoundtrackIds                   []string
	IsHD                            bool //,optional
	IsFolder                        bool
	ParentId                        string
	Type                            string
	People                          []BaseItemPerson
	Studios                         []StudioDto
	ParentLogoItemId                string
	ParentBackdropItemId            string
	ParentBackdropImageTags         []string
	LocalTrailerCount               int //,optional
	UserData                        UserItemDataDto
	SeasonUserData                  UserItemDataDto
	RecursiveItemCount              int //,optional
	ChildCount                      int //,optional
	SeriesName                      string
	SeriesId                        string
	SeasonId                        string
	SpecialFeatureCount             int //,optional
	DisplayPreferencesId            string
	Status                          string //,optional) = ['Continuing' or 'Ended'],
	AirTime                         string
	AirDays                         []string
	IndexOptions                    []string
	Tags                            []string
	Keywords                        []string
	PrimaryImageAspectRatio         float64 //,optional
	OriginalPrimaryImageAspectRatio float64 //,optional
	Artists                         []string
	Album                           string
	CollectionType                  string
	DisplayOrder                    string
	AlbumId                         string
	AlbumPrimaryImageTag            string
	SeriesPrimaryImageTag           string
	AlbumArtist                     string
	SeasonName                      string
	MediaStreams                    []MediaStream
	VideoType                       string //,optional) = ['VideoFile' or 'Iso' or 'Dvd' or 'BluRay' or 'HdDvd'],
	DisplayMediaType                string
	PartCount                       int //,optional
	MediaSourceCount                int //,optional
	SupportsPlaylists               bool
	BackdropImageTags               []string
	ScreenshotImageTags             []string
	ParentLogoImageTag              string
	ParentArtItemId                 string
	ParentArtImageTag               string
	SeriesThumbImageTag             string
	SeriesStudio                    string
	ParentThumbItemId               string
	ParentThumbImageTag             string
	ParentPrimaryImageItemId        string
	ParentPrimaryImageTag           string
	Chapters                        []ChapterInfoDto
	LocationType                    string //) = ['FileSystem' or 'Remote' or 'Virtual' or 'Offline'],
	IsoType                         string //,optional) = ['Dvd' or 'BluRay'],
	MediaType                       string
	EndDate                         time.Time //,optional
	HomePageUrl                     string
	ProductionLocations             []string
	Budget                          float64 //,optional
	Revenue                         float64 //,optional
	MovieCount                      int     //,optional
	SeriesCount                     int     //,optional
	EpisodeCount                    int     //,optional
	GameCount                       int     //,optional
	SongCount                       int     //,optional
	AlbumCount                      int     //,optional
	MusicVideoCount                 int     //,optional
	LockData                        bool    //,optional
	Width                           int     //,optional
	Height                          int     //,optional
	CameraMake                      string
	CameraModel                     string
	Software                        string
	ExposureTime                    float64 //,optional
	FocalLength                     float64 //,optional
	ImageOrientation                string  //,optional) = ['TopLeft' or 'TopRight' or 'BottomRight' or 'BottomLeft' or 'LeftTop' or 'RightTop' or 'RightBottom' or 'LeftBottom'],
	Aperture                        float64 //,optional
	ShutterSpeed                    float64 //,optional
	Latitude                        float64 //,optional
	Longitude                       float64 //,optional
	Altitude                        float64 //,optional
	IsoSpeedRating                  int     //,optional
	CanResume                       bool
	ResumePositionTicks             int64
	BackdropCount                   int
	ScreenshotCount                 int
	HasBanner                       bool
	HasArtImage                     bool
	HasLogo                         bool
	HasThumb                        bool
	HasPrimaryImage                 bool
	HasDiscImage                    bool
	HasBoxImage                     bool
	HasBoxRearImage                 bool
	HasMenuImage                    bool
	IsVideo                         bool
	IsAudio                         bool
	IsGame                          bool
	IsPerson                        bool
	IsRoot                          bool
	IsMusicGenre                    bool
	IsGameGenre                     bool
	IsGenre                         bool
	IsArtist                        bool
	IsAlbum                         bool
	IsStudio                        bool
}
type ExternalUrl struct {
	Name string
	Url  string
}
type MediaSourceInfo struct {
	Protocol                   string //) = ['File' or 'Http' or 'Rtmp' or 'Rtsp'],
	Id                         string
	Path                       string
	Type                       string //) = ['Default' or 'Grouping' or 'Cache'],
	Container                  string
	Size                       int64 //,optional
	Name                       string
	RunTimeTicks               int64 //,optional
	ReadAtNativeFramerate      bool
	VideoType                  string        //,optional) = ['VideoFile' or 'Iso' or 'Dvd' or 'BluRay' or 'HdDvd'],
	IsoType                    string        //,optional) = ['Dvd' or 'BluRay'],
	Video3DFormat              string        //,optional) = ['HalfSideBySide' or 'FullSideBySide' or 'FullTopAndBottom' or 'HalfTopAndBottom'],
	MediaStreams               []MediaStream //Items/{ItemId}]
	PlayableStreamFileNames    []string
	Formats                    []string
	Bitrate                    int               //,optional
	Timestamp                  string            //,optional) = ['None' or 'Zero' or 'Valid'],
	RequiredHttpHeaders        map[string]string //`2[[String`/Items/{ItemId}
	DefaultAudioStreamIndex    int               //,optional
	DefaultSubtitleStreamIndex int               //,optional
	DefaultAudioStream         MediaStream       //Items/{ItemId}
	VideoStream                MediaStream       //Items/{ItemId})
}

type MediaStream struct {
	Codec                string
	Language             string
	IsInterlaced         bool
	ChannelLayout        string
	BitRate              int //,optional
	BitDepth             int //,optional
	RefFrames            int //,optional
	PacketLength         int //,optional
	Channels             int //,optional
	SampleRate           int //,optional
	IsDefault            bool
	IsForced             bool
	Height               int     //,optional
	Width                int     //,optional
	AverageFrameRate     float32 //,optional
	RealFrameRate        float32 //,optional
	Profile              string
	Type                 string //) = ['Audio' or 'Video' or 'Subtitle' or 'EmbeddedImage'],
	AspectRatio          string
	Index                int
	IsExternal           bool
	IsTextSubtitleStream bool
	Path                 string
	PixelFormat          string
	Level                float64 //,optional
	IsAnamorphic         bool    //,optional
	IsCabac              bool    //,optional)
}

type MediaUrl struct {
	Url       string
	Name      string
	VideoSize string //,optional) = ['StandardDefinition' or 'HighDefinition']
}

type BaseItemPerson struct {
	Name            string
	Id              string
	Role            string
	Type            string
	PrimaryImageTag string
	HasPrimaryImage bool
}
type StudioDto struct {
	Name            string
	Id              string
	PrimaryImageTag string
	HasPrimaryImage bool
}
type UserItemDataDto struct {
	Rating                float64 //,optional
	PlayedPercentage      float64 //,optional
	UnplayedItemCount     int     //,optional
	PlaybackPositionTicks int64
	PlayCount             int
	IsFavorite            bool
	Likes                 bool      //,optional
	LastPlayedDate        time.Time //,optional
	Played                bool
	Key                   string
}

type ChapterInfoDto struct {
	StartPositionTicks int64
	Name               string
	ImageTag           string
	HasImage           bool
}

////////////////
//User Resonse//
////////////////

type UserDto struct {
	Name                            string
	ServerId                        string
	ConnectUserName                 string
	ConnectUserId                   string
	ConnectLinkType                 string //, optional) = ['LinkedUser' or 'Guest'],
	Id                              string
	PrimaryImageTag                 string
	HasPassword                     bool
	HasConfiguredPassword           bool
	LastLoginDate                   time.Time //, optional
	LastActivityDate                time.Time //, optional
	Configuration                   UserConfiguration
	Policy                          UserPolicy
	PrimaryImageAspectRatio         float64 //, optional
	OriginalPrimaryImageAspectRatio float64 //, optional
	HasPrimaryImage                 bool
}
type UserConfiguration struct {
	MaxParentalRating               int //, optional
	IsAdministrator                 bool
	AudioLanguagePreference         string
	PlayDefaultAudioTrack           bool
	SubtitleLanguagePreference      string
	IsHidden                        bool
	IsDisabled                      bool
	DisplayMissingEpisodes          bool
	DisplayUnairedEpisodes          bool
	EnableRemoteControlOfOtherUsers bool
	EnableSharedDeviceControl       bool
	EnableLiveTvManagement          bool
	EnableLiveTvAccess              bool
	EnableMediaPlayback             bool
	EnableContentDeletion           bool
	GroupMoviesIntoBoxSets          bool
	BlockedMediaFolders             []string
	BlockedChannels                 []string
	DisplayChannelsWithinViews      []string
	ExcludeFoldersFromGrouping      []string
	SubtitleMode                    string //) = ['Default' or 'Always' or 'OnlyForced' or 'None'],
	DisplayCollectionsView          bool   //
	DisplayFoldersView              bool
	EnableLocalPassword             bool
	OrderedViews                    []string
	IncludeTrailersInSuggestions    bool
	EnableCinemaMode                bool
	AccessSchedules                 []AccessSchedule ///Users/{Id}
	EnableUserPreferenceAccess      bool
	LatestItemsExcludes             []string
	BlockedTags                     []string
	HasMigratedToPolicy             bool
}
type AccessSchedule struct { ///Users/{Id} {
	DayOfWeek string //) = ['Sunday' or 'Monday' or 'Tuesday' or 'Wednesday' or 'Thursday' or 'Friday' or 'Saturday' or 'Everyday' or 'Weekday' or 'Weekend'],
	StartHour float64
	EndHour   float64
}
type UserPolicy struct {
	IsAdministrator                 bool
	IsHidden                        bool
	IsDisabled                      bool
	MaxParentalRating               int //, optional
	BlockedTags                     []string
	EnableUserPreferenceAccess      bool
	AccessSchedules                 []AccessSchedule ///Users/{Id}
	BlockedMediaFolders             []string
	BlockedChannels                 []string
	EnableRemoteControlOfOtherUsers bool
	EnableSharedDeviceControl       bool
	EnableLiveTvManagement          bool
	EnableLiveTvAccess              bool
	EnableMediaPlayback             bool
	EnableContentDeletion           bool
	EnableSync                      bool
	EnabledDevices                  []string
	EnableAllDevices                bool
}

//////////////////
//Login Response//
//////////////////

type AuthenticationResult struct {
	User        UserDto
	SessionInfo SessionInfoDto
	AccessToken string
	ServerId    string
}

type SessionInfoDto struct {
	SupportedCommands     []string
	QueueableMediaTypes   []string
	PlayableMediaTypes    []string
	Id                    string
	UserId                string
	UserPrimaryImageTag   string
	UserName              string
	AdditionalUsers       []SessionUserInfo
	ApplicationVersion    string
	Client                string
	LastActivityDate      time.Time
	NowViewingItem        BaseItemInfo
	DeviceName            string
	NowPlayingItem        BaseItemInfo
	DeviceId              string
	SupportsRemoteControl bool
	PlayState             PlayerStateInfo
	TranscodingInfo       TranscodingInfo
}

type SessionUserInfo struct {
	UserId   string
	UserName string
}

type BaseItemInfo struct {
	Name                string
	Id                  string
	Type                string
	MediaType           string
	RunTimeTicks        int64 //, optional
	PrimaryImageTag     string
	PrimaryImageItemId  string
	LogoImageTag        string
	LogoItemId          string
	ThumbImageTag       string
	ThumbItemId         string
	BackdropImageTag    string
	BackdropItemId      string
	PremiereDate        time.Time //, optional
	ProductionYear      int       //, optional
	IndexNumber         int       //, optional
	IndexNumberEnd      int       //, optional
	ParentIndexNumber   int       //, optional
	SeriesName          string
	Album               string
	Artists             []string
	MediaStreams        []MediaStream
	ChapterImagesItemId string
	Chapters            []ChapterInfoDto
	HasPrimaryImage     bool
}

type PlayerStateInfo struct {
	PositionTicks       int64 //, optional),
	CanSeek             bool
	IsPaused            bool
	IsMuted             bool
	VolumeLevel         int //, optional),
	AudioStreamIndex    int //, optional),
	SubtitleStreamIndex int //, optional),
	MediaSourceId       string
	PlayMethod          string //, optional) = ['Transcode' or 'DirectStream' or 'DirectPlay']
}

type TranscodingInfo struct {
	AudioCodec           string
	VideoCodec           string
	Container            string
	IsVideoDirect        bool
	IsAudioDirect        bool
	Bitrate              int     //, optional),
	Framerate            float32 //, optional),
	CompletionPercentage float64 //, optional),
	Width                int     //, optional),
	Height               int     //, optional),
	AudioChannels        int     //, optional)
}
