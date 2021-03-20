package Help

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetEditUrls() []string {
	client := &http.Client{Timeout: 5 * time.Second}

	var request *http.Request

	page := PAGE
	var articleIds []string
	for {
		url := getUrl(page, "false")
		request, _ = http.NewRequest("GET", url, strings.NewReader(""))

		response, err := client.Do(request)

		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		responseData, _ := ioutil.ReadAll(response.Body)

		articleIdsTmp := gjson.Get(string(responseData), "data.list.#.articleId")
		fmt.Println(articleIdsTmp)
		articleIdsTmp.ForEach(func(_, value gjson.Result) bool {
			articleIds = append(articleIds, value.String())
			return true
		})

		fmt.Println(len(articleIds))
		total := int(gjson.Get(string(responseData), "data.total").Int())
		if total <= len(articleIds) {
			break
		}
		page++

		fmt.Print(articleIds)

		time.Sleep(2 * time.Second)
	}

	return articleIds
}

func getUrl(page int, noMore string) string {
	var url bytes.Buffer

	url.WriteString(BlogListUrl)
	url.WriteString("page=" + strconv.Itoa(page))
	url.WriteString("&size=" + strconv.Itoa(PageSize))
	url.WriteString("&businessType=" + BusinessType)
	url.WriteString("&noMore=" + noMore)
	url.WriteString("&username=" + UserName)

	fmt.Println(url.String())
	return url.String()
}

func GetBlogContent(articleIds []string) {
	// 登录状态获取文章内容

}
