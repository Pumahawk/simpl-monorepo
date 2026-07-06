## IntellJ IDEA Run Configurations

These run configurations are exported from a multi-module IntelliJ IDEA project including all IAA microservices, having `common` has root project.

### Usage

Just copy the `.run` folder into the `common` project (or whatever is the root project in your configuration), they should automatically be recognized by IntelliJ IDEA.

### Configurations

The `.run` folders contains two type of Run configurations:

* Application Run configurations: the actual application configurations

* Compound Run configurations: containers allowing multiple applications to be started at once

The following is an overview of configurations required for Application Run configurations:

| Name                                   | Spring Profile     | Env Vars                                      | JVM Options                            |
|----------------------------------------|---------------------|-----------------------------------------------|----------------------------------------|
| AuthenticationProvider-Local-Authority | local-authority    | CRYPTO_SECRETKEYBASE64=KdSucCE6...            | -Djava.net.preferIPv4Stack=true        |
| AuthenticationProvider-Local-Consumer  | local-consumer     | CRYPTO_SECRETKEYBASE64=KdSucCE6...            | -Djava.net.preferIPv4Stack=true        |
| IdentityProvider-Local                 | local              | *(none)*                                      | *(none)*                               |
| Onboarding-Local                       | local              | *(none)*                                      | *(none)*                               |
| SecurityAttributesProvider-Local       | local              | *(none)*                                      | *(none)*                               |
| TierOne-Local-Authority                | local-authority    | *(none)*                                      | -Dio.netty.resolver.noPreferNative=true |
| TierOne-Local-Consumer                 | local-consumer     | *(none)*                                      | *(none)*                               |
| TierTwo-Local-Authority                | local-authority    | *(none)*                                      | *(none)*                               |
| TierTwo-Local-Consumer                 | local-consumer     | *(none)*                                      | *(none)*                               |
| UsersRoles-Local-Authority             | local-authority    | *(none)*                                      | *(none)*                               |
| UsersRoles-Local-Consumer              | local-consumer     | *(none)*                                      | *(none)*                               |
