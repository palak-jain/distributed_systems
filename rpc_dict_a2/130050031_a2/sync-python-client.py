import requests
import json

def main():
    url = 'http://localhost:6000/rpc'
    headers = {'content-type': 'application/json'}

    # Example echo method
    payload = json.dumps({
        'method': 'Dict.InsertWord',
        'params': [{'Key' : 'Roy'}],
        'jsonrpc': '2.0',
        'id': 0,
    })
    response = requests.post(
        url, data=payload, headers=headers)

    print (response)
   

if __name__ == '__main__':
    main()

