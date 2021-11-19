### Question - Falco 

A pod in the "space" namespace has produced an alert that a shell was opened in the container of this pod.

Find the falco rule that produces this alert and change the output to include the "user id", "container id", "container image repo"

### Solution

- [Falco](https://falco.org/docs/)

#### - Create a new falco rule

```sh

- Navigate to /etc/falco/falco_rules.yaml
- Find the rule titled "shell in a container"

Copy the rule to /etc/falco/falco_rules.local.yaml

```

```sh

- rule: Terminal shell in container
  desc: A shell was used as the entrypoint/exec point into a container with an attached terminal.
  condition: >
    spawned_process and container
    and shell_procs and proc.tty != 0
    and container_entrypoint
    and not user_expected_terminal_shell_in_container_conditions
  output: >
    %evt.time.s,%user.uid,%container.id,%container.image.repository
  priority: ALERT
  tags: [container, shell, mitre_execution]
  Save the /etc/falco/falco_rules.local.yaml file.
  
```

```sh

Then run systemctl restart falco.service to override the current rule

```
