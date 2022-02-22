CREATE TABLE `individual_bankroll` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `day_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '几天时间',
  `individual_code` varchar(10) NOT NULL DEFAULT '' COMMENT '股票编号',
  `individual_name` varchar(20) NOT NULL DEFAULT '' COMMENT '股票名称',
  `end_price` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '最新价格',
  `rose_ratio` float(10,4) NOT NULL DEFAULT '0.0000' COMMENT '涨跌幅',
  `turnover_ratio` float(10,4) NOT NULL DEFAULT '0.0000' COMMENT '换手率',
  `fund_amount_in` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '流入资金',
  `fund_amount_out` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '流出资金',
  `fund_real_in` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '净额',
  `ob_price` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '成交额（元）',
  `c_date` date NOT NULL DEFAULT '1970-01-01',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`),
  UNIQUE KEY `icd_index` (`individual_code`,`c_date`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='个股资金流';

CREATE TABLE `individual_stock` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `industry_code` varchar(10) NOT NULL DEFAULT '',
  `individual_code` varchar(10) NOT NULL DEFAULT '' COMMENT '股票编号',
  `individual_name` varchar(20) NOT NULL DEFAULT '' COMMENT '股票名称',
  `now_price` float(10,2) NOT NULL DEFAULT '0.00' COMMENT '当前价',
  `rose_ratio` float(12,4) NOT NULL DEFAULT '0.0000' COMMENT '涨跌幅',
  `turnover_ratio` float(12,4) NOT NULL DEFAULT '0.0000' COMMENT '换手率',
  `relative` float(12,4) NOT NULL DEFAULT '0.0000' COMMENT '量比',
  `amplitude_ratio` float(12,4) NOT NULL DEFAULT '0.0000' COMMENT '振幅',
  `ob_price` double(14,2) NOT NULL DEFAULT '0.00' COMMENT '成交额（元）',
  `circulate_stock` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '流通股',
  `circulate_value` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '流通市值（元）',
  `pe` float(10,4) NOT NULL DEFAULT '0.0000' COMMENT '市盈率',
  `c_date` date NOT NULL DEFAULT '1970-01-01',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`),
  UNIQUE KEY `icd_index` (`individual_code`,`c_date`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='个股信息';

CREATE TABLE `industry_bankroll` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `fund_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型，1表示概念，2表示行业',
  `day_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '几天时间',
  `industry_code` varchar(10) NOT NULL DEFAULT '' COMMENT '行业编号',
  `industry_name` varchar(20) NOT NULL DEFAULT '' COMMENT '行业名称',
  `industry_index` float(10,2) NOT NULL DEFAULT '0.00' COMMENT '行业指数',
  `rose_ratio` float(12,4) NOT NULL DEFAULT '0.0000' COMMENT '涨跌幅',
  `fund_amount_out` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '流出资金',
  `fund_amount_in` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '流入资金',
  `fund_real_in` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '净额',
  `company_num` smallint(5) NOT NULL DEFAULT '0' COMMENT '公司家数',
  `c_date` date NOT NULL DEFAULT '1970-01-01',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`),
  UNIQUE KEY `icd_index` (`industry_code`,`c_date`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='行业资金';

-- 序号	板块	涨跌幅(%)	总成交量（万手）	总成交额（亿元）	净流入（亿元）	上涨家数	下跌家数	均价	领涨股	最新价	涨跌幅(%)
CREATE TABLE `plate_bankroll` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `plate_code` varchar(10) NOT NULL DEFAULT '' COMMENT '行业编号',
  `plate_name` varchar(20) NOT NULL DEFAULT '' COMMENT '行业名称',
  `rose_ratio` float(12,4) NOT NULL DEFAULT '0.0000' COMMENT '涨跌幅',
  `ob_price` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '总成交量（手）',
  `ob_price` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '成交额（元）',
  `fund_real_in` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '净流入额',
  `rise_company_num` smallint(5) NOT NULL DEFAULT '0' COMMENT '上涨家数',
  `drop_company_num` smallint(5) NOT NULL DEFAULT '0' COMMENT '下跌家数',
  `avg_price` double(20,2) NOT NULL DEFAULT '0.00' COMMENT '均价',
  `c_date` date NOT NULL DEFAULT '1970-01-01',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`),
  UNIQUE KEY `icd_index` (`industry_code`,`c_date`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='板块资金';


CREATE TABLE `relat_industry_individual` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `industry_code` varchar(10) NOT NULL DEFAULT '' COMMENT '行业编号',
  `individual_code` varchar(10) NOT NULL DEFAULT '' COMMENT '股票编号',
  `c_date` date NOT NULL DEFAULT '1970-01-01',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `updated_at` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  PRIMARY KEY (`id`),
  UNIQUE KEY `icd_index` (`industry_code`,`individual_code`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COMMENT='行业和个股关联表';



create table bi_deal_detail
(
    deal_detail_id              int auto_increment comment '主键' primary key,
    bi_code varchar(10)   default ''                     not null comment '币code',
    bi_name varchar(20)   default ''                    not null comment '币名称',
    price_usd      float(10, 4)  default 0.0000      not null comment '价格',
    vol_usd       double(20, 2) default 0.00         not null comment '24小时交易额',
    turnover_ratio  float(10, 4)  default 0.0000   not null comment '换手率',
    rose_ratio      float(10, 4)  default 0.0000     not null comment '涨跌幅',
    c_date          date          default '1970-01-01'          not null,
    created_at      datetime      default '1970-01-01 00:00:00' not null,
    updated_at      datetime      default '1970-01-01 00:00:00' not null,
    constraint icd_index
        unique (deal_detail_id, c_date)
)
    comment '币交易明细' charset = utf8mb4;