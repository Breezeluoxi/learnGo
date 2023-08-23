update `flow_deal_data` set `seq` = `seq`+1 where `FPI_W1IVD_N3FLOYW` = '?';

insert into `flow_deal_data`
(`FPI_W1IVD_N3FLOW`,`FPI_W7PROCESSC_R4CODER`,`seq`,`FPI_D1PDROCESS_A1SOEQUE`)
values
('?','A01',1,1);