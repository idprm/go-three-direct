Go Yatta H3I

# SERVER

sudo nano /etc/systemd/system/server.service

[Unit]
Description=go-yatta-h3i server

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/app/go-yatta-h3i
ExecStart=/app/go-yatta-h3i/go-yatta-h3i server

[Install]
WantedBy=multi-user.target
=======================================================

# RENEWAL | Publisher

sudo nano /etc/systemd/system/publisher-renewal.service

[Unit]
Description=go-yatta-h3i publisher-renewal

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/app/go-yatta-h3i
ExecStart=/app/go-yatta-h3i/go-yatta-h3i publisher-renewal

[Install]
WantedBy=multi-user.target
=======================================================

# RETRY | Publisher

sudo nano /etc/systemd/system/publisher-retry.service

[Unit]
Description=go-yatta-h3i publisher-retry

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/app/go-yatta-h3i
ExecStart=/app/go-yatta-h3i/go-yatta-h3i publisher-retry

[Install]
WantedBy=multi-user.target
=======================================================

# RENEWAL | Consumer

sudo nano /etc/systemd/system/consumer-renewal@.service

[Unit]
Description=go-yatta-h3i consumer-renewal %i

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/app/go-yatta-h3i
ExecStart=/app/go-yatta-h3i/go-yatta-h3i consumer-renewal %i

[Install]
WantedBy=multi-user.target
=======================================================

# RETRY | Consumer

sudo nano /etc/systemd/system/consumer-retry@.service

[Unit]
Description=go-yatta-h3i consumer-retry %i

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/app/go-yatta-h3i
ExecStart=/app/go-yatta-h3i/go-yatta-h3i consumer-retry %i

[Install]
WantedBy=multi-user.target
=======================================================

# MO | Consumer

sudo nano /etc/systemd/system/consumer-mo@.service

[Unit]
Description=go-yatta-h3i consumer-mo %i

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/app/go-yatta-h3i
ExecStart=/app/go-yatta-h3i/go-yatta-h3i consumer-mo %i

[Install]
WantedBy=multi-user.target
=======================================================

# DR | Consumer

sudo nano /etc/systemd/system/consumer-dr@.service

[Unit]
Description=go-yatta-h3i consumer-dr %i

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/app/go-yatta-h3i
ExecStart=/app/go-yatta-h3i/go-yatta-h3i consumer-dr %i

[Install]
WantedBy=multi-user.target
=======================================================
