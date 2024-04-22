# Code Samples

This repository has examples on Golang, Nodejs and Python on how to consume our API to Dynamically scan Excel Spreadsheets and Word Docs.

# API File Upload and Query String Example

This Go program demonstrates how to consume the Nullafi API for file uploads for Data Type detection and obfuscation. 
It includes examples of uploading excel spreadsheets and word docs, setting the proper Content-Type header, including query strings in the API request, and saving the API response.

## Usage

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/Nullafi/nullafi-samples.git
   cd nullafi-samples
   ```

2. Modify the code example file with the API endpoint, API key, and query string parameters based on your environment and use-case

3. Run the program

4. The program will upload the sample files to the API endpoint, set the proper Content-Type header based on the file type, include query strings in the API request, and save the response files in the `response-files` folder.

## Sample Files

- `sample-test.xlsx`: Example Excel file for testing.
- `sample-test.docx`: Example Word document for testing.

## Query String Parameters

The following query string parameters are included in the API request:

- `namespace`: Namespace for the uploaded files (e.g., "go.apiclient.upload.sample").
- `obfuscatedDataTypes`: Types of data to obfuscate in the uploaded files (e.g., "PHONE_NUMBER,EMAIL_ADDRESS").
- `maskFormats`: Masking formats for obfuscated data (e.g., "LEAVE_LAST_FOUR,FULLY_OBFUSCATED") 
  - Note: Mask Formats matches in same order with the `obfuscatedDataTypes` so in the example "PHONE_NUMBER" will be "LEAVE_LAST_FOUR" and "EMAIL_ADDRESS" will be "FULLY_OBFUSCATED".
- `username`: Username for authentication (e.g., "golang-sample-user").
- `usergroup`: Usergroup for authorization (e.g., "golang-sample-usergroup").
- `storeOriginalValues`: Whether to store original values in our activity tracking (e.g., "true").

## Notes

- Replace `"http://api_url_here.com/api/scan-dynamic"` with your the API url setup on your environment.
- Update `"your_api_key_here"` with your API key for authorization.
- Modify the query string parameters as needed for your API requirements.
- The Content-Type headers are set based on the file types being uploaded (`application/vnd.openxmlformats-officedocument.spreadsheetml.sheet` for Excel, `application/vnd.openxmlformats-officedocument.wordprocessingml.document` for Word).
- The response files are saved in the `response-files` folder with the corresponding file extensions.

Feel free to adjust the instructions, file paths, API details, and query string parameters as needed for your specific setup.
