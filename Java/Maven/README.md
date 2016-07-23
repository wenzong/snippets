Maven
=============

## Reference

[Maven Getting Started Guide](https://www.google.com/search?q=Maven+Getting+Started+Guide)

## Create a Project

```
mvn archetype:generate \
    -DarchetypeArtifactId=maven-archetype-quickstart \
    -DgroupId=org.test.app \
    -DartifactId=testapp \
    -DinteractiveMode=false
```

## Maven Commands

+ versions(Updates):
    + Dependency: `mvn versions:display-dependency-updates`
    + Plugin: `mvn versions:display-plugin-updates`
+ dependency:
    + `mvn dependency:list`
    + `mvn dependency:tree -Ddetail`
    + `mvn dependency:build-classpath`
    + `mvn dependency:sources`(`mvn eclipse:eclipse`)
+ help:
    + `mvn help:effective-pom`
    + `mvn help:effective-settings`
    + `mvn help:system`

## Maven Plugins

+ compiler
+ checkstyle
+ eclipse
+ enforcer
