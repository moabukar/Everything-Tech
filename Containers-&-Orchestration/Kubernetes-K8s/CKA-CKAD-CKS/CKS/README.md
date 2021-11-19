[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

# CKS-Exercises

<p align="center">
  <img width="360" src="Cks-logo.png">
</p>

A curated collections of exercises to help prepare for the Certified Kubernetes Security Specialist. The exercises have been segregated into their respective domains as per the [CNCF curriculum](https://training.linuxfoundation.org/certification/certified-kubernetes-security-specialist/) for CKS.

### Note: 

- This repo is not meant to be an exam dump or a simulated exam. It is meant to be used as a guide for the type of exercises associated with each domain of the exam. We also provide exam like questions which give you an insight into what type of questions you will be asked and how you'll be able to solve such a question.


## Contributing

### Pull Request process
  - Fork the repo and create your branch from your forked repo
  - Please try to stick to the layout we follow in terms of the README markdown
  - Issue that PR
  - This will be reviewed by the team and merged

## Shortcuts and things to keep in mind when going through this repo

- NS = Namespace
- SA = Service account
- Po = Pod
- NetPol = Network policy
- PSP = Pod security policy
- RBAC = Role-based access control
- k = kubectl
- SVC = Service

## [Cluster setup (10%)](1-cluster-setup/)

- Using network policies to restrict cluster level access
- Using CIS benchmark to review security configuration of k8s components (kube-api, kubelet, etcd, kubedns) — Kube-bench
- Properly setting up ingress objects with security controls
- Protecting node metadata and endpoints
- Minimising the use of, and access to, GUI elements
- Verify platform binaries before deployment — binary verification using sha512 hash

## [Cluster Hardening (15%)](2-cluster-hardening/)

- Restricting access to the Kubernetes API
- Using RBAC (role-based access controls) to minimise exposure
- Using service account with minimal permissions and disabling defaults — opting out of automounting credentials for service accounts
- Keeping up to date with the latest Kubernetes versions

## [System Hardening (15%)](3-system-hardening/)

- Minimising footprint of host OS to reduce attack surface
- Minimising IAM roles: Principle of least privilege — Authentication and Authorisation
- Minimise external access to the network — denying external access to outside the cluster
- Using kernel hardening tools — such as seccomp and AppArmor correctly

## [Minimise Microservice Vulnerabilities (20%)](4-minimise-microservice-vulnerabilities/)

- Setting up appropriate OS level security domains like PSP, OPA and security contexts
- Managing K8s secrets
- Use container runtime sandboxes in multi-tenant environments like gVisor and kata-containers.
- Implementation of pod-to-pod encryption by using mTLS configurations

## [Supply Chain Security (20%)](5-supply-chain-security/)

- Minimise base image footprint — distroless, alpine and an image relevant to your build as well as following best practices when creating containers
- Securing your supply chain by signing and validating images and a whitelist of allowed image registries — ImagePolicyWebhook admission controller.
- Scan images for known vulnerabilities — Aquasec Trivy

## [Monitoring, Logging and Runtime security (20%)](6-monitoring-logging-runtime-security/)

- Performing analytics of syscall processes at host and container level to detect malicious activities — [Falco](https://falco.org/docs/)
- Detect threats within physical infrastructure, apps, networks, data, users and workloads
- Detect all phase of attack regardless of where it happens and its spread
- Performing deep analytical investigation and identification of bad actors within environment — [Sysdig](https://sysdig.com/)
- Ensuring immutability of containers at runtime
- Audit logs to monitor access

## [Exam style questions](7-mock-exam-questions)

- We have added some practice questions for you guys here
- Keep in mind these are not an exact copy of exam questions rather they are guidance for you to follow based on our K8s learning and experiences.

## [Are you preparing for the CKA/CKAD?](8-preparing-for-CKA-CKAD?)

- Here are a list of online resources to help you pass the CKA/CKAD exams so that you can go for the CKS (as the CKA is a prerequisite for the CKS)
- Feel free to make a PR in the folder if you find useful resources.

## More useful resources and materials

#### Further reading (articles & books)

1. [Where to begin with the CKS exam?](https://moabukar.medium.com/where-to-begin-with-the-cks-exam-5cf0dcc86f76)

#### Slack channels

1. [Kubernetes Community - #cks-exam-prep](https://kubernetes.slack.com)

#### K8s and Container Security resources

1. [Mumshad's KodeKloud "Certified Kubernetes Security Specialist" course](https://kodekloud.com/p/certified-kubernetes-security-specialist-cks)
1. [Killer.sh CKS practice exam](https://killer.sh/cks)

#### Other useful CKS repos

1. [Walid Shaari's CKS repo](https://github.com/walidshaari/Certified-Kubernetes-Security-Specialist)
1. [Abdennour's CKS repo](https://github.com/abdennour/certified-kubernetes-security-specialist)
1. [A simple and comorehensive CLI to prepare for the CKS exam by Abdennour](https://cks.kubernetes.tn/getting-started/installation/) 
1. [Kim's CKS Challenge series](https://github.com/killer-sh/cks-challenge-series)
