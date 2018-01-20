Inspired by https://aws.amazon.com/de/blogs/compute/announcing-go-support-for-aws-lambda/.

```
$ go test
$ go build -o main
$ zip lambda-go-playgound.zip main
$ aws lambda create-function \
--region eu-central-1 \
--function-name LambdaGoPlayground \
--zip-file fileb://./lambda-go-playgound.zip \
--runtime go1.x \
--role arn:aws:iam::<AccountID>:role/<LambdaExecutionRole> \
--handler main
```
Add an API Gateway trigger with
* API name: LambdaGoPlayground
* Deployment stage: dev
* Security: Open with access key
```
curl -k 'x-api-key:<APIKey>' -X POST -d 'Something' https://<APIID>.execute-api.<Region>.amazonaws.com/dev/LambdaGoPlayground
```