# Showcase Project
This project is meant to allow developers showcase HTML applications and present thier work to potential employers. Using the free tier of Google Appengine in most cases.

## Environment
* [Install App Engine SDK](https://cloud.google.com/appengine/docs/standard/go/download "appengine")
* [Setup Cloud Storage](https://cloud.google.com/appengine/docs/standard/python/googlecloudstorageclient/setting-up-cloud-storage)

## Source code

``` sh
git clone https://github.com/darkmane/showcase.git
```

## Customization
If you have a domain, you will want to [configure App Engine](https://console.cloud.google.com/appengine/settings/domains) and DNS to direct requests to the application.
### Map Fully Qualified Host Names to directories in Cloud Storage default bucket

The application will map the host name to a directory in a Google Cloud Storage bucket.

It follows the following rules to map a hostname to an environment variable.
* Uppercase the hostname
* Remove '-' from the resulting string
* Replace '.' with '\_'

If the resulting string is not a valid environment variable, the app looks for a 'domain' query parameter. 

If neither of those result in a directory, then finally it uses the original hostname as the directory.

You can add environment variables to the `app.yaml` in the `env_variables` section.

### Deploy
From within the directory with your project:
`gcloud app deploy app.yaml`
