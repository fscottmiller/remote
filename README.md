# Remote

So, I guess what I'm thinking here is that I don't know of any containerized workstation management system. There's a lot of value in using containers as dev machines, and they're really not that hard to set up if you're a big docker dev, but what about management at scale? We could make it happen.

## Roots

Yeah, I'm definitely not planning to do this from scratch by myself when smart people have already written good code toward this. Off the top of my head, I'm thinking about starting with VSCode. I like VSCode for its extensibility. It's good for pretty much any language with all its plugins, and there's already an [awesome image](https://github.com/cdr/code-server) for running it remotely. 

Let's start with an instance of one user (running in a Kubernetes cluster, not on local Docker). At minimum, you'll need **one pod running the code-server image**. Well, that's easy enough to implement.

We'll also have to consider storage persistence. Do we want to mount volumes to these containers, or do we not want persistent storage? Will these containers spin up every time a user wants to access their workstation, or remain online all the time? I don't think we need to limit ourselves based on these answers - maybe they could each be an implementation? To list a few:
1. No persistence. When a user needs their workstation, it spins up and clones their repository, but when they log off, the whole thing goes away. This would definitely be the least expensive option, and probably the simplest to implement. The only issue I see is that there needs to be some way to configure preinstalled tools - it would be awful to need to reinstall every time.
2. Persistence! When a workstation is created, a volume is also created and mounted. Custom tools could be installed in that volume, which would persist between development sessions. You also wouldn't need to clone the repo every time, and you wouldn't have to commit whenever you log off. You could have one fairly large mount, which could have all your work on it, or several small ones - maybe one for each project?
3. We probably don't want the containers online all the time to save on resource utilization.

Since the first scenario is the simplest, I'd like to start working on that one. So -- when a user needs their workstation, it spins up, clones their repository (let's assume it's just one for now), and saves nothing across sessions. However, we've gotta come up with a mechanism for saving the configuration. Again for the sake of simplicity, I'm going to start with custom docker images that extend the code-server image. 

Oh, and I've written all of this in a remote container workstation, and plan to do the entire project in the same manner :)