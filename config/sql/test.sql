-- 查询板块交易额
select ib.industry_code,
       ib.industry_name,
       sum(fund_real_in) as fund_real_in,
       sum(ob_price)     as ob_price,
       sum(circulate_value) as circulate_value,
       count(*)          as count_num,
       ob_price / circulate_value as contact_ratio
from industry_bankroll ib
         inner join individual_stock s on ib.industry_code = s.industry_code
where ib.c_date between date_sub(DATE_FORMAT(NOW(), '%Y-%m-%d'), INTERVAL 0 DAY) and date_sub(DATE_FORMAT(NOW(), '%Y-%m-%d'), INTERVAL 0 DAY)
group by ib.industry_code
order by contact_ratio desc;
