package constant
import(
	"log"
)
const (
	OK = 0
	PARAMS_ERROR = 10000
	NOTEXIST_ERROR = 10001
	AUTH_ERROR = 20000
	REPASS_ERROR = 20001
	NETWORK_ERROR = 30000
	DB_ERROR = 40000
)
/**
	GLOBLE RESULT CODE MAPPING
**/
var CODE_MAPPING map[int]string

func init(){
	log.Println("Constant init()")
	CODE_MAPPING = map[int]string{
		0:"成功",
		10000:"参数错误",
		10001:"数据不存在",
		20000:"访问受限",
		20001:"重复密码错误",
		30000:"网络错误",
		40000:"数据库错误",
	}
}