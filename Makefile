
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

test:
	@curl localhost:8080 \
	  -X POST \
	  -H "Content-Type: application/json" \
	  -H "ce-id: 123451234512345" \
	  -H "ce-specversion: 1.0" \
	  -H "ce-time: 2020-01-02T12:34:56.789Z" \
	  -H "ce-type: google.cloud.pubsub.topic.v1.messagePublished" \
	  -H "ce-source: //pubsub.googleapis.com/projects/MY-PROJECT/topics/MY-TOPIC" \
	  -d '{ \
			"message": { \
			  "data": "d29ybGQ=", \
			  "attributes": { \
				 "attr1":"attr1-value" \
			  } \
			}, \
			"subscription": "projects/MY-PROJECT/subscriptions/MY-SUB" \
		  }'
