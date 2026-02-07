function fn() {
  var config = {};
  // default base URL; can be overridden with -DbaseUrl=http://localhost:8080
  config.baseUrl = karate.properties['baseUrl'] || 'http://localhost:8080';
  return config;
}
