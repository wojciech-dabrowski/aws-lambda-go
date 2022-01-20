.PHONY: build

build:
	sam build

deploy:
	sam build && sam deploy --template-file ./.aws-sam/build/template.yaml --stack-name aws-lambda-go --s3-bucket $(DEPLOY_BUCKET) --s3-prefix calculate-student-grade/ --capabilities CAPABILITY_IAM --no-fail-on-empty-changeset
