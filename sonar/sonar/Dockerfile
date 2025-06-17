FROM sonarqube:community

ENV PLUGIN_1_URL=https://github.com/spotbugs/sonar-findbugs/releases/download/4.5.2/sonar-findbugs-plugin-4.5.2.jar
ENV PLUGIN_2_URL=https://github.com/jborgers/sonar-pmd/releases/download/4.0.3/sonar-pmd-plugin-4.0.3.jar

ENV SONARQUBE_EXTENSIONS_DIR=/opt/sonarqube/extensions/plugins

RUN curl -k -L $PLUGIN_1_URL -o $SONARQUBE_EXTENSIONS_DIR/sonar-github-plugin.jar
RUN curl -k -L $PLUGIN_2_URL -o $SONARQUBE_EXTENSIONS_DIR/sonar-jenkins-plugin.jar

