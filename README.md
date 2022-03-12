# rcu
Simple non-kernel-assisted (a.k.a userland) RCU library in Go.

Implementing RCU on garbage collected languages is actually quite simpler than
usual as we do not need to handle quiescent states (the garbage collector
handles that for us). So all we need is make atomic swaps so everything is
neatly lock-less but still thread-safe.

More on RCU can be read [here](https://en.wikipedia.org/wiki/Read-copy-update).
