package wget

import (
	"dev09/htmlparser"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
)

type Flags struct {
	R bool
}

type wget struct {
	dir               string
	assetsDir         string
	lastDownLoadedIdx int
	htmls             []htmlparser.HTMLPath
	hostregular       *regexp.Regexp
	visitedLinks      map[string]struct{}
	Flags
}

var (
	ErrInvalidUrl = errors.New("url должен быть полным")
	ErrLinkDomain = errors.New("ссылка не удволетворяет домену")
)

func NewWget(link string, flags Flags) (*wget, error) {
	url, err := url.Parse(link)
	if err != nil {
		fmt.Println(err)
	}
	hostName := url.Hostname()
	if hostName == "" {
		return nil, ErrInvalidUrl
	}
	re, err := regexp.Compile("https{0,1}://.{0,}" + hostName)
	if err != nil {
		return nil, err
	}
	w := wget{
		hostregular: re,
		Flags:       flags,
	}

	tempDir, err := ioutil.TempDir(".", "wgetout")
	if err != nil {
		defer os.Remove(tempDir)
		return nil, err
	}

	if err = os.Mkdir(path.Join(tempDir, "assets"), 0755); err != nil {
		return nil, err
	}

	w.dir = tempDir
	w.htmls = make([]htmlparser.HTMLPath, 0)
	w.visitedLinks = make(map[string]struct{})
	w.assetsDir = path.Join(tempDir, "assets")

	if err := w.download(link); err != nil {
		defer os.Remove(tempDir)
		return nil, err
	}

	if w.Flags.R {
		w.recursiveDownload()
	}
	w.cleanup()
	return &w, nil
}

func (w *wget) cleanup() {
	files, err := ioutil.ReadDir(w.assetsDir)
	if err != nil {
		return
	}

	if len(files) == 0 {
		os.Remove(w.assetsDir)
	}
}

func (w *wget) download(link string) error {
	resp, req, err := w.makeRequest(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	isHTML := resp.Header.Get("Content-type") == "text/html; charset=utf-8"
	dirpath := w.assetsDir
	if isHTML {
		dirpath = w.dir
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = w.makeFile(path.Join(dirpath, path.Base(req.URL.Path)), body)
	if err != nil {
		return err
	}

	clearLink := htmlparser.BuildClearLink(req.URL)
	if isHTML {
		if _, ok := w.visitedLinks[clearLink]; !ok {
			w.htmls = append(w.htmls, htmlparser.HTMLPath{
				ClearLink: clearLink,
				FilePath:  path.Join(dirpath, path.Base(req.URL.Path)),
			})
		}
	}
	w.visitedLinks[link] = struct{}{}
	w.visitedLinks[clearLink] = struct{}{}

	return nil
}

func (wget *wget) Dir() string {
	return wget.dir
}

func (wget *wget) DirAssets() string {
	return wget.assetsDir
}

// makeRequest знает как сделать запрос если ссылка удволетворяет домену
func (w *wget) makeRequest(link string) (*http.Response, *http.Request, error) {
	if _, ok := w.visitedLinks[link]; ok {
		return nil, nil, errors.New("ссылка уже посещенна")
	}

	if matched := w.hostregular.MatchString(link); !matched {
		fmt.Println("Невалидная ссылка:", link)
		return nil, nil, ErrLinkDomain
	}

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	return resp, req, nil
}

// makeFile знает как создать файл во временной директории
func (w *wget) makeFile(path string, body []byte) error {
	file, err := os.Create(path)
	if err != nil {
		file, err = ioutil.TempFile(path, "index")
		if err != nil {
			return err
		}
	}
	defer file.Close()
	if _, err = file.Write(body); err != nil {
		return err
	}

	return nil
}

func (w *wget) recursiveDownload() {
	currentLenght := len(w.htmls)

	for i := w.lastDownLoadedIdx; i < currentLenght; i++ {
		links, err := htmlparser.ParseHTML(w.htmls[i])
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, link := range links {
			if _, ok := w.visitedLinks[link]; ok {
				continue
			}

			w.download(link)
			w.visitedLinks[link] = struct{}{}
		}
	}

	w.lastDownLoadedIdx = len(w.htmls) - 1

	if len(w.htmls) != currentLenght {
		w.recursiveDownload()
	}
}
