[![pipeline status](https://gitlab.com/milescrabill/demo/badges/master/pipeline.svg)](https://gitlab.com/milescrabill/demo/commits/master)

Our pipeline _very slightly_ changes Gitlab's [Auto Devops](https://docs.gitlab.com/ee/topics/autodevops/index.html) template by simply following the instructions to enable a staging environment and manual promotion to production.

I also deleted some Postgres stuff. But that doesn't really matter.

Really, the magic here is that with a simple git push Gitlab does all the work of:
    - building a docker image for that version
    - running tests on the code for that version
    - pushing docker images to its per-repo registry
    - rolling the deployment out to a kubernetes cluster for staging
    - after a manual trigger, rolling the deployment to production

It even monitors the application _and_ the cluster using Prometheus that _it installs itself_.
