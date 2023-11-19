import tls_client
import requests

# session = tls_client.Session(client_identifier="chrome112")
session = requests.Session()

# url = "https://mktapi.autotrader.ca/api/ad/search/v2"

# querystring = {"hcp":"true","pr":",","haspr":"true","q":"dodge challenger","fcpf":"false","c2":"7,11,10,9","t":"10","p":"0","ptradloc":"100,T5A+2Y2","tapp":"true"}
url = "https://mktapi.autotrader.ca/api/ad/search/v2?c2=7,9,10,11&ptradloc=50,T5A%202Y2&haspr=True&hcp=True&q=dodge%20challenger&srt=2&t=20&p=0"

payload = ""
headers = {
    # "cookie": "visid_incap_1710137=imI%2B8LBCRJaM1OQaK8OIm6NkWWUAAAAAQUIPAAAAAAAiSaY71U3vUSekysVuHcX6; incap_ses_8220_1710137=%2FkABMMsNYEmAMMwCXk8TcqNkWWUAAAAAAe0nvEiDYBjYQV9WNLP1ZA%3D%3D",
    "User-Agent": "samsung SM-G955N|Android 7.1.2|8.14.0.459",
    "Authorization": "Basic MjE4MzA1MTg2Njo6",
    "Content-Type": "application/json; charset=utf-8",
    "Accept-Language": "en-ca"
}

response = session.get(url, data=payload, headers=headers, allow_redirects=True)

print(response.text)

