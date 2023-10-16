# Cloud Functions test project

This example was built on MacOs. The procedure should be fairly similar on other OS, but you may have to update shell commands.

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

### Test the HTTP function

```FUNCTION_TARGET=HelloWorld LOCAL_ONLY=true go run cmd/main.go```

After from another terminal:

```curl localhost:8080```

### Test the CloudEventFunction

```FUNCTION_TARGET=MyCloudEventFunction LOCAL_ONLY=true go run cmd/main.go```

This is an example of a Cloud Storage event

```
curl localhost:8080 -v \
  -X POST \
  -H "Content-Type: application/json" \
  -H "ce-id: 123451234512345" \
  -H "ce-specversion: 1.0" \
  -H "ce-time: 2022-12-31T00:00:00.0Z" \
  -H "ce-type: google.cloud.storage.object.v1.finalized" \
  -H "ce-source: //storage.googleapis.com/projects/_/buckets/image_bucket" \
  -H "ce-subject: objects/CF_debugging_architecture.png" \
  -d '{
        "bucket": "image_bucket",
        "contentType": "text/plain",
        "kind": "storage#object",
        "md5Hash": "...",
        "metageneration": "1",
        "name": "CF_debugging_architecture.png",
        "size": "352",
        "storageClass": "MULTI_REGIONAL",
        "timeCreated": "2022-12-31T00:00:00.0Z",
        "timeStorageClassUpdated": "2022-12-31T00:00:00.0Z",
        "updated": "2022-12-31T00:00:00.0Z"
      }'
```
