# Share

## Setup:
```
export AWS_ACCESS_KEY_ID=tid_xxx
export AWS_ENDPOINT_URL_S3=https://fly.storage.tigris.dev
export AWS_REGION=auto
export AWS_SECRET_ACCESS_KEY=tsec_xxx
export BUCKET_NAME=XXX
export DOMAIN=
```

## Usage:
```sh
# Share a file
share accounting-q3.zip

# Share a snippet:
echo "hello" | share

# Perfect for Taco Bell Programming
awk '{ print $1 }' access.log | sort | uniq -c | sort -nr | head -10 | share
```
