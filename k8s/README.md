Worblehat Fileserver Kubernetes Objects
=======================================

These are representative templates for the fileservers' objects.

Note that we presume the use of GafaelfawrIngress, which is a Rubin
Observatory-specific CRD that manages templated Ingress resources.

You would have to do a bit of work to decouple Worblehat from its
reliance on Gafaelfawr.

In reality, all these objects are constructed on the fly by
[JupyterLab Controller](https://github.com/lsst-sqre/jupyterlab-controller/)
rather than existing as standalone Kubernetes YAML.
