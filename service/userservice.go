package service

/*-----------------单机不加缓存并发在10w，加入缓存可以达到100w------------------*/
import (
	"IMsystem/models"
	"IMsystem/utils"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// LogInUser
// @summary 用户登录
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/loginUser [post]
func LogInUser(c *gin.Context) {
	//使用gin将信息拿到页面，注意注释必须紧贴函数
	data := models.UserBasic{}
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1, //0成功 -1失败
			"message": "用户不存在",
		})
		return
	}
	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1, //0成功 -1失败
			"message": "密码不正确",
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	c.JSON(200, gin.H{
		"code":    0, //0成功 -1失败
		"message": "登录成功",
		"data":    data,
	})
}

// CreateUser
// @summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "两次密码不一样!",
		})
		return
	}
	//增加对新增用户的校验
	validate1 := models.FindUserByName(user.Name)
	//后面两字段不生效，不在注册页面中显示
	validate2 := models.FindUserByEmail(user.Email)
	validate3 := models.FindUserByPhone(user.Phone)
	if validate1.Name != "" || validate2.Email != "" || validate3.Phone != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名已存在",
		})
		return
	}
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	user.LoginTime = time.Now()
	user.HeartBeatTime = time.Now()
	user.LoginOutTime = time.Now()
	err := models.CreateUser(user)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "新增用户成功",
			"data":    user,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "新增用户失败",
			"data":    user,
		})
		return
	}
}

// DeleteUser
// @summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"code":    -1,
		"message": "删除用户成功",
		"data":    user,
	})
}

// UpdateUser
// @summary 修改用户信息
// @Tags 用户模块
// @param id query string false "id"
// @param name query string false "用户名"
// @param password query string false "密码"
// @param email query string false "邮箱"
// @param phone query string false "电话"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	user.Name = c.Query("name")
	user.PassWord = c.Query("password")
	user.Email = c.Query("email")
	user.Phone = c.Query("phone")
	user.LoginTime = time.Now()
	user.HeartBeatTime = time.Now()
	user.LoginOutTime = time.Now()
	_, err := govalidator.ValidateStruct(user) //增加校验规则
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "更新用户失败",
			"data":    user,
		})
		return
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"code":    0,
			"message": "更新用户成功",
			"data":    user,
		})
	}
}

// GetUser
// @summary 查找用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/getUser [post]
func GetUser(c *gin.Context) {
	//pwd查询需要解密,分别判断name和pwd存在
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "不存在该用户",
			"data":    user,
		})
		return
	}
	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "密码不正确",
			"data":    user,
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data := models.GetUser(name, pwd)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "查找成功",
		"data":    data,
	})
}

var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(ws, c)

}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if err != nil {
		fmt.Println(err)
	}
	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println(err)
	}
}
