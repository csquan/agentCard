DROP TABLE IF EXISTS `t_agentPersonalInfo`;
CREATE TABLE `t_agentPersonalInfo`
(
    `f_id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `f_agentName`         text COMMENT '代理商名称',
    `f_mobileNo`          text COMMENT '手机号',
    `f_certType`          text COMMENT '证件类型',
    `f_certNo`            text COMMENT '证件号',
    `f_name`              text COMMENT '姓名',
    `f_mailbox`           text COMMENT '邮箱',
    `f_addr`              text COMMENT '地址',
    `f_certPhoto`         text COMMENT '证件图片',
    `f_certOther`         text COMMENT '其它材料',
    `f_password`          text COMMENT '密码',
    `f_created_at`        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'time',
    `f_updated_at`        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON update CURRENT_TIMESTAMP COMMENT 'time',
    PRIMARY KEY (`f_id`) /*T![clustered_index] CLUSTERED */
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='代理商个人信息表';

DROP TABLE IF EXISTS `t_agentProfitInfo`;
CREATE TABLE `t_agentProfitInfo`
(
    `f_id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `f_agentName`         text COMMENT '代理商名称',
    `f_settlementRule`    text COMMENT '结算规则(周结/月结)',
    `f_activeRebate`      text COMMENT '激活返佣',
    `f_withdrawThreshold` text COMMENT '提现门槛',
    `f_grade1proportion`  text COMMENT '充值档1的返还比例(0-50万usdt)',
    `f_grade2proportion`  text COMMENT '充值档2的返还比例(50-100万usdt-不包含50)',
    `f_grade3proportion`  text COMMENT '充值档3的返还比例(100万usdt-不包含100)',
    `f_created_at`        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'time',
    `f_updated_at`        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON update CURRENT_TIMESTAMP COMMENT 'time',
    PRIMARY KEY (`f_id`) /*T![clustered_index] CLUSTERED */
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='代理商收益规则表';


DROP TABLE IF EXISTS `t_profitHistory`;
CREATE TABLE `t_profitHistory`
(
    `f_id`               bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `f_agentName`        text COMMENT '代理商名称',
    `f_grant_at`         text COMMENT '收益发放时间',
    `f_settlementMoney`  text COMMENT '收益结算金额',
    `f_settlementProfit` text COMMENT '收益结算利润',
    `f_grant_at`         text COMMENT '收益发放时间',
    `f_settlementCycle`  text COMMENT '收益结算周期',
    `f_created_at`       timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'time',
    `f_updated_at`       timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON update CURRENT_TIMESTAMP COMMENT 'time',
    PRIMARY KEY (`f_id`) /*T![clustered_index] CLUSTERED */
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='代理商收益发放表';


DROP TABLE IF EXISTS `t_agentCardInfo`;
CREATE TABLE `t_agentCardInfo`
(
    `f_id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `f_agentName`  text COMMENT '代理商名称',
    `f_cardNo`     text COMMENT '卡号',
    `f_cardRebate` text COMMENT '卡是否返佣',
    `f_address`    text COMMENT '卡对应的usdt-trc20地址',
    `f_created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'time',
    `f_updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON update CURRENT_TIMESTAMP COMMENT 'time',
    PRIMARY KEY (`f_id`) /*T![clustered_index] CLUSTERED */
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='转账表';



DROP TABLE IF EXISTS `t_agentWithdraw`;
CREATE TABLE `t_agentWithdraw`
(
    `f_id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `f_agentName`    text COMMENT '代理商名称',
    `f_withdrawAddr` text COMMENT '提现地址',
    `f_created_at`   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'time',
    `f_updated_at`   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON update CURRENT_TIMESTAMP COMMENT 'time',
    PRIMARY KEY (`f_id`) /*T![clustered_index] CLUSTERED */
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='转账表';


