package mock

import (
	"app.nazul/models"
)

var BOOKS_DATA []models.Book

func init() {
	BOOKS_DATA = append(BOOKS_DATA, models.Book{
		Id:          "1",
		Isbn:        "9787520113113",
		Title:       "偏见",
		Author:      "十三邀",
		Image:       "https://img1.doubanio.com/view/subject/l/public/s29763168.jpg",
		Description: "谈话自有它的内在逻辑，它逼迫讲述者勾勒自己的轮廓、探视自己的内心。判断很可能片面、浅薄与武断，但背后，是我们对他人与时代真诚的理解欲望。",
	})
	BOOKS_DATA = append(BOOKS_DATA, models.Book{
		Id:          "2",
		Isbn:        "9787533952136",
		Title:       "庆熹纪事",
		Author:      "红猪侠",
		Image:       "https://img1.doubanio.com/view/subject/l/public/s29763169.jpg",
		Description: "风云开阖，忠贤灭门，他在阿鼻地狱中涅槃重生，却甘为深宫贱奴，为仇人之子驱使，只愿亲手撤藩地、平边患，一竟父志。",
	})
	BOOKS_DATA = append(BOOKS_DATA, models.Book{
		Id:          "3",
		Isbn:        "9787020135318",
		Title:       "故宫的古物之美",
		Author:      "祝勇",
		Image:       "https://img3.doubanio.com/view/subject/l/public/s29746360.jpg",
		Description: "国家宝藏的前世今生，故宫艺术的典藏读本\n了不起的中国古物，说不尽的华夏历史，再现一段文明的营造之美",
	})
	BOOKS_DATA = append(BOOKS_DATA, models.Book{
		Id:          "4",
		Isbn:        "9787208148109",
		Title:       "贸易打造的世界",
		Author:      "[美] 彭慕兰 / [美] 史蒂文·托皮克",
		Image:       "https://img3.doubanio.com/view/subject/l/public/s29645112.jpg",
		Description: "纵跨600年 | 83篇小史 | 30年研究精髓 |热销17年 | 3次再版更新\n《大分流》作者、美国历史学会会长彭慕兰\n拉美史、全球贸易史专家史蒂文•托皮克 共同书写\n全球化进程中的独立与依附、崇高与卑劣、失序与进步、不满与餍足",
	})
	BOOKS_DATA = append(BOOKS_DATA, models.Book{
		Id:          "5",
		Isbn:        "9787508682358",
		Title:       "失衡",
		Author:      "[美] 马胜学",
		Image:       "https://img3.doubanio.com/view/subject/l/public/s29752086.jpg",
		Description: "不久前，国际食物政策研究所一项面向全球的调查表明，全球有三分之一的人口处于营养失衡状态。所谓营养失衡，包括两方面：营养缺乏和营养过剩，前者一般出现在经济欠发达地区，而后者在发达国家更为常见。营养失衡正以各种形式在世界各地蔓延，贫困和富裕国家都受到了严重的影响。",
	})
	BOOKS_DATA = append(BOOKS_DATA, models.Book{
		Id:          "6",
		Isbn:        "9787570203345",
		Title:       "游荡者",
		Author:      "[德] 梁柯",
		Image:       "https://img3.doubanio.com/view/subject/l/public/s29739200.jpg",
		Description: "自小遭父母遗弃的徐猛被黑社会豢养，成为令人闻风丧胆的黑道煞星。一次行刺，他驾车逃离时不慎将路人撞成重伤，锒铛入狱。两年后，刑满释放的他发现自己的团伙已不知去向，江湖上所有的熟人都消失不见，他成了孤家寡人，被抛弃在这个城市。然而，噩梦还没有结束。一天清晨，他从梦中惊醒，发 现自己变成了一个陌生人！",
	})
}
