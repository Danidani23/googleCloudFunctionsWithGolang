# Cloud Functions test project

## How to work in local
All necessary commands are provided in the Makefile. Use them in this order
1. deployDependencies
2. gcloudInit


## Documentation

[GCP official docs](https://cloud.google.com/functions/docs/writing#entry-point)

[How to call functions directly](https://cloud.google.com/functions/docs/running/direct)

[Small project example in nodeJs](https://cloud.google.com/blog/topics/developers-practitioners/how-to-develop-and-test-your-cloud-functions-locally)

[Code snippets from GIThub with GO](https://github.com/GoogleCloudPlatform/functions-framework-go#quickstart-hello-world-on-your-local-machine)

## Local run

### How to use the same permissions as if it were deployed in the Cloud

In the Makefile there is the command to set the default account.
If you want to impersonate a service account, user:

```gcloud auth application-default login --impersonate-service-account=[YOUR_SERVICE_ACCOUNT_EMAIL]```