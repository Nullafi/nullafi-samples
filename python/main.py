import http.client
import mimetypes
import os
from urllib.parse import urlencode

API_URL = 'http://api_url_here.com/api/scan-dynamic'
API_KEY = 'your_api_key_here'

def test_upload(file_name, content_type):
    # Path to the file to upload
    file_path = f'../sample-files/{file_name}'

    # Read file data
    with open(file_path, 'rb') as file:
        file_data = file.read()

    # Construct query parameters
    query_params = urlencode({
        'namespace': 'python.apiclient.upload.sample',
        'obfuscatedDataTypes': 'PHONE_NUMBER,EMAIL_ADDRESS',
        'maskFormats': 'LEAVE_LAST_FOUR,FULLY_OBFUSCATED',
        'username': 'python-sample-user',
        'usergroup': 'python-sample-usergroup',
        'storeOriginalValues': 'true',
    })

    # Construct the full URL with query parameters
    url_with_params = f'{API_URL}?{query_params}'

    # Create HTTP connection
    conn = http.client.HTTPConnection('localhost', 8080)

    # Prepare headers
    headers = {
        'Authorization': f'Bearer {API_KEY}',
        'Content-Type': content_type,
    }

    # Send HTTP request
    conn.request('POST', url_with_params, file_data, headers)

    # Get response
    resp = conn.getresponse()
    data = resp.read()

    # Save response body to file
    with open(f'response-files/response-{file_name}', 'wb') as response_file:
        response_file.write(data)

    print(f'Response Status: {resp.status}')
    print(f'Response Body Length: {len(data)}')
    print(f'Response body saved as response-{file_name}')

    conn.close()

# Test the file uploads
test_upload('sample-test.xlsx', 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet')
test_upload('sample-test.docx', 'application/vnd.openxmlformats-officedocument.wordprocessingml.document')