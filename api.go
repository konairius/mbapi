package mbapi

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/youtube/vitess/go/event"
)

var (
	host     = flag.String("mb.host", "localhost", "The Mediabroswer you want to use as backend")
	user     = flag.String("mb.user", "", "The user to login to Mediabrowser")
	password = flag.String("mb.password", "", "The users Password")
)

//Backend is the main Construct that hold information about a Mediabrowser Server.
//It manages Authentication and actually executing the Query
type Backend struct {
	host   string
	client *http.Client
	auth   *AuthenticationResult
	UserId string
}

func NewBackend(host string) *Backend {
	b := &Backend{
		host:   host,
		client: &http.Client{},
	}
	return b
}

//NewRequest is the main Method of this Package.
//You get back the Complete http.Response because you to allow better control over what happens,
//Warps around "net/http" and manages the calls to the API
func (b *Backend) NewRequest(query Query) (*http.Response, error) {
	if b.host == "" {
		return nil, fmt.Errorf("Request Failed: Host not set")
	}
	q, err := RenderQuery(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to render Query: %v\n", err)
	}

	url := fmt.Sprintf("http://%s%s", b.host, q)
	// log.Printf("QueryUrl: %v\n", url)
	req, err := http.NewRequest(query.Method(), url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Request(%v): %v", url, err)
	}
	b.setHeader(req)
	//and execute
	resp, err := b.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to get: %v", err)
	}
	return resp, nil
}

//setHeader sets the Default header we use for requests
func (b *Backend) setHeader(req *http.Request) {
	hostname, _ := os.Hostname()
	//Set response type to something useful
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "MediaBrowser")
	req.Header.Set("Client", "Mediabrowser GO API")
	req.Header.Set("Device", hostname)
	req.Header.Set("DeviceId", hostname)
	req.Header.Set("Version", Version)
	//set Authentication
	if b.auth != nil {
		req.Header.Set("X-MediaBrowser-Token", b.auth.AccessToken)
	}

}

//Login Creates a LoginQuery and executes it the resulting AuthenticationResult is used to authenticate future Requests,
//this is probably the first Method you Call on a new Backend object.
func (b *Backend) Login(username, password string) error {

	password = fmt.Sprintf("%x", sha1.Sum([]byte(password)))

	query := &LoginQuery{
		Username: username,
		Password: password,
	}

	resp, err := b.NewRequest(query)
	if err != nil {
		event.Dispatch(err)
		return fmt.Errorf("Failed to Login: %v", err)
	}

	var auth AuthenticationResult
	err = RenderResponse(resp, &auth)
	if err != nil {
		event.Dispatch(err)
	}

	b.auth = &auth
	b.UserId = auth.User.Id
	return err
}

//StreamReader lets read from a "BaseItemDto" via the ReaderAt Interface.
//WARING: http Requests are SLOW, so if you may think about using a bigger buffer than you normally would.
type StreamReader struct {
	b    *Backend
	item *BaseItemDto
}

func (b *Backend) GetStreamReader(item *BaseItemDto) (*StreamReader, error) {
	reader := &StreamReader{
		b:    b,
		item: item,
	}

	if item.MediaSources[0].Size == 0 {
		return nil, fmt.Errorf("Invalid media Size: %s - %s", item.SeasonName, item.Name)
	}

	return reader, nil
}

func (s *StreamReader) ReadAt(p []byte, offset int64) (int, error) {

	resp, err := s.getResponse(s.getRequestSize(offset, int64(len(p))))
	if err != nil {
		return 0, fmt.Errorf("Failed to Read File: %v", err)
	}
	defer resp.Body.Close()
	//Find out how mutch we actualy got (Content-Length is bullcrap)
	// log.Printf("Content-Range:%s\n", resp.Header.Get("Content-Range"))
	// size, err := strconv.Atoi(strings.Split(resp.Header.Get("Content-Range"), "/")[1])
	// if err != nil {
	// 	return 0, fmt.Errorf("Failed to get size: %v", err)
	// }

	tn, n := 0, 0
	for tn < len(p) {
		// fmt.Printf("Read(tn=%d/%d,len(p)=%d)\n", tn, end-offset, len(p))
		n, err = resp.Body.Read(p[tn:])
		tn += n
		if err != nil && err != io.EOF {
			log.Printf("Unexpected Error: %v\n", err)
			break
		}
		if int64(tn)+offset >= s.item.MediaSources[0].Size {
			err = io.EOF
			break
		}
	}
	// if tn != len(p) {
	// 	log.Printf("Reader ended: %d/%d", tn, len(p))
	// }
	// log.Printf("Read(%#v)\n", p[:8])
	return tn, err
}

func (s *StreamReader) getResponse(start, end int64) (*http.Response, error) {
	query := &VideoStreamQuery{
		Id:            s.item.Id,
		MediaSourceId: s.item.MediaSources[0].Id, //We currently only handle single source media
		Static:        true,
	}
	url, err := RenderQuery(query)
	url = fmt.Sprintf("http://%s%s", s.b.host, url)
	// log.Printf("File: %s: %s\n", s.item.Name, url)
	if err != nil {
		return nil, fmt.Errorf("Failed to render Query: %v", err)
	}
	req, err := http.NewRequest(query.Method(), url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Request: %v", err)
	}
	s.b.setHeader(req)
	req.Header.Set("range", fmt.Sprintf("bytes=%d-%d", start, end))

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Http Requset Failed: %v", err)
	}
	return resp, nil
}

func (s *StreamReader) getRequestSize(off, req int64) (start, end int64) {
	end = off + req
	if end > int64(s.item.MediaSources[0].Size) {
		end = int64(s.item.MediaSources[0].Size)
	}
	return off, end
}
