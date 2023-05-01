# Worblehat

Worblehat is a very simple WebDAV server designed for use in the Rubin
Observatory's Science Platform.  It is designed to cooperate with
[Gafaelfawr](https://gafaelfawr.lsst.io/) to provide access control and
a correct user account security context.

## Deployment

Worblehat is designed to be run under Kubernetes.  At Rubin Observatory,
Worblehat fileservers are managed on users' behalfs as part of
[Nublado](https://github.com/lsst-sqre/phalanx/tree/main/applications/nublado)
and the specific objects created to support user fileservers are part
of [JupyterLab Controller](https://github.com/lsst-sqre/jupyterlab-controller).

Some representative Kubernetes YAML is available [here](./k8s) but be
aware that in reality, these objects are created on the fly by the
controller.

For each user fileserver, we construct a Job, which manages a single Pod
that contains the Worblehat fileserver.  We create a Service in front of
that Pod, and we create a GafaelfawrIngress, which in turn manages an
Ingress, that points to that Service.  Part of the Controller machinery
notes when the Pod exits, which causes Job completion, and cleans up the
remaining user objects.

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

