package database

import (
	"time"

	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

var configs = []model.Config{
	{
		Name:  "AUTH_USER",
		Value: "tU1qzr659H6VG3zGGwz38dIApGGVMmrY",
	},
	{
		Name:  "AUTH_PASS",
		Value: "RwrHIWKfNanVhdlN",
	},
	{
		Name:  "PARTNER_ID",
		Value: "linkit",
	},
	{
		Name:  "PRODUCT_ID",
		Value: "KoTest",
	},
	{
		Name:  "TRANSACTION_ID",
		Value: "KoTest123",
	},
	{
		Name:  "CHARGABLE_AMOUNT",
		Value: "2",
	},
	{
		Name:  "CORRELATION_ID",
		Value: "123",
	},
	{
		Name:  "COOLING_PERIOD",
		Value: "10",
	},
	{
		Name:  "QUARANTINE_DAY",
		Value: "10",
	},
	{
		Name:  "RENEWAL_DAY",
		Value: "2",
	},
	{
		Name:  "TRIAL_DAY",
		Value: "1",
	},
}

var schedules = []model.Schedule{
	{
		Name:       "RENEWAL_PUSH",
		PublishAt:  time.Now(),
		UnLockedAt: time.Now().Add(time.Hour),
		Status:     true,
	},
	{
		Name:       "RETRY_PUSH",
		PublishAt:  time.Now(),
		UnLockedAt: time.Now().Add(time.Hour),
		Status:     true,
	},
}

var contents = []model.Content{
	{
		Name:       "WELCOME",
		OriginAddr: "998790",
		Value:      "REG KEREN kamu sdh aktif. Kamu akan dikirimkan SMS utk akses layanan tarif 2200/SMS/2 hari, aktif s/d 180 hari. Stop: UNREG KEREN ke 99879 CS:02152964211",
	},
	{
		Name:       "REGISTRATION",
		OriginAddr: "998790",
		Value:      "Kamu akan berlangganan layanan REG KEREN tarif 2200/sms/2 hari, layanan aktif s/d 180hr. Balas YA utk lanjut.",
	},
	{
		Name:       "CONFIRMATION",
		OriginAddr: "998790",
		Value:      "Terimakasih, permintaan kamu diproses",
	},
	{
		Name:       "FIRSTPUSH",
		OriginAddr: "998791",
		Value:      "Kamu terdaftar di REG KEREN tarif 2200/sms/2 hari. Klik aplikasi REG KEREN https://bit.ly/3BGcVgj. (Tarif data berlaku). Stop: UNREG KEREN ke 99879 CS:02152964211",
	},
	{
		Name:       "RENEWAL",
		OriginAddr: "998791",
		Value:      "Layanan KEREN km aktif 180hr s/d @renewal_date di https://bit.ly/3BGcVgj. Tarif Rp2200/sms/2 hari selama 180hr. Stop: UNREG KEREN,CS:021-52964211",
	},
	{
		Name:       "UNSUB",
		OriginAddr: "998790",
		Value:      "Kamu sudah tidak berlangganan di layanan REG KEREN",
	},
	{
		Name:       "INSUFT",
		OriginAddr: "998790",
		Value:      "Maaf pulsa kamu tdk cukup, mohon isi ulang untuk bisa menikmati serunya layanan kami. Stop: UNREG KEREN ke 99879 CS:02152964211",
	},
	{
		Name:       "ERROR_KEYWORD",
		OriginAddr: "998790",
		Value:      "Keyword yang kamu masukkan salah. Ketik REG KEREN ke 99879. CS: 02152964211",
	},
	{
		Name:       "FAILED",
		OriginAddr: "998790",
		Value:      "Maaf, Anda belum berhasil berlangganan",
	},
	{
		Name:       "REMINDER",
		OriginAddr: "998790",
		Value:      "Layanan REG KEREN akan berakhir. Untuk perpanjang balas YA ke 99879 tarif 2200/SMS/2 hari, layanan aktif s/d 180 hari",
	},
	{
		Name:       "IS_ACTIVE",
		OriginAddr: "998790",
		Value:      "Kamu masih terdaftar di layanan REG KEREN tarif 2200/SMS/2 hari. Untuk berhenti berlangganan ketik UNREG KEREN kirim ke 99879",
	},
	{
		Name:       "PURGE",
		OriginAddr: "998790",
		Value:      "Kamu sudah tidak berlangganan di layanan REG KEREN",
	},
}

var services = []model.Service{
	{
		ID:              1,
		Code:            "99879",
		Name:            "KEREN",
		AuthUser:        "SD_210906_0180",
		AuthPass:        "y4tt43r4",
		Day:             2,
		Charge:          2200,
		TrialDay:        0,
		UrlNotifSub:     "https://tri.fortune360.mobi/api/subscription/subscribe",
		UrlNotifUnsub:   "https://tri.fortune360.mobi/api/subscription/unsubscribe",
		UrlNotifRenewal: "https://tri.fortune360.mobi/api/subscription/renewal",
		UrlPostback:     "http://kbtools.net/id-yatta-h3i.php",
		IsActive:        true,
	},
}

var keywords = []model.Keyword{
	{
		Name: "YT",
	},
	{
		Name: "YT1",
	},
	{
		Name: "YT2",
	},
	{
		Name: "YT3",
	},
}
