@Library('libpipelines@master') _

hose {
    EMAIL = 'eos@stratio.com'
    MODULE = 'capsule'
    REPOSITORY = 'capsule'
    PKGMODULESNAMES = ['capsule']
    BUILDTOOL = 'make'
    NEW_VERSIONING = 'true'
    DEVTIMEOUT = 30
    ANCHORE_POLICY = "production"
    ANCHORE_TEST = true

    DEV = { config ->
        doDocker(config)
    }
}
