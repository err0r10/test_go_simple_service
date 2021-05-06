# test_go_simple_service

1. Copy ```git clone https://github.com/err0r10/test_go_simple_service.git```
2. ```cd test_go_simple_service```
3. Build ```docker build -t test_go_simple_service .```
4. Run docker container ```docker run -d -p 5001:5001 test_go_simple_service```
5. Run curl ```curl --location --request POST 'http://127.0.0.1:5001/you' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "YouName"
}'```
