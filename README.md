# helloname

Assumptions: 
Build a 3 tier web app. Build an API to accept json requests and take an action
In this case check GET endpoint hello/:somename and return and print the name as well as to lookup the name in a database. If the name exists, increase a counter tracking number of GET requests for that name. If the name doesn't exist then insert the name into a table and increase the names count by 1. To delete the counts on all names post a json request to /counts..(curl -X DELETE -H "Content-Type: application/json" http://104.154.136.98/counts). To return all results outputed in json tupe of the name and count go to http://104.154.136.98/counts. Any undefined route will return a 404. 
To see system data go to http://104.154.136.98/health which will return disk and memory info.

Tech:
Go
Docker
Kubernetes
Helm
Terraform
Postgres
Gcloud CLI

Buckets:
nyt-hello-tf //tfvars and state
ny-helloname-creds //json creds


Build and Deploy:
GKE Cluster can be bootstraped by downloading the .tfvars file to the ./gcp dir of this repo. CloudSQL databases prohibit the reuse of identical db intance names for a fixed period so update the {db_instance_name =} var in tfvars to a uniqe value to launch the db. The Makefile conists of 8 operations. 

1. make plan - Runs terraform plan to see what will be launched in GCE and outputs what will occurr during the run.

2. make create - Creates a GKE cluster along with a cloud sql postgres instance

3. make destroy - Destroys all previsouly created GKE clusters and cloud sql instances

4. make install - Initial install of the chart used for deploying to the cluster

5. make deploy - Upgrades installed chart of the helloname app

6. make run - Runs go app in go

7. make build - Docker builds helloname app

8. make push - Docker push helloname app to docker hub registry 

Enjoy!
