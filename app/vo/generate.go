// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 代码生成器Vo
 * @author 半城风雨
 * @since 2021/11/15
 * @File : generate
 */
package vo

import "time"

// 数据库信息
type GenerateInfo struct {
	Name          string    `json:"name"`           // 表名
	Engine        string    `json:"engine"`         // 引擎
	Version       string    `json:"version"`        // 版本
	Collation     string    `json:"collation"`      // 编码
	Rows          int       `json:"rows"`           // 记录数
	DataLength    int       `json:"data_length"`    // 大小
	AutoIncrement int       `json:"auto_increment"` // 自增索引
	Comment       string    `json:"comment"`        // 表备注
	CreateTime    time.Time `json:"createTime"`     // 添加时间
	UpdateTime    time.Time `json:"updateTime"`     // 更新时间
}
