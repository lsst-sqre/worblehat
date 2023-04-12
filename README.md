# Worblehat

Worblehat is a very simple WebDAV server designed for use in the Rubin
Observatory's Science Platform.  It is designed to cooperate with
[Gafaelfawr](https://gafaelfawr.lsst.io/) to provide access control and
a correct user account security context.

## Warning

Worblehat assumes that it is being run in a container with the
filesystems to be served mounted, and further that the container's UID,
primary GID, and supplemental groups are all set correctly to restrict
access to the files the container sees.

This is typically done by protecting the mechanism that spawns a
Worblehat container behind a Gafaelfawr ingress that can determine the
user it should run as and restrict access to users providing an
appropriately-scoped security token.

Running it as root without securing network access to it would be a very
bad idea, since Worblehat itself gives unauthenticated write access to
whatever it's serving.

