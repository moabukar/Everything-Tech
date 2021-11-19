#### Documentation 

- [Falco docs](https://falco.org/docs/)
- [Falco supported field](https://falco.org/docs/rules/supported-fields/)

#### 1 - Simple rule

```sh

- rule: The program "curl" is run in a container
  desc: An event will trigger every time you run cat in a container
  condition: evt.type = execve and container.id != host and proc.name = cat
  output: "curl was run inside a container"
  priority: INFO

```

#### 2 - Added extra ouputt

```sh

- rule: The program "curl" is run in a container
  desc: An event will trigger every time you run cat in a container
  condition: evt.type = execve and container.id != host and proc.name = curl
  output: "curl was run inside a container(user=%username container=%container.name image=%container.image proc=%proc.cmdline)"
  priority: INFO

```

#### 3 - Adding Macros & List

```sh

- macro: custom_macro
  condition: evt.type = execve and container.id != host

- list: blacklist_binaries
  items: [curl,grep,wget,date]


- rule: The program "curl" is run in a container
  desc: An event will trigger every time you run cat in a container
  condition: evt.type = execve and container.id != host and proc.name = curl
  output: "curl was run inside a container(user=%username container=%container.name image=%container.image proc=%proc.cmdline)"
  priority: INFO

```
