# Insecure

This folder contains a script to generate a self-signed cert key pair, which is used to boot the web server with TLS. Please note in this project gRPC-Web API endpoints requires HTTP2, which means TLS is mandatory.

When the browser asks for security confirmation, you may also need to add `cert.pem` to the OS key chain to proceed.

The key pair is only valid for DNS Name localhost. In real deployment, you should use a solution like [Let's Encrypt](https://letsencrypt.org/).
