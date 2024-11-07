To run this aplication flow the below steps 
# Clone the repository
git clone git@github.com:bshinde/golangAssignment.git

# Navigate to the project directory
cd golangAssignment

# Install dependencies
go mod tidy

# Run the application
go run main.go

# curl request for testing the API
curl --location 'http://localhost:8080/validate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Name": "Test Name",
    "Pan": "ABCDE1234F",  
    "Mobile": "9876543210",
    "Email": "test@example.com"
}'
