@Library('libpipelines') _

hose {
    EMAIL = 'platform@stratio.com'
    BUILDTOOL_IMAGE = 'golang:1.22'
    BUILDTOOL = 'make'
    DEVTIMEOUT = 30
    ANCHORE_POLICY = "production"
    VERSIONING_TYPE = 'stratioVersion-3-3'
    UPSTREAM_VERSION = '0.7.2'
    DEPLOYONPRS = true
    GRYPE_TEST = true

    DEV = { config ->
        doDocker(conf:config, image:'capsule')
        doHelmChart(conf: config, helmTarget: "chart")
    }
}
