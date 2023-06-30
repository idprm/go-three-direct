CREATE UNIQUE INDEX uidx_service_msisdn ON yatta_db.subscriptions (service_id, msisdn);
CREATE INDEX idx_submited_id_msisdn ON yatta_db.transactions (submited_id, msisdn);