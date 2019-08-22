## API Authentication

API Authentication is based on a simple HMAC mechanism to prevent request forgery.


API credentials contain two components:

1. Public key
2. Secret key

The following headers are required:

`X-Public-Key`:  
The public key used to identify how this request is authorized

`X-Request-Signature <type> <signature>`:  
The signature type and signature value used to sign this request.

Currently SHA384 is the only supported signature method.

The request signature is an HMAC of the following payload data:

* Request path
* Current unix timestamp truncated to 5 minutes.
* Request body, up to the first kilobyte.

These values are joined without separators. 

The signature must be Base64 encoded.