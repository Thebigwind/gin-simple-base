package rest

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"gin-simple-base/service"
	"github.com/gin-gonic/gin"
)

/*
首先code是前端获取传到后端的, 我们不用管, 只要在HTTP请求中拿到这个参数即可.
ok, 第一步的code已经拿到.

然后第二步, 利用code获取openID 和session_key, 这里我们看微信官方文档给的接口
*/

// /wechat/applet_login?code=xxx [get]  路由
// 微信小程序登录
func AppletWeChatLogin(c *gin.Context) {
	code := c.Query("code") //  获取code
	// 根据code获取 openID 和 session_key
	wxLoginResp, err := service.WXLogin(code)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	// 保存登录态
	//session := sessions.Default(c)
	//session.Set("openid", wxLoginResp.OpenId)
	//session.Set("sessionKey", wxLoginResp.SessionKey)

	// 这里用openid和sessionkey的串接 进行MD5之后作为该用户的自定义登录态
	mySession := GetMD5Encode(wxLoginResp.OpenId + wxLoginResp.SessionKey)
	// 接下来可以将openid 和 sessionkey, mySession 存储到数据库中,
	// 但这里要保证mySession 唯一, 以便于用mySession去索引openid 和sessionkey
	c.String(200, mySession)

}

// 将一个字符串进行MD5加密后返回加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// 校验微信返回的用户数据
func ValidateUserInfo(rawData, sessionKey, signature string) bool {
	signature2 := GetSha1(rawData + sessionKey)
	return signature == signature2
}

// SHA-1 加密
func GetSha1(str string) string {
	data := []byte(str)
	has := sha1.Sum(data)
	res := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return res
}
