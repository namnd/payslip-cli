# Employee Monthly Pay Slip

## Assumption
1. All amount are rounded up to dollar
1. Name must have double quote if more than one word
1. One single word name can omit double quote

## Build docker image
~~~
docker build . --tag=payslip-cli:latest
~~~

## Run application
~~~
docker run -it --rm payslip-cli
~~~
