#!/usr/bin/env python
# -*- coding: utf-8 -*-

import string
import requests
from base64 import b64encode

def auth(key):
    source = "api:" + key
    auth = "Basic " + b64encode(source)
    headers = {"Authorization": auth}
    return headers


def main():

    key = "d5e76748-7b9a-11e9-bb91-f21898b25098"

    data = {
        "from": "user@example.hedwi.com>",
        "to": "support@hedwi.com",
        "subject": "Hello",
        "text": "Testing email from Hedwi!"
    }
    url = "https://api.hedwi.com/mail/example.hedwi.com"

    use_headers = True
    if use_headers:
        headers = auth(key)
        #使用headers
        r = requests.post(url, headers = headers, data = data)
        print r.content
    else:
        #使用requests basic auth
        r = requests.post(url, auth = ("api", key), data = data)
        print r.content

if __name__ == "__main__":
    main()
