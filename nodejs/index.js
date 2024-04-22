const http = require('http');
const fs = require('fs');
const { URLSearchParams } = require('url');

const API_URL = 'http://api_url_here.com/api/scan-dynamic';
const API_KEY = 'your_api_key_here';

function testUpload(fileName, contentType) {
    // Path to the file to upload
    const filePath = `../sample-files/${fileName}`;

    // Read file data
    const fileData = fs.readFileSync(filePath);

    // Construct query parameters
    const queryParams = new URLSearchParams({
        namespace: 'node.apiclient.upload.sample',
        obfuscatedDataTypes: 'PHONE_NUMBER,EMAIL_ADDRESS',
        maskFormats: 'LEAVE_LAST_FOUR,FULLY_OBFUSCATED',
        username: 'node-sample-user',
        usergroup: 'node-sample-usergroup',
        storeOriginalValues: 'true',
    });

    // Construct the full URL with query parameters
    const urlWithParams = `${API_URL}?${queryParams.toString()}`;

    // Configure the request options
    const options = {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${API_KEY}`,
            'Content-Type': contentType,
        },
    };

    // Create an HTTP request
    const req = http.request(urlWithParams, options, (res) => {
        const chunks = [];

        res.on('data', (chunk) => {
            chunks.push(chunk);
        });

        res.on('end', () => {
            const responseData = Buffer.concat(chunks);

            console.log(`Response Status: ${res.statusCode}`);
            console.log(`Response Body Length: ${responseData.length}`);

            // Save response body as file
            fs.writeFileSync(`response-files/response-${fileName}`, responseData);
            console.log(`Response body saved as response-${fileName}`);
        });
    });

    // Handle request errors
    req.on('error', (error) => {
        console.error('Error sending HTTP request:', error);
    });

    // Send file data in the request
    req.write(fileData);

    // End the request
    req.end();
}

// Test the file uploads
testUpload('sample-test.xlsx', 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet');
testUpload('sample-test.docx', 'application/vnd.openxmlformats-officedocument.wordprocessingml.document');
