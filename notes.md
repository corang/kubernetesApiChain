# Purpose
To allow initContainer based dependency chaining of containers regardless of k8s environment (istio + mysql makes me want to rip my hair out). For this to work the readinessProbe(s) of the container(s) in the pod you're dependent on have to be functional.

# Notes
- Should use a specific serviceAccount
- ClusterRole should only have exactly the permissions needed to view the status of pods in it's own namespace.
  - Experiment with how restrictive permissions can get
- ClusterRoleBinding only refers to the new clusterRole and serviceAccount
- image should be extremely small to leave as small a footprint as possible

## Manifests
- Helm chart?

## Go app
- configurable entirely through environment variables
  - amount of retries
    - defaults to infinite
  - delay for retries
    - defaults to `10s`
  - name of pod
    - required
  - configurable time delay after api success
    - defaults to none (`0s`)
    - takes format of linux `sleep` command
- https://12factor.net/