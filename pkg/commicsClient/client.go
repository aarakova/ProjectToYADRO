package commicsclient

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	fast "github.com/valyala/fastjson"
)

const MaxOffsetDone = "max offset is done"

func NewClient(url string, maxOffset int, formatter Formatter) *Clint {
	return &Clint{maxOffset: maxOffset, formatter: formatter, pullUrl: url, offset: 1, jsonArena: &fast.Arena{}}
}

type Formatter interface {
	NormalizeText(string) []string
}

type Clint struct {
	formatter Formatter
	jsonArena *fast.Arena
	pullUrl   string
	offset    int
	maxOffset int
}

func (c *Clint) next() (string, *fast.Value, error) {
	response, err := http.Get(fmt.Sprintf("%s/%d/info.0.json", c.pullUrl, c.offset))
	if err != nil {
		return "", nil, err
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Errorf("cannot close response body: %v", err.Error())
		}
	}()

	var body []byte
	body, err = io.ReadAll(response.Body)
	if err != nil {
		return "", nil, err
	}

	jsonValue, err := fast.Parse(string(body))
	if err != nil {
		return "", nil, err
	}

	field := c.jsonArena.NewObject()
	field.Set("url", jsonValue.Get("img"))
	arr := c.jsonArena.NewArray()
	for i, v := range c.formatter.NormalizeText(jsonValue.Get("transcript").String()) {
		arr.SetArrayItem(i, c.jsonArena.NewString(v))
	}
	field.Set("keywords", arr)

	c.jsonArena.Reset()

	c.offset++

	return jsonValue.Get("num").String(), field, nil
}

func (c *Clint) ConvertAll() string {
	var err error
	var field *fast.Value
	var id string
	result := c.jsonArena.NewObject()
	for {
		if c.offset >= c.maxOffset && c.maxOffset != -1 {
			return result.String()
		}
		if id, field, err = c.next(); err == nil {
			result.Set(id, field)
			continue
		}
		log.Error(err)
		break
	}

	return result.String()
}
