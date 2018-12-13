// Usys_log
package usql

//用户日志表
type SysLog struct {
	id            [20]byte  //  NOT NULL COMMENT 'id',
	act_time      [20]byte  //  DEFAULT NULL COMMENT '时间',
	act_user_id   [20]byte  //  DEFAULT NULL COMMENT '操作用户',
	act_module_id [20]byte  //  DEFAULT NULL COMMENT '操作模块',
	act_name      [200]byte //  DEFAULT NULL COMMENT '操作名称',
	act_object    [20]byte  //  DEFAULT NULL COMMENT '操作对象',
	act_result    [20]byte  //  DEFAULT NULL COMMENT '操作结果',
}
