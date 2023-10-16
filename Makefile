
# Variables
PROJECT_ID ?= ai-trainer-dev


deployDependencies:
	# Enable cloud functions
	@gcloud services enable cloudfunctions.googleapis.com
	# get the framework for go
	@go get github.com/GoogleCloudPlatform/functions-framework-go

	# Tidying things up
	@go mod tidy
	@go mod vendor


gcloudInit:
	# Logging in to GCP and set it the application default
	@gcloud auth application-default login
	# Setting the project
	@gcloud config set project $(PROJECT_ID)


deploy:
	@gcloud functions deploy myfunction \
       --runtime go113 \
       --trigger-topic test-topic \
       --entry-point MyCloudEventFunction
