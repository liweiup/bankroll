package service

import "bankroll/global"

//行业资金
//`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
//`fund_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型，1表示概念，2表示行业',
//`day_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '几天时间',
//`industry_code` int(5) NOT NULL DEFAULT '0' COMMENT '行业编号',
//`industry_name` varchar(20) NOT NULL DEFAULT '' COMMENT '行业名称',
//`industry_index` float(8,2) NOT NULL DEFAULT '0.00' COMMENT '行业指数',
//`rose_ratio` float(4,2) NOT NULL DEFAULT '0.00' COMMENT '涨跌幅',
//`fund_amount_out` double(12,2) NOT NULL DEFAULT '0.00' COMMENT '流出资金',
//`fund_amount_in` double(12,2) NOT NULL DEFAULT '0.00' COMMENT '流入资金',
//`fund_real_in` double(12,2) NOT NULL DEFAULT '0.00' COMMENT '净额',
//`company_num` smallint(5) NOT NULL DEFAULT '0' COMMENT '公司家数',
//`leader_company_name` varchar(20) NOT NULL DEFAULT '' COMMENT '领涨股票名',
//`leader_rose_ratio` float(4,2) NOT NULL DEFAULT '0.00' COMMENT '领涨股涨幅',
//`leader_price` float(6,2) NOT NULL DEFAULT '0.00' COMMENT '领涨股当前价',
type IndustryBankroll struct {
	global.GVA_MODEL
	FundType int `json:"fund_type"`
	DayNum int `json:"day_num"`
	IndustryCode int `json:"industry_code"`
	IndustryName string `json:"industry_name"`
	IndustryIndex     float64 `json:"industry_index"`
	RoseRatio         float64 `json:"rose_ratio"`
	FundAmountOut     float64 `json:"fund_amount_out"`
	FundAmountIn      float64 `json:"fund_amount_in"`
	FundRealIn        float64 `json:"fund_real_in"`
	CompanyNum        int     `json:"company_num"`
	LeaderCompanyName string  `json:"leader_company_name"`
	LeaderCompanyCode int     `json:"leader_company_code"`
	LeaderRoseRatio   float64 `json:"leader_rose_ratio"`
	LeaderPrice       float64 `json:"leader_price"`
	CDate             string `json:"c_date"`
}
//个股资金流
//`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
//`day_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '几天时间',
//`symbol_code` int(5) NOT NULL DEFAULT '0' COMMENT '股票编号',
//`symbol_name` varchar(20) NOT NULL DEFAULT '' COMMENT '股票名称',
//`price` float(8,2) NOT NULL DEFAULT '0.00' COMMENT '当前价',
//`rose_ratio` float(4,2) NOT NULL DEFAULT '0.00' COMMENT '涨跌幅',
//`hand_ratio` float(4,2) NOT NULL DEFAULT '0.00' COMMENT '换手率',
//`fund_amount_out` double(12,2) NOT NULL DEFAULT '0.00' COMMENT '流出资金',
//`fund_amount_in` double(12,2) NOT NULL DEFAULT '0.00' COMMENT '流入资金',
//`fund_real_in` double(12,2) NOT NULL DEFAULT '0.00' COMMENT '净额',
//`trade_amount` double(12,2) NOT NULL DEFAULT '0.00' COMMENT '净额',
//`big_trade_in` double(12,2) NOT NULL DEFAULT '0.00' COMMENT '大单流入'
type IndividualBankroll struct {
	global.GVA_MODEL
	FundType int
	DayNum int
	IndustryCode int
	IndustryName string
	IndustryIndex float32
	RoseRatio float32
	FundAmountOut float64
	FundAmountIn float64
	FundRealIn float64
	CompanyNum int
	LeaderCompanyName string
	LeaderRoseRatio float32
	LeaderPrice float32
}