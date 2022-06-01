@Library('libpipelines') _

hose {
    EMAIL = 'eos@stratio.com'
    BUILDTOOL = 'make'
    DEVTIMEOUT = 30
    ANCHORE_POLICY = "production"
    BUILDTOOL_IMAGE = "stratio/keos-builder:0.1.0-e97b880"
    VERSIONING_TYPE = 'stratioVersion-3-3'
    UPSTREAM_VERSION = '0.1.0'

    DEV = { config ->
        doDocker(conf:config, image:'capsule')
    }
}
