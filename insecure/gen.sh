openssl req \
-newkey rsa:2048 \
-x509 \
-nodes \
-keyout key.pem \
-new \
-out cert.pem \
-subj /CN=localhost \
-reqexts SAN \
-extensions SAN \
-config <(cat /usr/local/etc/openssl/openssl.cnf \
    <(printf '[SAN]\nsubjectAltName=DNS:localhost,IP:127.0.0.1')) \
-sha256 \
-days 3650
