"c:\Program Files\Java\jdk1.8.0_211\bin\java.exe" -server -Xms1024M -Xmx1024M -XX:MetaspaceSize=128M -XX:MaxMetaspaceSize=128M -XX:+PrintGCDetails -XX:+PrintGCTimeStamps -XX:+PrintGCDateStamps -XX:+PrintGCCause -Xloggc:gc.log -jar super-spring-app.jar --spring.profiles.active=prod