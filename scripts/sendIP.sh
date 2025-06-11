ba_encoded_token="aHR0cHM6Ly9vYXBpLmRpbmd0YWxrLmNvbS9yb2JvdC9zZW5kP2Q5NzcyMWVkNzlmMmU3ZWU2NmYzNmExMWZkN2I4NTI1"
Ding_Webhook_Token=$(echo "$ba_encoded_token" | base64 -d)
IPIP="IP已变更"

sendDingIP() {
    curl -s "$Ding_Webhook_Token" \
    -H 'Content-Type: application/json' \
    -d '{
        "msgtype": "markdown",
        "markdown": {
            "title": "动态IP已变更",
            "text": "'"$*"'"
        },
        "at": {
            "isAtAll": false,
            "atMobile": ["19******52"]
        }
    }' 2>&1 > /dev/null
}