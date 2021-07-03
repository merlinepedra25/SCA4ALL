package tomcat_brute

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type userpass struct {
	username string
	password string
}

var (
	up        = []userpass{{"admin", ""}, {"Tomcat-manager", "manager"}, {"root", ""}, {"test", "test"}, {"root", "admin"}, {"admin", "admin"}, {"root", "123456"}, {"admin", "123456"}, {"admin", "admanager"}, {"admin", "admin"}, {"ADMIN", "ADMIN"}, {"admin", "adrole1"}, {"admin", "adroot"}, {"admin", "ads3cret"}, {"admin", "adtomcat"}, {"admin", "advagrant"}, {"admin", "password"}, {"admin", "password1"}, {"admin", "Password1"}, {"admin", "tomcat"}, {"admin", "vagrant"}, {"both", "admanager"}, {"both", "admin"}, {"both", "adrole1"}, {"both", "adroot"}, {"both", "ads3cret"}, {"both", "adtomcat"}, {"both", "advagrant"}, {"both", "tomcat"}, {"cxsdk", "kdsxc"}, {"j2deployer", "j2deployer"}, {"manager", "admanager"}, {"manager", "admin"}, {"manager", "adrole1"}, {"manager", "adroot"}, {"manager", "ads3cret"}, {"manager", "adtomcat"}, {"manager", "advagrant"}, {"manager", "manager"}, {"ovwebusr", "OvW*busr1"}, {"QCC", "QLogic66"}, {"role1", "admanager"}, {"role1", "admin"}, {"role1", "adrole1"}, {"role1", "adroot"}, {"role1", "ads3cret"}, {"role1", "adtomcat"}, {"role1", "advagrant"}, {"role1", "role1"}, {"role1", "tomcat"}, {"role", "changethis"}, {"root", "admanager"}, {"root", "admin"}, {"root", "adrole1"}, {"root", "adroot"}, {"root", "ads3cret"}, {"root", "adtomcat"}, {"root", "advagrant"}, {"root", "changethis"}, {"root", "owaspbwa"}, {"root", "password"}, {"root", "password1"}, {"root", "Password1"}, {"root", "r00t"}, {"root", "root"}, {"root", "toor"}, {"tomcat", ""}, {"tomcat", "admanager"}, {"tomcat", "admin"}, {"tomcat", "adrole1"}, {"tomcat", "adroot"}, {"tomcat", "ads3cret"}, {"tomcat", "adtomcat"}, {"tomcat", "advagrant"}, {"tomcat", "changethis"}, {"tomcat", "password"}, {"tomcat", "password1"}, {"tomcat", "s3cret"}, {"tomcat", "tomcat"}, {"xampp", "xampp"}, {"server_admin", "owaspbwa"}, {"admin", "owaspbwa"}, {"demo", "demo"}, {"root", "root123"}, {"root", "password"}, {"root", "root@123"}, {"root", "root888"}, {"root", "root"}, {"root", "a123456"}, {"root", "123456a"}, {"root", "5201314"}, {"root", "111111"}, {"root", "woaini1314"}, {"root", "qq123456"}, {"root", "123123"}, {"root", "000000"}, {"root", "1qaz2wsx"}, {"root", "1q2w3e4r"}, {"root", "qwe123"}, {"root", "7758521"}, {"root", "123qwe"}, {"root", "a123123"}, {"root", "123456aa"}, {"root", "woaini520"}, {"root", "woaini"}, {"root", "100200"}, {"root", "1314520"}, {"root", "woaini123"}, {"root", "123321"}, {"root", "q123456"}, {"root", "123456789"}, {"root", "123456789a"}, {"root", "5211314"}, {"root", "asd123"}, {"root", "a123456789"}, {"root", "z123456"}, {"root", "asd123456"}, {"root", "a5201314"}, {"root", "aa123456"}, {"root", "zhang123"}, {"root", "aptx4869"}, {"root", "123123a"}, {"root", "1q2w3e4r5t"}, {"root", "1qazxsw2"}, {"root", "5201314a"}, {"root", "1q2w3e"}, {"root", "aini1314"}, {"root", "31415926"}, {"root", "q1w2e3r4"}, {"root", "123456qq"}, {"root", "woaini521"}, {"root", "1234qwer"}, {"root", "a111111"}, {"root", "520520"}, {"root", "iloveyou"}, {"root", "abc123"}, {"root", "110110"}, {"root", "111111a"}, {"root", "123456abc"}, {"root", "w123456"}, {"root", "7758258"}, {"root", "123qweasd"}, {"root", "159753"}, {"root", "qwer1234"}, {"root", "a000000"}, {"root", "qq123123"}, {"root", "zxc123"}, {"root", "123654"}, {"root", "abc123456"}, {"root", "123456q"}, {"root", "qq5201314"}, {"root", "12345678"}, {"root", "000000a"}, {"root", "456852"}, {"root", "as123456"}, {"root", "1314521"}, {"root", "112233"}, {"root", "521521"}, {"root", "qazwsx123"}, {"root", "zxc123456"}, {"root", "abcd1234"}, {"root", "asdasd"}, {"root", "666666"}, {"root", "love1314"}, {"root", "QAZ123"}, {"root", "aaa123"}, {"root", "q1w2e3"}, {"root", "aaaaaa"}, {"root", "a123321"}, {"root", "123000"}, {"root", "11111111"}, {"root", "12qwaszx"}, {"root", "5845201314"}, {"root", "s123456"}, {"root", "nihao123"}, {"root", "caonima123"}, {"root", "zxcvbnm123"}, {"root", "wang123"}, {"root", "159357"}, {"root", "1A2B3C4D"}, {"root", "asdasd123"}, {"root", "584520"}, {"root", "753951"}, {"root", "147258"}, {"root", "1123581321"}, {"root", "110120"}, {"root", "qq1314520"}, {"admin", "admin123"}, {"admin", "password"}, {"admin", "admin@123"}, {"admin", "admin888"}, {"admin", "root"}, {"admin", "a123456"}, {"admin", "123456a"}, {"admin", "5201314"}, {"admin", "111111"}, {"admin", "woaini1314"}, {"admin", "qq123456"}, {"admin", "123123"}, {"admin", "000000"}, {"admin", "1qaz2wsx"}, {"admin", "1q2w3e4r"}, {"admin", "qwe123"}, {"admin", "7758521"}, {"admin", "123qwe"}, {"admin", "a123123"}, {"admin", "123456aa"}, {"admin", "woaini520"}, {"admin", "woaini"}, {"admin", "100200"}, {"admin", "1314520"}, {"admin", "woaini123"}, {"admin", "123321"}, {"admin", "q123456"}, {"admin", "123456789"}, {"admin", "123456789a"}, {"admin", "5211314"}, {"admin", "asd123"}, {"admin", "a123456789"}, {"admin", "z123456"}, {"admin", "asd123456"}, {"admin", "a5201314"}, {"admin", "aa123456"}, {"admin", "zhang123"}, {"admin", "aptx4869"}, {"admin", "123123a"}, {"admin", "1q2w3e4r5t"}, {"admin", "1qazxsw2"}, {"admin", "5201314a"}, {"admin", "1q2w3e"}, {"admin", "aini1314"}, {"admin", "31415926"}, {"admin", "q1w2e3r4"}, {"admin", "123456qq"}, {"admin", "woaini521"}, {"admin", "1234qwer"}, {"admin", "a111111"}, {"admin", "520520"}, {"admin", "iloveyou"}, {"admin", "abc123"}, {"admin", "110110"}, {"admin", "111111a"}, {"admin", "123456abc"}, {"admin", "w123456"}, {"admin", "7758258"}, {"admin", "123qweasd"}, {"admin", "159753"}, {"admin", "qwer1234"}, {"admin", "a000000"}, {"admin", "qq123123"}, {"admin", "zxc123"}, {"admin", "123654"}, {"admin", "abc123456"}, {"admin", "123456q"}, {"admin", "qq5201314"}, {"admin", "12345678"}, {"admin", "000000a"}, {"admin", "456852"}, {"admin", "as123456"}, {"admin", "1314521"}, {"admin", "112233"}, {"admin", "521521"}, {"admin", "qazwsx123"}, {"admin", "zxc123456"}, {"admin", "abcd1234"}, {"admin", "asdasd"}, {"admin", "666666"}, {"admin", "love1314"}, {"admin", "QAZ123"}, {"admin", "aaa123"}, {"admin", "q1w2e3"}, {"admin", "aaaaaa"}, {"admin", "a123321"}, {"admin", "123000"}, {"admin", "11111111"}, {"admin", "12qwaszx"}, {"admin", "5845201314"}, {"admin", "s123456"}, {"admin", "nihao123"}, {"admin", "caonima123"}, {"admin", "zxcvbnm123"}, {"admin", "wang123"}, {"admin", "159357"}, {"admin", "1A2B3C4D"}, {"admin", "asdasd123"}, {"admin", "584520"}, {"admin", "753951"}, {"admin", "147258"}, {"admin", "1123581321"}, {"admin", "110120"}, {"admin", "qq1314520"}}
	httpProxy string
)

func httpRequset(username string, password string, loginurl string) int {
	//设置跳过https证书验证，超时和代理
	var tr *http.Transport
	if httpProxy != "" {
		uri, _ := url.Parse(httpProxy)
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyURL(uri),
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{
		Timeout:   time.Duration(5) * time.Second,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse //不允许跳转
		}}
	req, err := http.NewRequest(strings.ToUpper("GET"), loginurl+"/manager/html", strings.NewReader(""))
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		return resp.StatusCode
	}
	return 999999
}

func makeGuesses(users, passwords []string) []userpass {
	var guesses = make([]userpass, 0)
	for _, username := range users {
		for _, password := range passwords {
			guesses = append(guesses, userpass{username, password})
		}
	}
	return guesses
}

func Check(url string) (username string, password string) {
	//httpProxy = "http://127.0.0.1:8080"
	if httpRequset("", "", url) == 401 {
		for uspa := range up {
			code := httpRequset(up[uspa].username, up[uspa].password, url)
			if code == 200 {
				fmt.Printf("tomcat-brute-sucess|%s:%s--%s", up[uspa].username, up[uspa].password, url)
				fmt.Println()
				return up[uspa].username, up[uspa].password
			}
		}
	}
	return "", ""
}
