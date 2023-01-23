# Contributing Guidelines

## Upgrading

* change the `promtail` upstream version in Chart dependencies (`helm/promtail/Chart.yaml`)
* run `helm dependency update helm/promtail` to update the Chart.lock file
* re-generate `helm/promtail/values.schema.json`:
  * `helm schema-gen helm/promtail/values.yaml > helm/promtail/values.schema.json` to re-generate the file.
  * `sed -i 's/"type": "null"/"type": ["string", "null"]/g' helm/promtail/values.schema.json` to accept strings for all null values.
* update the link in the `Configuration` section of the Readmeto point to the new tag configuration.

## Testing

### Install the new Promtail version

* Choose a test cluster: `opsctl login gauss`
* Either create a draft PR or directly use your branch

#### The "easy" way

* Deploy the new Promtail version using opsctl (replace `test-branch` with your branch's name): `opsctl deploy --use-kubeconfig promtail-app@test-branch`
* Make sure promtail is deployed with the right version: `kubectl get app -n giantswarm promtail`

#### The "hard" way

If for some reason you can't use the `opsctl deploy` command or don't want to:
* Follow the instructions to manually deploy an app [here](https://intranet.giantswarm.io/docs/dev-and-releng/how-to-manually-update-an-app/)

### Ensure Promtail is scraped by prometheus

* Open prometheus on the MC from the installation: `opsctl open -i gauss -a prometheus --workload-cluster gauss`
* Go to the *Status->Targets* section and search for `promtail` in the search bar. You should see as many endpoints as there are nodes on the cluster (8/8 on gauss)
* You might want to take screenshots as it could prove useful for the PR.
* Go to the *Graph* section and execute the following query: `promtail_build_info`. As for the previous step, you should see as many metrics as there are nodes on the cluster (8 on gauss) with value 1.
* You might take screenshots as well.

### Ensure logs are shipped to Loki

* Open the grafana instance from gorilla: `opsctl open -i gorilla -a grafana`
* Go to the *Explore* section and select `Loki` as datasource.
* Execute a query based on the following example: `sum(count_over_time({installation="gauss", cluster_id="gauss", hostname=~"ip-.*.eu-west-1.compute.internal"}[10m])) by (hostname)`
* You should see as many different colors as there are nodes on the targeted cluster (8 for gauss). Take a screenshot for the PR.
* Execute another query based on this example: `{installation="gauss", cluster_id="gauss", pod=~"k8s-api-server-ip-.*"}`
* You should see logs from the api-server pods on the targeted cluster. Take a screenshot for the PR.

### Clean environment

* Once you're finished testing, simply resume flux: `flux resume kustomization -n flux-giantswarm collection; flux resume kustomization -n flux-giantswarm flux`. (You might consider not resuming flux when working on test installations as many users might need it to stay suspended)
