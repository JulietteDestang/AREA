# Overview

This project is an automation platform between APIs. On a website usable on both mobile devices and classic PCs, you can create an account to select **actions** and **reactions**.

# Detailed explication

What is an **action** ? What is a **reaction** ?

## Action

An **action** is a [webhook](https://www.redhat.com/en/topics/automation/what-is-a-webhook) that will start a dedicated **reaction**.

### Some examples

- Weather action : triggering a reaction when the temperature in a specific city is below or above a specific number.
- Covid action : triggering a reaction when the covid cases or critical cases are over a given number.
- Crypto current : triggering a reaction when a choosen crypto is over or under a given number.

## Reaction

A **reaction** is a component activated by an **action** : these components perform a specific task by activating a **trigger**.

### Some examples 

- Deezer service : adding a specific song to a specific playlist.
- Discord service : sending a specific message in a specific channel.
- Spotify service : adding a specific song to the user's queue.

You can link action and reactions through the *wallet* page.

# Specs

The project is divided in three parts :
- An application server
- A web client
- A mobile client

<!-- @cond -->

## Languages

```mermaid
graph LR
A --> B
A((AREA)) --> C
B[App mobile] --> E{Flutter}
A --> F[API]
F --> G{golang}
C[App web] --> D{ReactJS}
```

## API description

```mermaid
graph LR
R --> S{Backend}
S --> R
S --> D((Database))
D --> S
M[Mobile] --> R{{Router}}
W[Web] --> R
```
<!-- @endcond -->

# about.json

The `about.json` is a file that contains informations about the client and the server (including active services).

This means that all active actions and reactions with their own descriptions are stored in this file.

The application server answers the call `http://localhost:8080/about.json` that leads to this file.  

***

# Our team

Developers  

| [<img src="https://github.com/Azzzen.png?size=85" width=85><br><sub>[Axel Zenine]</sub>](https://github.com/Azzzen) | [<img src="https://github.com/ErwanSimonetti.png?size=85" width=85><br><sub>[Erwan Simonetti]</sub>](https://github.com/ErwanSimonetti) | [<img src="https://github.com/JulietteDestang.png?size=85" width=85><br><sub>[Juliette Destang]</sub>](https://github.com/JulietteDestang) | [<img src="https://github.com/HKtueur1.png?size=85" width=85><br><sub>[Timothée De Boynes]</sub>](https://github.com/HKtueur1) | [<img src="https://github.com/BlanchoMartin.png?size=85" width=85><br><sub>[Martin Blancho]</sub>](https://github.com/BlanchoMartin)
| :---: | :---: | :---: | :---: | :---: |
