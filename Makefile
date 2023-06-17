export

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
# Phony avoids conflicts for the file named as the main command
.PHONY: help
# Adds files to staging the commits the changes and finally push to remote repository

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

gitPush: ### Command to aid in pushing code to the current branch. Runs `git add . && git commit && git push
	chmod +x ./cmd/gitPush.sh
	./cmd/gitPush.sh
.PHONY: gitPush

mockeryGenerateBlogUsecase: ### generates testing mocks using mockery tool
	mockery --dir=internal/entity/utilities --name=PaymentUsecase --filename=payment.go --output=internal/entity/mocks --outpkg=mocks
.PHONY: mockeryGenerateBlogUsecase

testCover: ### Used to run tests with coverage and display the output.Scans all the files and runs the tests if available
	go test ./... -coverprofile=cover.out
.PHONY: goTestCoverProfile
