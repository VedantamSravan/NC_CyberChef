var app = angular.module("myApp", ["ngRoute"]);

app.config(function($routeProvider) {
  $routeProvider
    .when("/cyberchef", {
      template: '<iframe src="/cyberchef/" title="CyberChef"></iframe>'
    })
    .when("/status", {
      template: '<h2>Server Status</h2><p>CyberChef is available and server is running securely over HTTPS.</p>'
    })
    .otherwise({
      redirectTo: "/cyberchef"
    });
});
