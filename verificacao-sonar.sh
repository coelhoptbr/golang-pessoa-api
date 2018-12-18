# para rodar: bash estearquivo.sh

sonar-scanner \
  -Dsonar.projectKey=pessoa-api-go \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://localhost:9000 \
  -Dsonar.login=f8de502fc8803835323910dfd8dcbf1eaf2739d3