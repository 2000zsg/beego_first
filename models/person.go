package models

/*/
人结构体定义 包含三个字段，与前端内容一致
 */
type Person struct {
	Name string
	Age int
	Sex string
}
type Air struct {
Name string
Birthday float64
Address string
Nick string
}
type User struct {
	Name string `form:"user_name"`
	Birthday float64 `form:"birthday"`
	Address string `form:"address"`
	Nick string `form:"nick"`
	Password string `form:"password"`
}
type ResponseResult struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
