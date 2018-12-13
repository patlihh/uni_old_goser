// Usys_operator_log
package usql

//系统操作日记表
type SysOperatorLog struct {
	ID            [20]byte  // not null comment 'ID',
	userID        [20]byte  //  comment '用户ID',
	operationType [50]byte  //  comment '操作类型',
	operationTime [30]byte  //  comment '操作时间',
	description   [50]byte  // comment '描述',
	result        [10]byte  //  comment '操作结果',
	exceptionInfo [200]byte //  comment '异常信息',
	menuID        [20]byte  //  comment '功能菜单ID',
	userIp        [20]byte  //  comment '用户IP(本机用localhost)',
	attribute2    [20]byte  //  comment '备用字段2',
	attribute3    [20]byte  //  comment '备用字段3',
}
