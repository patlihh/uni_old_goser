// Usys_language
package usql

//语言表
type SysLanguage struct {
	code_id      [20]byte  //  NOT NULL COMMENT '语言编码',
	code_name    [200]byte //   DEFAULT NULL COMMENT '语言名称',
	remark       [500]byte //  DEFAULT NULL COMMENT '备注',
	lang_img     [500]byte //   DEFAULT NULL COMMENT '语言图标',
	login_img    [500]byte //   DEFAULT NULL COMMENT '登录按钮图片',
	username_img [500]byte //   DEFAULT NULL COMMENT '用户名图片',
	password_img [500]byte //  DEFAULT NULL COMMENT '密码图片',
	logout_img   [500]byte //   DEFAULT NULL COMMENT '注销图片',
}
